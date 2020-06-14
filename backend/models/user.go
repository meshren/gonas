package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gonas/utils"
	"log"
	"time"
)

type UserToken struct {
	ID        uint
	Username  string
	ExpiredAt time.Time
}

type User struct {
	BaseModel
	Username    string      `json:"username"`
	Password    string      `json:"password"`
	DeviceID    uint        `json:"device_id"`
	LastLoginIp uint32      `json:"last_login_ip"`
	LastLoginAt time.Time   `json:"last_login_at"`
	LoginCount  uint64      `json:"login_count"`
	Files       []File      `gorm:"many2many:user_files",json:"files"`
	Directories []Directory `json:"directories"`
	DirectoryID uint
}

func (u *User) directories(directoryID uint) (directories []Directory, err error) {
	panic("implement me")
}

func (u *User) GenToken() (string, error) {
	userToken, err := u.GenUserToken()
	if err != nil {
		return "", err
	}
	return EncryptedToken(&userToken)
}

func FindUserByID(id uint) (user User, err error) {
	db, err := connection()
	log.Println("db1: ", db)
	log.Println("err: ", err)
	if err != nil {
		utils.ErrDetail(err)
		return
	}
	defer db.Close()
	log.Println("db2: ", db)
	db.Where("id=?", id).First(&user)
	return
}

func (u *User) UserFiles() (files []File, err error) {
	return FilesByUser(u)
}

func FilesDirectoryUserID(userID uint, directoryID uint) (files []File, err error) {
	db, err := connection()
	if err != nil {
		return
	}
	var user User
	user.ID = userID
	db.Model(&user).Where("user_files.directory_id=?", directoryID).Related(&files, "Files")
	return
}

func FilesByUser(user *User) (files []File, err error) {
	db, err := connection()
	if err != nil {
		return
	}
	db.Model(&user).Related(&files, "Files")
	return
}

func (u *User) Login() (token string, err error) {
	db, err := connection()
	if err != nil {
		return "", err
	}
	defer db.Close()
	saltPwd := u.Password + viper.GetString("backend.manager-salt")
	passMd5 := utils.Md5Encrypt(saltPwd)
	var user User
	db.Where("username=? AND password=?", u.Username, passMd5).First(&user)
	if user.ID == 0 {
		return "", nil
	}
	userToken, err := user.GenUserToken()
	if err != nil {
		return
	}
	return EncryptedToken(&userToken)
}

func (u *User) GenUserToken() (token UserToken, err error) {
	expiredAt, err := utils.AddTime(time.Now(), "+24h")
	if err != nil {
		utils.ErrDetail(err)
		return
	}
	token = UserToken{
		ID:        u.ID,
		Username:  u.Username,
		ExpiredAt: expiredAt,
	}
	return
}

func EncryptedToken(userToken *UserToken) (token string, err error) {
	tokenB, err := json.Marshal(userToken)
	if err != nil {
		return
	}
	tokenS := string(tokenB)
	key := viper.GetString("backend.manager-key")
	token, err = utils.EnPwdCode([]byte(tokenS), []byte(key))
	return
}

func (u *User) Create() error {
	db, err := connection()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(u)
	return nil
}

type UserI interface {
	Login() (string, error)
	Create() error
	GenToken() (string, error)
	UserFiles() ([]File, error)
}

// 解密用户token
func DecryptToken(t string) (token UserToken, err error) {
	key := viper.GetString("backend.manager-key")
	tokenB, err := utils.DePwdCode(t, []byte(key))
	err = json.Unmarshal(tokenB, &token)
	return
}

func CheckAuth(c *gin.Context) (userID uint, err error) {
	token, err := c.Cookie("Admin-Token")
	if err != nil {
		return 0, err
	}
	decryptToken, err := DecryptToken(token)
	if err != nil {
		return
	}
	return decryptToken.ID, nil
}

func UserDirectories(userID uint, directoryPID uint) (directories []Directory, err error) {
	db, err := connection()
	if err != nil {
		return
	}
	db.Where("user_id=? AND parent_id=?", userID, directoryPID).Find(&directories)
	return
}

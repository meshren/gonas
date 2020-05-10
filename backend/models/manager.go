package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"gonas/utils"
	"time"
)

type Manager struct {
	gorm.Model
	Username    string `json:"username"`
	Password    string `json:"password"`
	LastLoginAt time.Time
	LastLoginIp uint32
	LoginCount  uint64
}

type ManagerToken struct {
	ID        uint
	Username  string
	ExpiredAt time.Time
}

func FindManagerByID(id int) (error, Manager) {
	db, err := connection()
	defer db.Close()
	if err != nil {
		utils.ErrDetail(err)
		return err, Manager{}
	}
	manager := Manager{}
	db.First(&manager, id)
	return nil, manager
}

func (m *Manager) CheckPassword() (ok bool, err error) {
	db, err := connection()
	defer db.Close()
	manager := Manager{}
	if err != nil {
		return false, err
	}
	saltPwd := m.Password + viper.GetString("backend.manager-salt")
	passMd5 := utils.Md5Encrypt(saltPwd)
	db.Debug().Where("username=? AND password=?", m.Username, passMd5).First(&manager)
	m.ID = manager.ID
	return manager.ID != 0, nil
}

func GenToken(m *Manager) (token ManagerToken, err error) {
	expiredAt, err := utils.AddTime(time.Now(), "+24h")
	if err != nil {
		utils.ErrDetail(err)
		return
	}
	token = ManagerToken{
		ID:        m.ID,
		Username:  m.Username,
		ExpiredAt: expiredAt,
	}
	return
}

func EncryptToken(token *ManagerToken) (t string, err error) {
	tokenB, err := json.Marshal(token)
	if err != nil {
		return
	}
	tokenS := string(tokenB)
	key := viper.GetString("backend.manager-key")
	t, err = utils.EnPwdCode([]byte(tokenS), []byte(key))
	return
}

func DecryptToken(t string) (token ManagerToken, err error) {
	key := viper.GetString("backend.manager-key")
	tokenB, err := utils.DePwdCode(t, []byte(key))
	err = json.Unmarshal(tokenB, &token)
	return
}

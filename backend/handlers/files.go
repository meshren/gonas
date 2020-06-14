package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"gonas/models"
	"gonas/utils"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
)

func UploadFiles(c *gin.Context) (err error){
	file, header, err := c.Request.FormFile("file")
	directoryIDS := c.Request.Form.Get("directory_id")
	directoryID, err := strconv.Atoi(directoryIDS)
	if err != nil {
		directoryID = 0
	}
	if err != nil {
		utils.ClientJson(c, http.StatusBadRequest, "", utils.CodeProcessFailed, "读取上传文件失败，请重试！")
		return
	}
	hash, err := fileMd5(file)
	if err != nil {
		utils.ClientJson(c, http.StatusBadRequest, "", utils.CodeProcessFailed, "获取上传文件哈希值失败，请重试！")
		return
	}
	fileName := header.Filename
	log.Printf("file name: %s, hash: %s\n", fileName, hash)
	pwd, err := os.Getwd()
	if err != nil {
		utils.ClientJson(c, http.StatusBadRequest, "", utils.CodeProcessFailed, "获取上传文件目录失败，请重试！")
		return
	}
	dst := path.Join(pwd, "/storage", hash)
	if err = c.SaveUploadedFile(header, dst); err != nil {
		utils.ClientJson(c, http.StatusBadRequest, "", utils.CodeProcessFailed, "保存文件失败，请重试！")
		return
	}
	// todo 获取文件类型
	fileModel := &models.File{
		Display:      fileName,
		Hash:         hash,
		Type:         0,
		Size:         header.Size,
		OriginalName: fileName,
		DeviceID:     0,
	}
	var fileID uint
	if fileID, err = fileModel.Create(); err != nil {
		utils.ErrDetail(err)
		utils.ClientJson(c, http.StatusBadRequest, "", utils.CodeProcessFailed, "新增文件记录失败，请重试！")
	}
	var userFile models.UserFiles
	userFile.UserID, _ = models.CheckAuth(c)
	userFile.FileID = fileID
	userFile.DirectoryID = uint(directoryID)
	if _, err := userFile.Create(); err != nil {
		return err
	}
	utils.ClientJson(c, http.StatusOK, "", utils.CodeSuccess, "上传成功！")
	return
}

func createFile(bytes []byte) (file *os.File, err error) {
	fileName := "tmpFile"
	file, err = os.Create(fileName)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		return
	}
	return
}

func AllFiles(c *gin.Context) error {
	userID, err := models.CheckAuth(c)
	if err != nil {
		utils.ClientJson(c, http.StatusBadRequest, "", utils.CodeProcessFailed, "用户登录超时")
		return nil
	}
	directoryIDString := c.Query("directory_id")
	directoryID, err := strconv.Atoi(directoryIDString)
	if err != nil {
		utils.ClientJson(c, http.StatusBadRequest, "", utils.CodeProcessFailed, "用户目录不存在！")
		return err
	}
	type Result struct {
		Files       []models.File      `json:"files"`
		Directories []models.Directory `json:"directories"`
	}
	var result Result
	files, err := models.FilesDirectoryUserID(userID, uint(directoryID))
	directories, err := models.UserDirectories(userID, uint(directoryID))
	if err != nil {
		utils.JsonRepose(c, result)
		return err
	}
	result.Files = files
	result.Directories = directories
	utils.JsonRepose(c, result)
	return nil
}

func fileMd5(file multipart.File) (md5Str string, err error) {
	h := md5.New()
	if _, err := file.Seek(0, 0); err != nil {
		return md5Str, err
	}
	if _, err := io.Copy(h, file); err != nil {
		return md5Str, err
	}
	md5Str = hex.EncodeToString(h.Sum(nil))
	return md5Str, nil
}

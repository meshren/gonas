package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"gonas/utils"
)

// mysql 连接
func connection() (db *gorm.DB, err error) {
	user := viper.GetString("backend.mysql.user")
	password := viper.GetString("backend.mysql.password")
	host := viper.GetString("backend.mysql.host")
	port := viper.GetString("backend.mysql.port")
	database := viper.GetString("backend.mysql.database")
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err = gorm.Open("mysql", args)
	if err != nil {
		utils.ErrDetail(err)
		return
	}
	return
}

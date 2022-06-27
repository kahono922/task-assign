package utils

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"task-assign/models"
)

var DB *gorm.DB
var err error
var (
	user = "root"
	password = "1234"
	host = "localhost"
	port = "3306"
	dbname = "task_assign"
)
func Connect(){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	DB, err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		panic(err)
	}
	automigrate(&models.Worker{})
}
func automigrate(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
}
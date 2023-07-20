package config

import (
	"fmt"
	"github.com/tdatIT/gin-gorm-restapi/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = "123123@Dat"
	dbName   = "gin_db_api"
)

var DB *gorm.DB

func InitConnection() {
	defer fmt.Println("Database has connected")

	cnt := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(cnt))
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&model.Product{})
}

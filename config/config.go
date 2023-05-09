package config

import (
	"fmt"

	"crud/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DBUsername string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
}

var JWTSecret = []byte("secret_key")

func InitDB() {
	config := Config{
		DBUsername: "root",
		DBPassword: "",
		DBPort:     "3306",
		DBHost:     "127.0.0.1",
		DBName:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.Agents{})
	DB.AutoMigrate(&models.Customers{})
	DB.AutoMigrate(&models.Properties{})
	DB.AutoMigrate(&models.Transactions{})
	DB.AutoMigrate(&models.Users{})
}

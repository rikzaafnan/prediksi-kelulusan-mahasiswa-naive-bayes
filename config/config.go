package config

import (
	"prediksi-kelulusan-mahasiswa-naive-bayes/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/prediksi-kelulusan?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&models.ProfileLulusan{}, &models.User{})
}

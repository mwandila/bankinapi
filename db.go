package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100"`
	Email string `gorm:"unique;not null"`
	Age   int    `gorm:"default:3"`
}

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	err = DB.AutoMigrate(&User{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}

	// Ensure all existing users have an email before adding the NOT NULL constraint
	DB.Model(&User{}).Where("email IS NULL").Update("email", "default@example.com")
	// Add NOT NULL constraint
	DB.Exec("ALTER TABLE users ALTER COLUMN email SET NOT NULL")
}

package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	FirstName         string `json:"first_name" gorm:"size:100"`
	LastName          string `json:"last_name" gorm:"size:100"`
	Email             string `json:"email" gorm:"unique;not null"`
	Phone             string `json:"phone" gorm:"size:20"`
	Address           string `json:"address" gorm:"size:255"`
	IdentityCard      string `json:"identity_card" gorm:"size:20"`
	IdentityCardImage string `json:"identity_card_image" gorm:"size:255"`
	ProfileImage      string `json:"profile_image" gorm:"size:255"`
}
type BankAccount struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	UserID        uint    `json:"user_id" gorm:"foreignkey:User"`
	AccountNumber string  `json:"account_number" gorm:"size:20"`
	AccountType   string  `json:"account_type" gorm:"size:20"`
	Balance       float64 `json:"balance"`
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
	err = DB.AutoMigrate(&BankAccount{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}

}

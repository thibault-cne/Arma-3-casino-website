package clientservices

import (
	"fmt"

	"casino.website/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateNewClient(username string, accessType int) string {
	client := NewClient(username, accessType)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("An error occured while openning the database : %s", err.Error())
		return ""
	}

	password := client.Password
	hash, err := hashPassword(password)

	if err != nil {
		fmt.Printf("An error occured while hashing the password : %s", err.Error())
		return ""
	}

	client.Password = hash
	db.Save(client)
	return password
}

func CheckClientModerationType(userId int) bool {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("An error occured while openning the database : %s", err.Error())
		return false
	}

	var user models.Client

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return false
	}

	if user.AccessType >= 2 {
		return true
	}

	return false
}

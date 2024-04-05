package main

import (
	"oryzonly/user"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/oryzonly?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	userRepository := user.NewRepository(db)

	user := user.User{
		Name:           "Ardia",
		Email:          "ardi@example.com",
		PasswordHash:   "password123",
		AvatarFileName: "ardi.jpg",
		Role:           "User",
		Occupation:     "Programmer",
		Phone:          "081234567687",
		Nationality:    "Indonesia",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	userRepository.Create(user)
}

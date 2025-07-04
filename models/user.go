package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

func (u *User) HashPassword(password string) error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hasedPassword)

	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
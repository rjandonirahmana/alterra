package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Name     string `json:"name" form:"name" `
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"tooken" form:"token"`
}

const SECRET_KEY = "golang alterra"

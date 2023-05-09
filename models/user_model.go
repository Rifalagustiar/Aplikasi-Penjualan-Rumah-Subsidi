package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	UserType string `json:"user_type" form:"user_type"`
}

type UserResponse struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	UserType string `json:"user_type" form:"user_type"`
	Token    string `json:"token" form:"token"`
}

type SuperAdmin struct {
	gorm.Model
	Users
}

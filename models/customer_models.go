package models

import "github.com/jinzhu/gorm"

type Customers struct {
	gorm.Model
	FirstName    string         `json:"first_name" form:"first_name" gorm:"not null"`
	LastName     string         `json:"last_name" form:"last_name" gorm:"not null"`
	Email        string         `json:"email" form:"email" gorm:"unique;not null"`
	Phone        string         `json:"phone" form:"phone" gorm:"not null"`
	Transactions []Transactions `json:"transactions,omitempty" gorm:"foreignkey:CustomerID"`
}

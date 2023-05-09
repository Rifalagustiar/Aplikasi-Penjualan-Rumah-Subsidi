package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Transactions struct {
	gorm.Model
	PropertyID       uint       `json:"property_id" form:"property_id" gorm:"not null"`
	Property         Properties `json:"property" gorm:"foreignkey:PropertyID"`
	CustomerID       uint       `json:"customer_id" form:"customer_id" gorm:"not null"`
	Customer         Customers  `json:"customer" gorm:"foreignkey:CustomerID"`
	TransactionDate  time.Time  `json:"transaction_date" form:"transaction_date" gorm:"not null"`
	Harga            int        `json:"harga" form:"harga" gorm:"not null"`
	PaymentMethod    string     `json:"payment_method" form:"payment_method"`
	PaymentReference string     `json:"payment_reference" form:"payment_reference"`
}

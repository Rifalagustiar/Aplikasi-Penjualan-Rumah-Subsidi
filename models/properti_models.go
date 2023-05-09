package models

import "github.com/jinzhu/gorm"

type Properties struct {
	gorm.Model
	Alamat        string         `json:"alamat" form:"alamat" gorm:"not null"`
	PropertyType  string         `json:"property_type" form:"property_type" gorm:"not null"`
	Kamar_tidur   int            `json:"kamar_tidur" form:"kamar_tidur" gorm:"not null"`
	Kamar_mandi   int            `json:"kamar_mandi" form:"kamar_mandi" gorm:"not null"`
	Luas_tanah    int            `json:"luas_tanah" form:"luas_tanah" gorm:"not null"`
	Luas_bangunan int            `json:"luas_bangunan" form:"luas_bangunan" gorm:"not null"`
	Harga         int            `json:"harga" form:"harga" gorm:"not null"`
	Deskripsi     string         `json:"deskripsi" form:"deskripsi"`
	Image         string         `json:"image" form:"image"`
	Transactions  []Transactions `json:"transactions,omitempty" gorm:"foreignkey:PropertyID"`
}

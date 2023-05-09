package models

import "github.com/jinzhu/gorm"

type Agents struct {
	gorm.Model
	Name       string            `json:"name" form:"name" gorm:"not null"`
	Email      string            `json:"email" form:"email" gorm:"unique;not null"`
	Hp         string            `json:"hp" form:"hp" gorm:"not null"`
	Deskripsi  string            `json:"deskripsi" form:"deskripsi"`
	Properties []AgentProperties `json:"properties,omitempty" gorm:"foreignkey:AgentID"`
}

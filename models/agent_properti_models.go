package models

import "github.com/jinzhu/gorm"

type AgentProperties struct {
	gorm.Model
	AgentID    uint `json:"agent_id" form:"agent_id" gorm:"not null"`
	PropertyID uint `json:"property_id" form:"property_id" gorm:"not null"`
}

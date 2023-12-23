package models

import (
	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	gorm.Model
	Name  string `json:"name" binding:"required"`
	Age   uint8  `json:"age" binding:"required"`
	Email string `json:"email" binding:"required"`
}

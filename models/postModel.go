package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

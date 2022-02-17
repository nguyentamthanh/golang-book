package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Id       uint `json:"id" gorm:"primaryKey"`
	CreateAt time.Time
	Title    string `json:"title"`
	Author   string `json:"author"`
	Rating   int    `json:"rating"`
	Year     int    `json:"year"`
}

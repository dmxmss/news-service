package entities

import (
	"gorm.io/gorm"
	"time"
)

type News struct {
	gorm.Model
	Title string `gorm:"not null"`
	Contents string
	AuthorID int
	user User `gorm:"foreignKey:AuthorID"`
	Approved bool `gorm:"default:false"`
	ApprovedAt time.Time
	Tags []Tag `gorm:"many2many:news_tags"`
	Source string `gorm:"not null"`
}

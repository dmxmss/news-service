package entities

import (
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Title string `gorm:"not null"`
	News []News `gorm:"many2many:collection_news"`
	AuthorID int
	User User `gorm:"foreign_key:AuthorID"`
}

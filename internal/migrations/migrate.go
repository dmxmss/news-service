package main

import (
	"github.com/CherryRadiator/hakathon2025Spring/entities"
	"github.com/CherryRadiator/hakathon2025Spring/internal"
	"github.com/CherryRadiator/hakathon2025Spring/config"
	"gorm.io/gorm"
)

func main() {
	conf := config.GetConfig()

	tasksRepo, err := internal.NewPgNewsRepository(conf)
	if err != nil {
		panic(err)
	}

	migrate(tasksRepo.GetDb())
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.News{})
	db.AutoMigrate(&entities.User{})
}

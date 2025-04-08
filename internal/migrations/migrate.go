package main

import (
	"github.com/dmxmss/news-service/config"
	"github.com/dmxmss/news-service/internal"
	"github.com/dmxmss/news-service/entities"
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

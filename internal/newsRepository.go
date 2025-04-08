package internal

import (
	"github.com/dmxmss/news-service/entities"
	"github.com/dmxmss/news-service/config"
	e "github.com/dmxmss/news-service/error"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"fmt"
	"errors"
)

type NewsRepository interface {
	PostNews(int, entities.PostNewsDto) error
	GetNewsById(int) (*entities.News, error)
	PatchNewsById(int, entities.PatchNewsDto) error
	DeleteNewsById(int) (*entities.News, error)
	SearchNews(*entities.SearchNewsParams) ([]entities.News, error)
	GetDb() *gorm.DB
}

type NewsPgRepository struct {
	db *gorm.DB
}

func NewPgNewsRepository(config *config.Config) (NewsRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})	
	if err != nil {
		return nil, e.ErrDbInitError
	}
	
	return &NewsPgRepository{db}, nil
}

func (nr *NewsPgRepository) PostNews(user_id int, news entities.PostNewsDto) error {
	result := nr.db.Create(&entities.News{Title: news.Title, Contents: news.Contents, AuthorID: user_id, Tags: news.Tags})
	if result.Error != nil {
		return e.ErrDbTransactionFailed
	}
	return nil
}

func (nr *NewsPgRepository) GetNewsById(id int) (*entities.News, error) {
	var news entities.News
	result := nr.db.Take(&news, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, e.ErrDbNewsNotFound
		}
	}

	return &news, nil
}

func (nr *NewsPgRepository) GetDb() *gorm.DB {
	return nr.db
}

func (nr *NewsPgRepository) PatchNewsById(id int, patchNews entities.PatchNewsDto) error {
	result := nr.db.Model(&entities.News{}).Where("id = ?", id).Updates(patchNews)
	if result.Error != nil {
		return e.ErrDbTransactionFailed
	}

	return nil
}

func (nr *NewsPgRepository) DeleteNewsById(id int) (*entities.News, error) {
	news, err := nr.GetNewsById(id)
	if err != nil {
		return nil, err
	}
	result := nr.db.Delete(&news)
	if result.Error != nil {
		return nil, e.ErrDbTransactionFailed
	}
	return news, nil
}

func (nr *NewsPgRepository) SearchNews(params *entities.SearchNewsParams) ([]entities.News, error) {
	query := nr.db.Model(&entities.News{}).Preload("Tags")

	if params.Keyword != nil {
		query = query.Where("title LIKE ? OR contents LIKE ?", "%"+*params.Keyword+"%", "%"+*params.Keyword+"%")
	}

	var news []entities.News
	err := query.Order("created_at DESC").Find(&news).Error
	return news, err
}

package service

import (
	"github.com/CherryRadiator/hakathon2025Spring/entities"
	e "github.com/CherryRadiator/hakathon2025Spring/error"
	"github.com/CherryRadiator/hakathon2025Spring/internal"
	"github.com/CherryRadiator/hakathon2025Spring/config"
)

type NewsService interface {
	PostNews(int, entities.PostNewsDto) error
	GetNewsById(int) (*entities.News, error)
	PatchNewsById(int, int, entities.PatchNewsDto) error
	DeleteNewsById(int) (*entities.DeleteNewsDto, error)
	SearchNews(*entities.SearchNewsParams) ([]entities.News, error)
}

func NewNewsService(conf *config.Config) (NewsService, error) {
	newsRepo, err := internal.NewPgNewsRepository(conf)
	if err != nil {
		return nil, err
	}

	return &NewsServiceImpl{newsRepo: newsRepo}, nil
}

type NewsServiceImpl struct {
	newsRepo internal.NewsRepository
}

func (ns *NewsServiceImpl) PostNews(user_id int, news entities.PostNewsDto) error {
	return ns.newsRepo.PostNews(user_id, news)
}

func (ns *NewsServiceImpl) GetNewsById(id int) (*entities.News, error) {
	return ns.newsRepo.GetNewsById(id)
}

func (ns *NewsServiceImpl) PatchNewsById(user_id int, news_id int, patchNews entities.PatchNewsDto) error {
	news, err := ns.newsRepo.GetNewsById(news_id)
	if err != nil {
		return err
	}

	if news.AuthorID != user_id {
		return e.ErrUserIsNotAuthor
	}

	return ns.newsRepo.PatchNewsById(news_id, patchNews)
}

func (ns *NewsServiceImpl) DeleteNewsById(id int) (*entities.DeleteNewsDto, error) {
	news, err := ns.newsRepo.DeleteNewsById(id)
	if err != nil {
		return nil, err
	}

	return &entities.DeleteNewsDto{Title: news.Title, Contents: news.Contents, AuthorID: news.AuthorID}, err
}

func (ns *NewsServiceImpl) SearchNews(params *entities.SearchNewsParams) ([]entities.News, error) {
	return ns.newsRepo.SearchNews(params)	
}

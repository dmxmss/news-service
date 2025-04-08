package server

import (
	"strings"

	"github.com/dmxmss/news-service/config"
	"github.com/dmxmss/news-service/entities"
	e "github.com/dmxmss/news-service/error"
	"github.com/dmxmss/news-service/service"
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"
	"strconv"
)

type GinServer struct {
	app *gin.Engine
	conf *config.Config
	newsService service.NewsService
}

type SearchNewsParams = entities.SearchNewsParams

func NewGinServer(conf *config.Config) (*GinServer, error) {
	r := gin.Default()
	newsService, err := service.NewNewsService(conf)
	if err != nil {
		return nil, err
	}

	r.Use(ErrorCatchMiddleware())


	return &GinServer{app: r, newsService: newsService, conf: conf}, nil
}

func (s *GinServer) RegisterHandlers(conf *config.Config) {
	RegisterHandlersWithOptions(s.app, s, GinServerOptions{
		Middlewares: []MiddlewareFunc{
			func(c *gin.Context) {
				if strings.HasPrefix(c.Request.URL.Path, "/news") && (c.Request.Method == "POST" || c.Request.Method == "PATCH" || c.Request.Method == "DELETE") {
					s.JWTAuthMiddleware(conf)(c)
					return
				}
				c.Next()
			},
		},
	})
}

func (s *GinServer) Start() {
	s.app.Run(fmt.Sprintf("%s:%s", s.conf.App.Address, s.conf.App.Port))
}

func (s *GinServer) SearchNews(c *gin.Context, params SearchNewsParams) {
	news, err := s.newsService.SearchNews(&params)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, news)
}

func (s *GinServer) MakeNews(c *gin.Context) {
	v, exists := c.Get("user_id")
	user_id, ok := v.(string)
	if !exists || !ok {
		c.Error(e.ErrNotAuthorized)
		return
	}

	var news entities.PostNewsDto
	if err := c.ShouldBindJSON(&news); err != nil {
		c.Error(e.ErrInvalidRequestData)
		return
	}

	id, _ := strconv.Atoi(user_id)
	err := s.newsService.PostNews(id, news)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (s *GinServer) GetNewsById(c *gin.Context, id int) {
	news, err := s.newsService.GetNewsById(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, news)
}

func (s *GinServer) PatchNewsById(c *gin.Context, id int) {
	v, exists := c.Get("user_id")
	obj, ok := v.(string)
	if !exists || !ok {
		c.Error(e.ErrNotAuthorized)
		return
	}
	user_id, _ := strconv.Atoi(obj)

	var patchNews entities.PatchNewsDto
	if err := c.ShouldBindJSON(&patchNews); err != nil {
		c.Error(e.ErrInvalidRequestData)
		return
	}

	err := s.newsService.PatchNewsById(user_id, id, patchNews)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (s *GinServer) DeleteNewsById(c *gin.Context, id int) {
	news, err := s.newsService.DeleteNewsById(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, news)
}

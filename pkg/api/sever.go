package api

import (
	"os"

	"github.com/afthaab/urlshortner/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(urlHandler *handler.UrlHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	// defining the EndPoints
	engine.GET("/:url", urlHandler.ResolveUrl)
	engine.POST("/api/v1", urlHandler.ShortenUrl)

	return &ServerHTTP{
		engine: engine,
	}
}

func (s *ServerHTTP) Start() {
	s.engine.Run(os.Getenv("APP_PORT"))
}

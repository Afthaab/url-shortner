package handler

import (
	"net/http"

	"github.com/afthaab/urlshortner/pkg/domain"
	services "github.com/afthaab/urlshortner/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	urlUseCase services.UrlUseCase
}

func NewUrlHandler(urlusecase services.UrlUseCase) *UrlHandler {
	return &UrlHandler{
		urlUseCase: urlusecase,
	}
}

func (u *UrlHandler) ShortenUrl(c *gin.Context) {
	bodyReq := domain.Request{}

	err := c.Bind(&bodyReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error in binding the JSON Data",
		})
		return
	}
	ip := c.ClientIP()
	statuscode, err, res := u.urlUseCase.ShortenUrl(bodyReq, ip)
	if err != nil {
		c.JSON(statuscode, gin.H{
			"Error":    err,
			"Response": res,
		})
		return
	}

}

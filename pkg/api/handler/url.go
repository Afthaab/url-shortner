package handler

import (
	services "github.com/afthaab/urlshortner/pkg/usecase/interface"
)

type UrlHandler struct {
	urlUseCase services.UrlUseCase
}

func NewUrlHandler(urlusecase services.UrlUseCase) *UrlHandler {
	return &UrlHandler{
		urlUseCase: urlusecase,
	}
}

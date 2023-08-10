package usecase

import (
	interfaces "github.com/afthaab/urlshortner/pkg/repository/interface"
	services "github.com/afthaab/urlshortner/pkg/usecase/interface"
)

type urlUseCase struct {
	urlRepo interfaces.UrlRepository
}

func NewUrlUseCase(repo interfaces.UrlRepository) services.UrlUseCase {
	return &urlUseCase{
		urlRepo: repo,
	}
}

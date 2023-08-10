//go:build wireinject
// +build wireinject

package di

import (
	"github.com/afthaab/urlshortner/pkg/api"
	"github.com/afthaab/urlshortner/pkg/api/handler"
	interfaces "github.com/afthaab/urlshortner/pkg/repository/interface"
	"github.com/afthaab/urlshortner/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(repoInterface interfaces.UrlRepository) (*api.ServerHTTP, error) {
	wire.Build(usecase.NewUrlUseCase, handler.NewUrlHandler, api.NewServerHTTP)
	return &api.ServerHTTP{}, nil
}

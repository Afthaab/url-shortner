package interfaces

import "github.com/afthaab/urlshortner/pkg/domain"

type UrlUseCase interface {
	ShortenUrl(bodyReq domain.Request, ip string) (int, error, domain.Response)
	ResolveTheUrl(url string) (int, string, error)
}

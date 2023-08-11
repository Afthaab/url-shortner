package usecase

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/afthaab/urlshortner/pkg/domain"
	interfaces "github.com/afthaab/urlshortner/pkg/repository/interface"
	services "github.com/afthaab/urlshortner/pkg/usecase/interface"
	"github.com/afthaab/urlshortner/pkg/utility"
	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type urlUseCase struct {
	urlRepo interfaces.UrlRepository
}

func NewUrlUseCase(repo interfaces.UrlRepository) services.UrlUseCase {
	return &urlUseCase{
		urlRepo: repo,
	}
}

func (u *urlUseCase) ShortenUrl(bodyReq domain.Request, ip string) (int, error, domain.Response) {

	var limit time.Duration

	// implementing rate limiting
	val, err := u.urlRepo.FindTheIp(ip)
	if err == redis.Nil {
		err := u.urlRepo.SaveTheIp(ip)
		if err != nil {
			return http.StatusInternalServerError, errors.New("Could not save the ip in the database"), domain.Response{}
		}
	} else {
		value, _ := strconv.Atoi(val)
		if value <= 0 {
			// Checking the time remaining until the rate limit is reset.
			limit, _ = u.urlRepo.CheckTheTime(ip)
			return http.StatusServiceUnavailable, errors.New("Rate limit exceeded"), domain.Response{
				RateLimiting: limit,
			}
		}
	}

	// check if the input is an actual url
	if !govalidator.IsURL(bodyReq.URL) {
		return http.StatusBadRequest, errors.New("Invalid URL"), domain.Response{}
	}

	// check for the domain error
	if !utility.RemoverDomaiError(bodyReq.URL) {
		return http.StatusServiceUnavailable, errors.New("Cannot proceed with this domain"), domain.Response{}
	}

	// enforce https, SSL
	bodyReq.URL = utility.EnforceHTTP(bodyReq.URL)

	//check if the user provided any custom short URL
	var id string
	if bodyReq.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = bodyReq.CustomShort
	}

	// check if the user provided short is already there in the db
	value, _ := u.urlRepo.FindTheURL(id)
	if value == "" {
		return http.StatusForbidden, errors.New("Custo URL is already in use"), domain.Response{}
	}

	if bodyReq.Expiry == 0 {
		bodyReq.Expiry = 24
	}

	// set the url in the database
	err = u.urlRepo.SetTheUrl(bodyReq.URL, id, bodyReq.Expiry)
	if err != nil {
		return http.StatusInternalServerError, errors.New("Could not connect to the Database"), domain.Response{}
	}

	// create the response if all is well
	rateRemianing, _ := strconv.Atoi(val)
	response := domain.Response{
		URL:           bodyReq.URL,
		CustomShort:   os.Getenv("DOMAIN"),
		Expiry:        bodyReq.Expiry,
		RateRemaining: rateRemianing,
		RateLimiting:  limit,
	}

	return http.StatusOK, nil, response

}

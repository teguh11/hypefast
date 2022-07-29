package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/teguh11/hypefast/helpers"
	"github.com/teguh11/hypefast/models"
)

var redirectData []*models.ShortenData

const (
	hostURL = "http://localhost:8080"
)

type ShortenService struct {
}

type ShortenServiceInterface interface {
	ShortenURL(ctx context.Context, data models.ShortenRequest) (result string, err error)
	Statistic(randomString string) (data models.ShortenData, err error)
	RedirectURL(randomString string) (realURL string, err error)
}

func InitShortenService() *ShortenService {
	return &ShortenService{}
}

func (s *ShortenService) ShortenURL(ctx context.Context, data models.ShortenRequest) (result string, err error) {
	shortenURL, err := helpers.ShortenURL(data.URL, 6)
	if err != nil {
		return
	}

	shortenResult := fmt.Sprintf("%s/%s", hostURL, shortenURL)

	x := &models.ShortenData{
		URL:           data.URL,
		ShortenValue:  shortenResult,
		CreatedAt:     time.Now().String(),
		RedirectCount: 0,
	}
	redirectData = append(redirectData, x)
	y, _ := json.Marshal(redirectData)
	fmt.Println("data =>", string(y))
	return shortenResult, nil
}

func (s *ShortenService) Statistic(randomString string) (data models.ShortenData, err error) {
	shortenURL := fmt.Sprintf("%s/%s", hostURL, randomString)
	for _, v := range redirectData {
		if v.ShortenValue == shortenURL {
			return *v, nil
		}
	}
	return data, errors.New("data not found")
}

func (s *ShortenService) RedirectURL(randomString string) (realURL string, err error) {
	shortenURL := fmt.Sprintf("%s/%s", hostURL, randomString)
	for i, v := range redirectData {
		if v.ShortenValue == shortenURL {
			fmt.Println("url =>", v.URL)
			realURL = v.URL
			v.RedirectCount += 1
			redirectData[i] = v
		}
	}
	if realURL == "" {
		return realURL, errors.New("url not found")
	}
	y, _ := json.Marshal(redirectData)
	fmt.Println("data after =>", string(y))
	return
}

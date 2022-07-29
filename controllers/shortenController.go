package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teguh11/hypefast/helpers"
	"github.com/teguh11/hypefast/models"
	"github.com/teguh11/hypefast/services"
)

type ShortenControllers struct {
	Shorten services.ShortenServiceInterface
}

func InitShortenControllers() *ShortenControllers {
	shorten := &services.ShortenService{}
	return &ShortenControllers{
		Shorten: shorten,
	}
}

func (c *ShortenControllers) ShortenURL(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var shortenRequest models.ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&shortenRequest)
	if err != nil {
		helpers.Response(rw, http.StatusInternalServerError, map[string]interface{}{
			"err_msg": err.Error(),
			"data":    nil,
		})
		return
	}

	result, err := c.Shorten.ShortenURL(ctx, shortenRequest)
	if err != nil {
		helpers.Response(rw, http.StatusInternalServerError, map[string]interface{}{
			"err_msg": err.Error(),
			"data":    nil,
		})
		return
	}

	helpers.Response(rw, http.StatusOK, map[string]interface{}{
		"err_msg": nil,
		"data":    result,
	})
}

func (c *ShortenControllers) Statistic(rw http.ResponseWriter, r *http.Request) {
	randomString := mux.Vars(r)["randomString"]

	statURL, err := c.Shorten.Statistic(randomString)
	if err != nil {
		helpers.Response(rw, http.StatusOK, map[string]interface{}{
			"err_msg": nil,
			"data":    randomString,
		})
	}
	helpers.Response(rw, http.StatusOK, map[string]interface{}{
		"err_msg": nil,
		"data":    statURL,
	})
}

func (c *ShortenControllers) RedirectURL(rw http.ResponseWriter, r *http.Request) {
	randomString := mux.Vars(r)["randomString"]

	realURL, err := c.Shorten.RedirectURL(randomString)
	if err != nil {
		helpers.Response(rw, http.StatusOK, map[string]interface{}{
			"err_msg": nil,
			"data":    randomString,
		})
	}
	http.Redirect(rw, r, realURL, http.StatusFound)
}

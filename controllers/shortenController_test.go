package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teguh11/hypefast/controllers"
)

func TestStatistic(t *testing.T) {
	t.Parallel()
	r := httptest.NewRequest(http.MethodGet, "/stat/{randomString}", nil)
	w := httptest.NewRecorder()

	shortenController := controllers.InitShortenControllers()

	shortenController.Statistic(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

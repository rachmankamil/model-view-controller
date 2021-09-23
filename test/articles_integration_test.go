package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rachmankamil/kampus-merdeka-b/config"
	"github.com/rachmankamil/kampus-merdeka-b/controllers"
	"github.com/rachmankamil/kampus-merdeka-b/models"
	"github.com/stretchr/testify/assert"
)

func setup() {
	articles := []models.Article{
		{
			Title:  "test123",
			Author: "rachmankamil",
		},
		{
			Title:  "test456",
			Author: "kamil",
		},
	}
	config.DB.Create(&articles)
}

func TestGetArticleController(t *testing.T) {
	e := InitEcho()
	setup()
	defer TearDown()

	t.Run("valid test case", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/articles", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

		if assert.NoError(t, controllers.GetArticleController(echoContext)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "success")
			assert.Contains(t, rec.Body.String(), "{\"data\":[{")
		}

	})

}

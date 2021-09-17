package swapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type article struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

var articles = []article{
	{
		ID:        1,
		Name:      "Article1",
		CreatedAt: time.Now(),
	},
	{
		ID:        2,
		Name:      "Article2",
		CreatedAt: time.Now().AddDate(0, 0, -2),
	},
}

func getArticles(httpResponse http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		result, err := json.Marshal(articles)
		if err != nil {
			http.Error(httpResponse, err.Error(), http.StatusInternalServerError)
			return
		}

		httpResponse.Write(result)
		return
	}
	http.Error(httpResponse, "method not alowed", http.StatusBadRequest)
}

func getSWPlanet(httpResponse http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		result, _ := http.Get("https://swapi.dev/api/planets/1/")

		responseData, _ := ioutil.ReadAll(result.Body)
		defer result.Body.Close()

		var planetSW Planets
		json.Unmarshal(responseData, &planetSW)

		fmt.Printf("%+v", planetSW)
		result2, err := json.Marshal(planetSW)
		if err != nil {
			http.Error(httpResponse, err.Error(), http.StatusInternalServerError)
			return
		}

		httpResponse.Write(result2)
		return
	}
	http.Error(httpResponse, "method not alowed", http.StatusBadRequest)
}

func mainRaw() {
	http.HandleFunc("/articles", getArticles)
	http.HandleFunc("/getPlanets", getSWPlanet)

	fmt.Println("starting web......")
	http.ListenAndServe("0.0.0.0:8080", nil)
}

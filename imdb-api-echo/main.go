package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Search       interface{} `json:"Search"`
	TotalResults int         `json:"totalResults,string"`
	Response     bool        `json:"Response"`
}

type Imdb struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	Dvd        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}

func GetImdbByIdController(c echo.Context) error {
	resp, err := http.Get("https://www.omdbapi.com/?apikey=a818034f&i=" + c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": err,
		})
	}

	respData, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": err,
		})
	}

	var responseObject Imdb
	json.Unmarshal(respData, &responseObject)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get by id",
		"imdb":     responseObject,
	})
}

func SearchImdbController(c echo.Context) error {
	link := "https://www.omdbapi.com/?apikey=a818034f&s=" + c.QueryParam("search")
	if c.QueryParam("page") != "" {
		link += "&page=" + c.QueryParam("page")
	}

	if c.QueryParam("type") != "" {
		link += "&type=" + c.QueryParam("type")
	}
	resp, err := http.Get(link)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err,
		})
	}

	respData, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err,
		})
	}
	var responseObject Response
	json.Unmarshal(respData, &responseObject)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":     "success search imdb",
		"imdb":         responseObject.Search,
		"totalResults": responseObject.TotalResults,
	})
}

func main() {
	e := echo.New()
	e.GET("/:id", GetImdbByIdController)
	e.GET("/", SearchImdbController)
	e.Start(":8000")
}

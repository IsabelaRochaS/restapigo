package controller

import (
	"encoding/json"
	swapiModel "github.com/isabelarochas/restapigo/model"
	"net/http"
	"net/url"
)

type SwapiInterface interface {
	GetQuantityFilms(name string) int
}

type SwapiStruct struct {}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (s SwapiStruct) GetQuantityFilms(name string) int {
	swapi := swapiModel.Swapi{}
	_ = getJson("https://swapi.dev/api/planets/?search=" + url.QueryEscape(name), &swapi)
	if swapi.Count == 1 {
		return len(swapi.StarWarsPlanet[0].Films)
	}
	return 0
}
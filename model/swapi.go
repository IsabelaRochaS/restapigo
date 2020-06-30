package model

type Swapi struct {
	Count int `json:"count"`
	StarWarsPlanet []struct {
		Name string `json:"name"`
		Films []string `json:"films"`
	} `json:"results"`
}

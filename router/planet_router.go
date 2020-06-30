package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/isabelarochas/restapigo/config/dao"
	. "github.com/isabelarochas/restapigo/controller"
	. "github.com/isabelarochas/restapigo/model"
	"gopkg.in/mgo.v2/bson"
)

var planDao PlanetDAOInterface
var swapiController SwapiInterface

func init() {
	planDao = &PlanetDAO{}
	swapiController = SwapiStruct{}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	planets, err := planDao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}
	log.Println("Get All OK")
	respondWithJson(w, http.StatusOK, planets)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planet, err := planDao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Planet ID")
		log.Println(err.Error())
		return
	}
	log.Println("Get By ID return " + planet.Name)
	respondWithJson(w, http.StatusOK, planet)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var planet Planet
	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		log.Println(err.Error())
		return
	}
	planet.ID = bson.NewObjectId()
	planet.MoviesCount = swapiController.GetQuantityFilms(planet.Name)

	validate := validatePlanet(planet)

	if validate != "" {
		respondWithError(w, http.StatusInternalServerError, validate)
		log.Println(validate)
		return
	}

	if err := planDao.Create(planet); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}
	log.Println("Planet " + planet.Name + " created!")
	respondWithJson(w, http.StatusCreated, planet)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var planet Planet
	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		log.Println(err.Error())
		return
	}
	if err := planDao.Update(params["id"], planet); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}
	log.Println(planet.Name + "updated with success!")
	respondWithJson(w, http.StatusOK, map[string]string{"result": planet.Name + " updated with success!"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := planDao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}
	log.Println("Planet deleted with success!")
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func validatePlanet(planet Planet) string {
	if planet.Name == ""{
		return "Name is required."
	}
	if planet.Terrain == "" {
		return "Terrain is required"
	}
	if planet.Climate == "" {
		return "Climate is required"
	}
	return ""
}

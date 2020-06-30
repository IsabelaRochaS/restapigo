package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/isabelarochas/restapigo/config"
	. "github.com/isabelarochas/restapigo/config/dao"
	planetrouter "github.com/isabelarochas/restapigo/router"
)

var dao = PlanetDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/planets", planetrouter.GetAll).Methods("GET")
	r.HandleFunc("/planet/{id}", planetrouter.GetByID).Methods("GET")
	r.HandleFunc("/planet", planetrouter.Create).Methods("POST")
	r.HandleFunc("/planet/{id}", planetrouter.Update).Methods("PUT")
	r.HandleFunc("/planet/{id}", planetrouter.Delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}

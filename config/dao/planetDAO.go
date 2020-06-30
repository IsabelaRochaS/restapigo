package dao

import (
	"log"

	. "github.com/isabelarochas/restapigo/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PlanetDAOInterface interface {
	Connect()
	GetAll() ([]Planet, error)
	GetByID(id string) (Planet, error)
	Create(planet Planet) error
	Delete(id string) error
	Update(id string, planet Planet) error
}

type PlanetDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "planets"
)

func (m *PlanetDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
	log.Println("Connected with database.")
}

func (m *PlanetDAO) GetAll() ([]Planet, error) {
	var planets []Planet
	err := db.C(COLLECTION).Find(bson.M{}).All(&planets)
	return planets, err
}

func (m *PlanetDAO) GetByID(id string) (Planet, error) {
	var planet Planet
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&planet)
	return planet, err
}

func (m *PlanetDAO) Create(planet Planet) error {
	err := db.C(COLLECTION).Insert(&planet)
	return err
}

func (m *PlanetDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *PlanetDAO) Update(id string, planet Planet) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &planet)
	return err
}
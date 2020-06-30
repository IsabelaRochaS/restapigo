package router

import (
	"bytes"
	"errors"
	. "github.com/isabelarochas/restapigo/model"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"testing"
)

type swapiMock struct{}

type planetDaoInterfaceMock struct {}
type planetDaoInterfaceMockError struct {}

func (p planetDaoInterfaceMock) Connect() {}
func (p planetDaoInterfaceMockError) Connect() {}

func (p swapiMock) GetQuantityFilms(name string) int {
	return 2
}

func (p planetDaoInterfaceMock) GetAll() ([]Planet, error) {
	planet := Planet{ID: "1", Name: "Coruscant", Climate: "temperate", Terrain: "cityscape, mountains", MoviesCount: 4}
	return []Planet{planet}, nil
}

func (p planetDaoInterfaceMockError) GetAll() ([]Planet, error) {
	return nil, errors.New("error")
}

func (p planetDaoInterfaceMock) GetByID(id string) (Planet, error) {
	planet := Planet{ID: "2", Name: "Kamino", Climate: "temperate", Terrain: "ocean", MoviesCount: 1}
	return planet, nil
}

func (p planetDaoInterfaceMockError) GetByID(id string) (Planet, error) {
	return Planet{}, errors.New("error")
}

func (p planetDaoInterfaceMock) Create(planet Planet) error {
	return nil
}

func (p planetDaoInterfaceMockError) Create(planet Planet) error {
	return errors.New("error")
}

func (p planetDaoInterfaceMock) Delete(id string) error {
	return nil
}

func (p planetDaoInterfaceMockError) Delete(id string) error {
	return errors.New("error")
}

func (p planetDaoInterfaceMock) Update(id string, planet Planet) error {
	return nil
}

func (p planetDaoInterfaceMockError) Update(id string, planet Planet) error {
	return errors.New("error")
}

func TestGetAll(t *testing.T) {
	planDao = planetDaoInterfaceMock{}

	req, err := http.NewRequest("GET", "/planets", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAll)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, "[{\"id\":\"" + bson.ObjectId("1").Hex() + "\",\"name\":\"Coruscant\",\"climate\":\"temperate\",\"terrain\":\"cityscape, mountains\",\"movies_count\":4}]", resp.Body.String())
}

func TestGetAllError(t *testing.T) {
	planDao = planetDaoInterfaceMockError{}

	req, err := http.NewRequest("GET", "/planets", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAll)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
}

func TestGetByID(t *testing.T) {
	planDao = planetDaoInterfaceMock{}

	req, err := http.NewRequest("GET", "/planet/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(GetByID)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, "{\"id\":\"" + bson.ObjectId("2").Hex() + "\",\"name\":\"Kamino\",\"climate\":\"temperate\",\"terrain\":\"ocean\",\"movies_count\":1}", resp.Body.String())
}

func TestGetByIDError(t *testing.T) {
	planDao = planetDaoInterfaceMockError{}

	req, err := http.NewRequest("GET", "/planets", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(GetByID)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}

func TestCreate(t *testing.T) {
	planDao = planetDaoInterfaceMock{}
	swapiController = swapiMock{}

	var jsonStr = []byte("{\"name\":\"Tatooine\",\"terrain\":\"desert\",\"climate\":\"arid\"}")

	req, err := http.NewRequest("POST", "/planet", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(Create)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, 201, resp.Code)
}

func TestCreateError(t *testing.T) {
	planDao = planetDaoInterfaceMockError{}
	swapiController = swapiMock{}

	var jsonStr = []byte("{\"name\":\"Tatooine\",\"terrain\":\"desert\",\"climate\":\"arid\"}")

	req, err := http.NewRequest("POST", "/planet", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(Create)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
}

func TestUpdate(t *testing.T) {
	planDao = planetDaoInterfaceMock{}

	var jsonStr = []byte("{\"name\":\"Tatooine\",\"terrain\":\"desert\",\"climate\":\"arid\"}")

	req, err := http.NewRequest("PUT", "/planet", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(Update)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}

func TestUpdateError(t *testing.T) {
	planDao = planetDaoInterfaceMockError{}

	var jsonStr = []byte("{\"name\":\"Tatooine\",\"terrain\":\"desert\",\"climate\":\"arid\"}")

	req, err := http.NewRequest("PUT", "/planet", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(Update)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
}

func TestDelete(t *testing.T) {
	planDao = planetDaoInterfaceMock{}

	var jsonStr = []byte("{\"name\":\"Tatooine\",\"terrain\":\"desert\",\"climate\":\"arid\"}")

	req, err := http.NewRequest("DELETE", "/planet{id}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(Delete)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}

func TestDeleteError(t *testing.T) {
	planDao = planetDaoInterfaceMockError{}

	var jsonStr = []byte("{\"name\":\"Tatooine\",\"terrain\":\"desert\",\"climate\":\"arid\"}")

	req, err := http.NewRequest("DELETE", "/planet{id}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(Delete)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
}
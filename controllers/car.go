package controllers

import (
	//"../models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"v2/mongo/models"
)

type CarController struct {
	session *mgo.Session
}

var collection = "cars"
var dbName = "automobiles"

func NewCarController(s *mgo.Session) *CarController {
	return &CarController{s}
}

func (cc CarController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprintf(w, "Welcome to Future Automobiles")
}

func (cc CarController) CreateCar(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	car := models.Car{}

	// decode json
	json.NewDecoder(r.Body).Decode(&car)

	// generate an ID
	car.Id = bson.NewObjectId()

	// store car in mongo db
	cc.session.DB(dbName).C(collection).Insert(car)

	// convert car to json
	c, err := json.Marshal(car)

	// handle error
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s\n", c)


}
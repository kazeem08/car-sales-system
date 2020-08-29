package main

import (
	"../models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

type CarController struct {
	session *mgo.Session
}

func NewCarController(s *mgo.Session) *CarController {
	return &CarController{s}
}

func (cc CarController) CreateCar(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	c := models.Car{}

	// decode json
	json.NewDecoder(r.Body).Decode(&c)

}
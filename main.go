package main

import (
	"./controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)


func main() {
	r := httprouter.New()
	carController := controllers.NewCarController(getSession())
	r.GET("/", carController.CreateCar)

	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	// connect to local mongo
	s, err := mgo.Dial("localhost")

	// check if connection error, is mongo running
	if err != nil {
		panic(err)
	}
	return s
}


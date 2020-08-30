package main

import (
	//"./controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
	"v2/mongo/controllers"
)


func main() {
	r := httprouter.New()
	carController := controllers.NewCarController(getSession())
	r.GET("/", carController.Index)
	r.POST("/cars", carController.CreateCar)

	http.ListenAndServe("localhost:8080", r)
}

// Connect to mongoBb
func getSession() *mgo.Session {
	// connect to local mongo
	s, err := mgo.Dial("localhost")

	// check if connection error, is mongo running
	if err != nil {
		panic(err)
	}
	return s
}


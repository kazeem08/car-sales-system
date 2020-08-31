package main

import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"v2/mongo/controllers"
)


func TestGetCars(t *testing.T){
	//car := models.Car{
	//	Brand: "Toyota",
	//		Model: "Camry 2010",
	//		Color: "white",
	//		Door: 4,
	//		Condition: "used",
	//		Price: 2000000,
	//		Quantity: 10,
	//
	//}

		handler := controllers.NewCarController(getSession())
		router := httprouter.New()
		router.GET("/cars", handler.GetCars)

		req, _ := http.NewRequest("GET", "/cars", nil)
		rr := httptest.NewRecorder()

		fmt.Println(req.URL)

		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Wrong status")
		}
	}

func TestCreateCars(t *testing.T){
	car := `{"brand": "Toyota",
		"model": "Camry 2010",
		"color": "white",
		"door": 4,
		"Condition": "used",
		"price": 2000000,
		"quantity": 10}`

	//value, _ := json.Marshal(car)
	data := bytes.NewBufferString(string(car))

	handler := controllers.NewCarController(getSession())
	router := httprouter.New()
	router.POST("/cars", handler.GetCars)

	req, err := http.NewRequest("POST", "/cars", data)

	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()

	fmt.Println(rr.Code)

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}

//
func getSession() *mgo.Session {
	// connect to local mongo
	s, err := mgo.Dial("localhost")

	// check if connection error, is mongo running
	if err != nil {
		panic(err)
	}
	return s
}

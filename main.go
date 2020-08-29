package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type User struct {
	Name string
	Email string
	Sex string
}

var users = []User {
	{   Name: "kazeem",
		Email: "test@gmail.com",
		Sex: "Male",
	},
	{   Name: "Lanre",
		Email: "test2@gmail.com",
		Sex: "Female",
	},
}

func main() {
	r := httprouter.New()
	r.GET("/", GetUsers)

	http.ListenAndServe("localhost:8080", r)
}

func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	var result []User
	for _, user := range users {
		if user.Name == "kazeem" {
			result = append(result, user)
		}
	}

	uj, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}
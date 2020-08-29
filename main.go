package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)




func main() {
	r := httprouter.New()
	r.GET("/", GetUsers)

	http.ListenAndServe("localhost:8080", r)
}


package models

import "gopkg.in/mgo.v2/bson"

type Car struct {
	Name string `json:"name"`
	Brand string `json:"brand"`
	Door int `json:"door"`
	Condition string `json:"condition"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	Id bson.ObjectId `json:"id" bson: "_id">`
}
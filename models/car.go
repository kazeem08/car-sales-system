package models

import "gopkg.in/mgo.v2/bson"

type Car struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Door int `json:"door"`
	Color string `json:"color"`
	Condition string `json:"condition"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	Id bson.ObjectId `json:"id" bson: "_id">`
}
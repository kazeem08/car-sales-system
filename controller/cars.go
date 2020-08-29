package controller

import "gopkg.in/mgo.v2"

type CarsController struct {
	session *mgo.Session
}

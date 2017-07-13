package services

import (
	"gopkg.in/mgo.v2"
	"log"
)

var Session = mgo.Session{}

func InitMongo() bool {
	url := "mongodb://localhost"
	log.Println("Establishing MongoDB connection...")
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal("Cannot connect to MongoDB!")
		return true
	} else {
		Session = *session
		return false
	}
}

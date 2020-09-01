package main

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var lock sync.Mutex

//Defaultskip : Stores the value of the default offset
var Defaultskip = int64(0)

//Defaultlimit : Stores the value of the default limit
var Defaultlimit = int64(10)

var skip = Defaultskip
var limit = Defaultlimit

//participant : Stores the details of a participant
type participant struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Rsvp  string `json:"rsvp,omitempty" bson:"rsvp,omitempty"`
}

//Meeting : Stores the details of a meeting
type Meeting struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title,omitempty" bson:"title,omitempty"`
	Participants []participant      `json:"participants,omitempty" bson:"participants,omitempty"`
	Starttime    string             `json:"starttime,omitempty" bson:"starttime,omitempty"`
	Endtime      string             `json:"endtime,omitempty" bson:"endtime,omitempty"`
	Creationtime string             `json:"creationtime,omitempty" bson:"creationtime,omitempty"`
}

func (person *participant) cons() {
	if person.Rsvp == "" {
		person.Rsvp = "Not Answered"
	}
	if person.Email == "" {
		person.Email = "defaultmail@email.com"
	}
	if person.Name == "" {
		person.Name = person.Email
	}
}

func (obj *Meeting) def() {
	if obj.Title == "" {
		obj.Title = "Untitled Meeting"
	}
	if obj.Starttime == "" {
		obj.Starttime = string(time.Now().Format(time.RFC3339))
	}
	if obj.Endtime == "" {
		obj.Endtime = string(time.Now().Local().Add(time.Hour * time.Duration(1)).Format(time.RFC3339))
	}
	if obj.Creationtime == "" {
		obj.Creationtime = string(time.Now().Format(time.RFC3339))
	}
	for i := range obj.Participants {
		obj.Participants[i].cons()
	}
}

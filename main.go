package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type participant struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Rsvp  string `json:"rsvp,omitempty" bson:"rsvp,omitempty"`
}

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
		obj.Starttime = string(time.Now().Format("2006-01-02 15:04:05"))
	}
	if obj.Endtime == "" {
		obj.Endtime = string(time.Now().Local().Add(time.Hour * time.Duration(1)).Format("2006-01-02 15:04:05"))
	}
	if obj.Creationtime == "" {
		obj.Creationtime = string(time.Now().Format("2006-01-02 15:04:05"))
	}
	for i := range obj.Participants {
		obj.Participants[i].cons()
	}
}

/*
func CreateMeetingEndpoint(response http.ResponseWriter, request *http.Request) {}
func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request)     {}
func GetMeetingEndpoint(response http.ResponseWriter, request *http.Request)    {}
*/

func CreateMeetingEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Got message")
	fmt.Println(request.Method)
	if request.Method == "POST" {
		response.Header().Set("content-type", "application/json")
		var meet Meeting
		_ = json.NewDecoder(request.Body).Decode(&meet)
		meet.def()
		// fmt.Println(meet)
		collection := client.Database("appointy").Collection("meetings")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		result, _ := collection.InsertOne(ctx, meet)
		meet.ID = result.InsertedID.(primitive.ObjectID)
		// var res Meeting
		json.NewEncoder(response).Encode(meet)
		fmt.Println(meet)

	}
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Got message")
	fmt.Println(request.Method)
	if request.Method == "GET" {
		response.Header().Set("content-type", "application/json")
		fmt.Println((request.URL.Query()["id"][0]))
		id, _ := primitive.ObjectIDFromHex(request.URL.Query()["id"][0])
		var meet Meeting
		collection := client.Database("appointy").Collection("meetings")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		err := collection.FindOne(ctx, Meeting{ID: id}).Decode(&meet)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(meet)
	}
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://arnav:arnav0512@cluster0.l3dls.mongodb.net/<dbname>?retryWrites=true&w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)

	http.HandleFunc("/meetings", CreateMeetingEndpoint)
	// http.HandleFunc("/people", GetPersonEndpoint)
	http.HandleFunc("/meeting/", GetPersonEndpoint)
	http.ListenAndServe(":12345", nil)
}

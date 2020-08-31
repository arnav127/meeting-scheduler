package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
func CreateMeetingEndpoint(response http.ResponseWriter, request *http.Request) {}
func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request)     {}
func GetMeetingEndpoint(response http.ResponseWriter, request *http.Request)    {}
*/

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://arnav:arnav0512@cluster0.l3dls.mongodb.net/<dbname>?retryWrites=true&w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)

	http.HandleFunc("/meetings", CreateMeeting)
	http.HandleFunc("/articles/", GetParticipants)
	// http.HandleFunc("/people", GetPersonEndpoint)
	http.HandleFunc("/meeting/", GetMeeting)
	http.ListenAndServe(":12345", nil)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CheckParticipant : Returns a list of active meetings of the person
func CheckParticipant(email string) []Meeting {
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Find()
	opts.SetSort(bson.D{{"starttime", 1}})
	cursor, _ := collection.Find(ctx, bson.D{
		{"participants.email", email},
	}, opts)
	var meetingsreturn []Meeting
	var meet Meeting
	for cursor.Next(ctx) {
		cursor.Decode(&meet)
		meetingsreturn = append(meetingsreturn, meet)
	}
	return meetingsreturn
}

// GetParticipants : Gets a list of active meetings of the person
func GetParticipants(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		response.Header().Set("content-type", "application/json")
		fmt.Println((request.URL.Query()["participant"][0]))
		email := request.URL.Query()["participant"][0]
		participantmeetings := CheckParticipant(email)
		json.NewEncoder(response).Encode(participantmeetings)
	}
}

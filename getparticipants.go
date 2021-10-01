package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	opts.SetSort(bson.D{{Key: "starttime", Value: 1}})
	opts.Skip = &skip
	opts.Limit = &limit
	cursor, _ := collection.Find(ctx, bson.D{
		{Key: "participants.email", Value: email},
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
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
		return
	}
	response.Header().Set("content-type", "application/json")
	fmt.Println((request.URL.Query()["participant"][0]))
	if len(request.URL.Query()["limit"]) != 0 {
		limit, _ = strconv.ParseInt(request.URL.Query()["limit"][0], 0, 64)
	}
	if len(request.URL.Query()["ofset"]) != 0 {
		skip, _ = strconv.ParseInt(request.URL.Query()["offset"][0], 0, 64)
	}
	email := request.URL.Query()["participant"][0]
	participantmeetings := CheckParticipant(email)
	if len(participantmeetings) == 0 {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Participant not present" }`))
		return
	}
	json.NewEncoder(response).Encode(participantmeetings)
	skip = Defaultskip
	limit = Defaultlimit
}

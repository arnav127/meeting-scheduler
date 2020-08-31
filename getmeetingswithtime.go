package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//GetMeetingwithTime : Gives the meeting with the provided id
func GetMeetingwithTime(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	fmt.Println((request.URL.Query()["start"][0]))
	fmt.Println((request.URL.Query()["end"][0]))
	CheckstartTime := request.URL.Query()["start"][0]
	CheckendTime := request.URL.Query()["end"][0]
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, _ := collection.Find(ctx, bson.M{})
	var meetingsreturn []Meeting
	var meet Meeting
	for cursor.Next(ctx) {
		cursor.Decode(&meet)
		if (CheckstartTime <= meet.Starttime) && (CheckendTime >= meet.Endtime) {
			meetingsreturn = append(meetingsreturn, meet)
		}
	}
	json.NewEncoder(response).Encode(meetingsreturn)

}

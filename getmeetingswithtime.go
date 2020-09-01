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

//CheckMeetingwithTime : Returns the meetings within the time
func CheckMeetingwithTime(CheckStartTime string, CheckEndTime string) []Meeting {
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Find()
	opts.SetSort(bson.D{{"starttime", 1}})
	filter := bson.D{
		{"starttime", bson.M{"$gt": CheckStartTime}},
		{"endtime", bson.M{"$lt": CheckEndTime}},
	}
	cursor, _ := collection.Find(ctx, filter, opts)
	var meetingsreturn []Meeting
	var meet Meeting
	for cursor.Next(ctx) {
		cursor.Decode(&meet)
		meetingsreturn = append(meetingsreturn, meet)
	}
	return meetingsreturn
}

//GetMeetingwithTime : Gives the meetings within the time
func GetMeetingwithTime(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	fmt.Println((request.URL.Query()["start"][0]))
	fmt.Println((request.URL.Query()["end"][0]))
	CheckStartTime := request.URL.Query()["start"][0]
	CheckEndTime := request.URL.Query()["end"][0]
	meetingswithtime := CheckMeetingwithTime(CheckStartTime, CheckEndTime)
	json.NewEncoder(response).Encode(meetingswithtime)

}

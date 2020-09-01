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

//CheckMeetingwithTime : Returns the meetings within the time
func CheckMeetingwithTime(CheckStartTime string, CheckEndTime string) []Meeting {
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "starttime", Value: 1}})
	opts.Skip = &skip
	opts.Limit = &limit
	filter := bson.D{
		{Key: "starttime", Value: bson.M{"$gt": CheckStartTime}},
		{Key: "endtime", Value: bson.M{"$lt": CheckEndTime}},
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
	if len(request.URL.Query()["limit"]) != 0 {
		limit, _ = strconv.ParseInt(request.URL.Query()["limit"][0], 0, 64)
	}
	if len(request.URL.Query()["ofset"]) != 0 {
		skip, _ = strconv.ParseInt(request.URL.Query()["offset"][0], 0, 64)
	}
	meetingswithtime := CheckMeetingwithTime(CheckStartTime, CheckEndTime)
	json.NewEncoder(response).Encode(meetingswithtime)
	skip = Defaultskip
	limit = Defaultlimit
}

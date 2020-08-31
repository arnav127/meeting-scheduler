package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CheckMeetingwithID : Checks the meeting with the provided id
func CheckMeetingwithID(id primitive.ObjectID) (Meeting, error) {
	var meet Meeting
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Meeting{ID: id}).Decode(&meet)
	return meet, err
}

//GetMeetingwithID : Gives the meeting with the provided id
func GetMeetingwithID(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		response.Header().Set("content-type", "application/json")
		fmt.Println((request.URL.Query()["id"][0]))
		id, _ := primitive.ObjectIDFromHex(request.URL.Query()["id"][0])
		meetingwithID, err := CheckMeetingwithID(id)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(meetingwithID)
	}
}

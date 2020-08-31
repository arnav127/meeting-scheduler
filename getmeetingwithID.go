package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetMeetingwithID : Gives the meeting with the provided id
func GetMeetingwithID(response http.ResponseWriter, request *http.Request) {
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

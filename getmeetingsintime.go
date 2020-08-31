package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//GetMeetingsinTime : Returns the meetings that are currently going on
func GetMeetingsinTime(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		response.Header().Set("content-type", "application/json")
		fmt.Println((request.URL.Query()["participant"][0]))
		email := request.URL.Query()["participant"][0]
		collection := client.Database("appointy").Collection("meetings")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		cursor, _ := collection.Find(ctx, bson.M{})
		var meetingsreturn []Meeting
		var meet Meeting
		for cursor.Next(ctx) {
			cursor.Decode(&meet)
			if iswithintime(meet) {
				for _, person := range meet.Participants {
					if person.Email == email {
						meetingsreturn = append(meetingsreturn, meet)
						break
					}
				}
			}
		}
		json.NewEncoder(response).Encode(meetingsreturn)

	}
}

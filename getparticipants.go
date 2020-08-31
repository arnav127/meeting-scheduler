package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetParticipants(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		response.Header().Set("content-type", "application/json")
		fmt.Println((request.URL.Query()["participant"][0]))
		email := request.URL.Query()["participant"][0]
		collection := client.Database("appointy").Collection("meetings")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		cursor, _ := collection.Find(ctx, bson.M{})
		var meetingsreturn []Meeting
		for cursor.Next(ctx) {
			var meet Meeting
			cursor.Decode(&meet)
			for _, person := range meet.Participants {
				if person.Email == email {
					meetingsreturn = append(meetingsreturn, meet)
					break
				}
			}
		}
		json.NewEncoder(response).Encode(meetingsreturn)

	}
}

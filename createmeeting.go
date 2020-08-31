package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMeeting(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		response.Header().Set("content-type", "application/json")
		var meet Meeting
		_ = json.NewDecoder(request.Body).Decode(&meet)
		meet.def()
		collection := client.Database("appointy").Collection("meetings")
		// people := client.Database("appointy").Collection("people")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		result, _ := collection.InsertOne(ctx, meet)
		meet.ID = result.InsertedID.(primitive.ObjectID)
		json.NewEncoder(response).Encode(meet)
		fmt.Println(meet)

	}
}

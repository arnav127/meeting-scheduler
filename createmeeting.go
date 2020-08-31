package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateMeeting : Adds another meeting to the database
func CreateMeeting(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var meet Meeting
	_ = json.NewDecoder(request.Body).Decode(&meet)
	meet.def()
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, meet)
	meet.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(response).Encode(meet)
	fmt.Println(meet)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParticipantsBusy(thismeet Meeting) bool {
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, _ := collection.Find(ctx, bson.M{})
	var meet Meeting
	for cursor.Next(ctx) {
		cursor.Decode(&meet)
		if thismeet.Starttime <= meet.Endtime ||
			thismeet.Endtime >= meet.Starttime {
			for _, person := range meet.Participants {
				for _, thisperson := range thismeet.Participants {
					if thisperson == person {
						return true
					}
				}
			}
		}
	}
	return false
}

//CreateMeeting : Adds another meeting to the database
func CreateMeeting(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var meet Meeting
	_ = json.NewDecoder(request.Body).Decode(&meet)
	if ParticipantsBusy(meet) {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "Participants RSVP clash" }`))
		return
	}
	if meet.Starttime < meet.Creationtime {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "Meeting cannot start in the past" }`))
		return
	}
	meet.def()
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, meet)
	meet.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(response).Encode(meet)
	fmt.Println(meet)
}

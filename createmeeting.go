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

//ParticipantsBusy : Checks if the participants are not RSVP in any other meeting during this time
func ParticipantsBusy(thismeet Meeting) bool {
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var meet Meeting
	for _, thisperson := range thismeet.Participants {
		if thisperson.Rsvp == "Yes" {
			filter := bson.M{
				"participants.email": thisperson.Email,
				"participants.rsvp":  "Yes",
				"endtime":            bson.M{"$gt": string(time.Now().Format(time.RFC3339))},
			}
			cursor, _ := collection.Find(ctx, filter)
			for cursor.Next(ctx) {
				cursor.Decode(&meet)
				fmt.Println(meet, thismeet)
				if (thismeet.Starttime >= meet.Starttime && thismeet.Starttime <= meet.Endtime) ||
					(thismeet.Endtime >= meet.Starttime && thismeet.Endtime <= meet.Endtime) {
					return true
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
	meet.def()
	if meet.Starttime < meet.Creationtime {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "Meeting cannot start in the past" }`))
		return
	}
	if meet.Starttime > meet.Endtime {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "Invalid time" }`))
		return
	}
	lock.Lock()
	defer lock.Unlock()
	if ParticipantsBusy(meet) {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "Participants RSVP clash" }`))
		return
	}
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, meet)
	meet.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(response).Encode(meet)
	fmt.Println(meet)
}

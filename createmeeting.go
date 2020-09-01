package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ParticipantsBusy : Checks if the participants are not RSVP in any other meeting during this time
func ParticipantsBusy(thismeet Meeting) error {
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
				if (thismeet.Starttime >= meet.Starttime && thismeet.Starttime <= meet.Endtime) ||
					(thismeet.Endtime >= meet.Starttime && thismeet.Endtime <= meet.Endtime) {
					returnerror := "Error 400: Participant " + thisperson.Name + " RSVP Clash"
					return errors.New(returnerror)
				}
			}
		}
	}
	return nil
}

//CreateMeeting : Adds another meeting to the database
func CreateMeeting(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var meet Meeting
	_ = json.NewDecoder(request.Body).Decode(&meet)
	meet.def()
	if meet.Starttime < meet.Creationtime {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Meeting cannot start in the past" }`))
		return
	}
	if meet.Starttime > meet.Endtime {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Invalid time" }`))
		return
	}
	lock.Lock()
	defer lock.Unlock()
	err := ParticipantsBusy(meet)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
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

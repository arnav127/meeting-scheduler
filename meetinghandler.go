package main

import "net/http"

//MeetingHandler : Function to handle "/meetings" and call CreateMeeting or GetMeetingID accoording to request
func MeetingHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		CreateMeeting(response, request)
	}
	if request.Method == "GET" {
		GetMeetingwithTime(response, request)
	}
}

<p align="center">
   <img height="200" src="images/icon_transparent.png" >
</p>


# Meeting Scheduler  [![Go Report Card](https://goreportcard.com/badge/arnav127/meeting-scheduler)](https://goreportcard.com/report/arnav127/meeting-scheduler) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)
Meeting Scheduler is a basic version of HTTP JSON API written in Golang and works on the port 12345 by default.

### Meeting Attributes
Meetings have the following Attributes.
- Id
- Title
- Participants
- Start Time
- End Time
- Creation Timestamp  
*The format of the timestamp is RFC3339, eg: `2020-12-01T14:11:57+05:30`*  <br><br>

Participants have the following Attributes. 
- Name
- Email
- RSVP (Yes/No/MayBe/Not Answered)

### Supported Operations:
- **Schedule a meeting** <br> 
A POST request using the URL `'/meetings'` which returns the meeting in JSON format
- **Get a meeting using id** <br>
A GET request using the URL `'/meeting/<id here>'` which returns the meeting in JSON format
- **List all meetings within a time frame** <br>
A GET request using the URL `‘/meetings?start=<start time here>&end=<end time here>’` which returns an array of meetings in JSON format that are within the time range
- **List all meetings of a participant** <br>
A GET request using the URL `‘/articles?participant=<email id>’` which returns an array of meetings in JSON format that have the participant

### Other Features
- Support for pagination using limit and offset
- Thread Safe creation of meetings using Mutex Lock
- The response from the API is ordered by the start time of the meetings
- Multiple checks on the time of the meeting
- Meetings of a person at same time with RSVP Yes are not allowed

## Benchmarks
* POST New Meetings:
```
Running tool: /usr/bin/go test -benchmem -run=^$ -bench ^(BenchmarkMaingetpost)$
goos: linux
goarch: amd64
BenchmarkMaingetpost-8   	    1206	    855918 ns/op	   17942 B/op	     131 allocs/op
PASS
ok  	_/home/arnav/prog/appointy/meeting-scheduler	2.109s
```
* GET meeting from ID:
```
Running tool: /usr/bin/go test -benchmem -run=^$ -bench ^(BenchmarkMaingetmeet)$
goos: linux
goarch: amd64
BenchmarkMaingetmeet-8   	    1309	    826992 ns/op	   16761 B/op	     119 allocs/op
PASS
ok  	_/home/arnav/prog/appointy/meeting-scheduler	2.160s
```
* GET participants meetings:
```
Running tool: /usr/bin/go test -benchmem -run=^$ -bench ^(BenchmarkMaingetparticipant)$
goos: linux
goarch: amd64
BenchmarkMaingetparticipant-8   	    1216	    862029 ns/op	   16787 B/op	     120 allocs/op
PASS
ok  	_/home/arnav/prog/appointy/meeting-scheduler	2.152s
```
* GET meeting during time
```
Running tool: /usr/bin/go test -benchmem -run=^$ -bench ^(BenchmarkMaingettime)$
goos: linux
goarch: amd64
BenchmarkMaingettime-8   	    1202	    964725 ns/op	   16820 B/op	     120 allocs/op
PASS
ok  	_/home/arnav/prog/appointy/meeting-scheduler	2.189s
```

## Sample Responses
* GET participants meetings
```
curl "http://localhost:12345/articles/?participant=arnavdixit127@gmail.com&offset=1&limit=1" | json_pp
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   261  100   261    0     0   254k      0 --:--:-- --:--:-- --:--:--  254k
[
   {
      "_id" : "5f4e3d6ce7a896b6576a04bb",
      "creationtime" : "2020-09-01T17:54:12+05:30",
      "endtime" : "2021-09-01T10:52:12+05:30",
      "participants" : [
         {
            "email" : "arnavdixit127@gmail.com",
            "name" : "Arnav Dixit",
            "rsvp" : "Yes"
         }
      ],
      "starttime" : "2021-09-01T09:52:12+05:30",
      "title" : "Title"
   }
]
```
* GET meeting from ID
```
curl "http://localhost:12345/meeting/5f4e410c724c1867e9755e3b" | json_pp
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   319  100   319    0     0   155k      0 --:--:-- --:--:-- --:--:--  155k
{
   "_id" : "5f4e410c724c1867e9755e3b",
   "creationtime" : "2020-09-01T18:09:40+05:30",
   "endtime" : "2020-19-01T15:11:57+05:30",
   "participants" : [
      {
         "email" : "arnav@gmail.com",
         "name" : "Arnavvvvvvvvvv",
         "rsvp" : "Yes"
      },
      {
         "email" : "manast95@gmail.com",
         "name" : "Sudhanshu",
         "rsvp" : "No"
      }
   ],
   "starttime" : "2020-12-01T14:11:57+05:30",
   "title" : "The Meet"
}
```

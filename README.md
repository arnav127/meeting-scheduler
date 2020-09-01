# Meeting Scheduler
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

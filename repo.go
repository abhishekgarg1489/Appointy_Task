package main

import "fmt"

var currentId int

var meetings Meetings

// Give us some seed data
func init() {
	Create_meeting(Meeting{Title: "Let's Meet"})
	Create_meeting(Meeting{Title: "Host Meetup"})
}

func Find_meeting(id int) Meeting {
	for _, m := range meeting {
		if m.Id == id {
			return m
		}
	}
	return Meeting{}
}

func Create_meeting(m Meeting) Meeting {
	currentId += 1
	m.Id = currentId
	meetings = append(meetings, m)
	return m
}

func Destroy_meeting(id int) error {
	for i, m := range meetings {
		if m.Id == id {
			meetings = append(meetings[:i], meetings[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find meeting with id of %d to delete", id)
}

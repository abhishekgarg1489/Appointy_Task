package main

import "time"

type Meeting struct {
	Id           string    `json:id`
	Title        string    `json:"Title"`
	participants int       `json:"participants"`
	start_time   time.Time `json:"start"`
	end_time     time.Time `json:"end"`
}

var Meetings []Meeting

/*
type participants struct {
	Name  string `json:"name"`
	email string `json:"email"`
	rsvp  bool   `json:"rsvp"`
}
*/

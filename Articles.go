package main

import "net/http"

type Article struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Articles []Article

var articles = Articles{
	Article{
		"Index",
		"GET",
		"/",
		Index,
	},
	Article{
		"MeetingIndex",
		"GET",
		"/meetings",
		MeetingIndex,
	},
	Article{
		"MeetingCreate",
		"POST",
		"/meetings",
		MeetingCreate,
	},
	Article{
		"MeetingShow",
		"GET",
		"/meetings/{id}",
		MeetingShow,
	},
}

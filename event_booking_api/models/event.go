package models

import "time"

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	// add it to db
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}

package models

import (
	"fmt"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required` // these are called "Struct Text" and say "gin.ShouldBindJSON()" package they should be inside json body of request
	Description string    `binding:"required`
	Location    string    `binding:"required`
	DateTime    time.Time `binding:"required`
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
	fmt.Println(e)
}

func GetAllEvents() []Event {
	return events
}

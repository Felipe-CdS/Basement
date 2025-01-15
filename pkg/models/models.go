package models

import (
	"errors"
	"time"
)

var ErrNotFound = errors.New("models: no matching record found")
var ErrDbOperation = errors.New("models: database operation failed")
var ErrActityTypeNotFound = errors.New("models: activity type doesnt exist")
var ErrNotFinished = errors.New("models: there is one or more activities not finished yet")
var ErrAlreadyExists = errors.New("models: entry already exists")
var ErrInvalidInsert = errors.New("models: something wrong with the insert data")

type TagEnum string

const (
	ActivityTag TagEnum = "activity"
	IdolTag     TagEnum = "idol"
)

type Activity struct {
	ID          int
	StartTime   time.Time
	EndTime     time.Time
	Title       string
	Description string

	// Age = EndTime - StartTime
	// Special Property that the DB returns on some queries
	Age []uint8

	Tags []Tag // For query with tags search
}

type ActivityDayOverview struct {
	Date     time.Time
	TotalAge []uint8 // Total seconds of that day
	TotalSec int     // Total seconds of that day
}

type Tag struct {
	ID   int
	Type TagEnum
	Name string
}

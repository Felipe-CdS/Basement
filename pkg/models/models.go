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

type TagEnum string

const (
	ActivityTag TagEnum = "activity"
	IdolTag     TagEnum = "idol"
)

type Activity struct {
	ID          int
	StartTime   time.Time
	EndTime     time.Time
	Description string
}

type Tag struct {
	ID   int
	Type TagEnum
	Name string
}

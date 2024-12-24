package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")
var ErrActityTypeNotFound = errors.New("models: activity type doesnt exist")
var ErrAlreadyExists = errors.New("models: entry already exists")

type ActivityEnum string

const (
	Study            ActivityEnum = "study"
	ProgrammingWork  ActivityEnum = "programming_work"
	ProgrammingHobby ActivityEnum = "programming_hobby"
	ReadStudy        ActivityEnum = "read_study"
	Garbage          ActivityEnum = "garbage"
	ReadFun          ActivityEnum = "read_fun"
	Korean           ActivityEnum = "korean"
)

type Activity struct {
	ID           int
	activityType ActivityEnum
	StartTime    time.Time
	EndTime      time.Time
}

type DayStats struct {
	ID               int
	Date             time.Time
	Study            int
	ProgrammingWork  int
	ProgrammingHobby int
	ReadStudy        int
	Garbage          int
	ReadFun          int
	Korean           int
}

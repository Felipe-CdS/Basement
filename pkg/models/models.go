package models

import (
	"errors"
	"time"
)

var errNoRecord = errors.New("models: no matching record found")

type ActivityEnum string

const (
	Study            ActivityEnum = "study"
	ProgrammingWork  ActivityEnum = "programming_work"
	ProgrammingHobby ActivityEnum = "programming_hobby"
	ReadStudy        ActivityEnum = "read_study"
	Garbage          ActivityEnum = "garbage"
	ReadFun          ActivityEnum = "read_fun"
)

type Activity struct {
	ID           int
	activityType ActivityEnum
	StartTime    time.Time
	EndTime      time.Time
}

type DayStats struct {
	ID               string
	Date             time.Time
	Study            int
	ProgrammingWork  int
	ProgrammingHobby int
	ReadStudy        int
	Garbage          int
	ReadFun          int
}

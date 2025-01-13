package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	"nugu.dev/basement/pkg/models"
	"nugu.dev/basement/views/activity_views"
)

func (app *application) Activities(w http.ResponseWriter, r *http.Request) {

	// Create New Activity
	if r.Method == "POST" {
		app.startActivity(w, r)
		return
	}

	// Finish Open Activity
	if r.Method == "PATCH" {
		app.finishActivity(w, r)
		return
	}

	last, queryErr := app.activitiesRepository.GetLastActivity()

	if queryErr != nil && queryErr != models.ErrNotFound {
		http.Error(w, queryErr.Error(), http.StatusInternalServerError)
		return
	}

	listDone, queryErr := app.activitiesRepository.GetDailyLog(time.Now())

	if queryErr != nil && queryErr != models.ErrNotFound {
		http.Error(w, queryErr.Error(), http.StatusInternalServerError)
		return
	}

	var component templ.Component

	// if last activity doesnt exists or already ended, then delete cookie
	if queryErr == models.ErrNotFound || !last.EndTime.IsZero() {
		component = activity_views.ActivityIndex(false, listDone) // start button
	} else {
		component = activity_views.ActivityIndex(true, listDone) // stop button
	}

	setCookie(w, last.StartTime, last.EndTime)
	component.Render(r.Context(), w)
}

func (app *application) startActivity(w http.ResponseWriter, r *http.Request) {

	_, err := app.activitiesRepository.StartActivity()

	if err != nil {
		switch err {
		case models.ErrNotFinished:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	last, queryErr := app.activitiesRepository.GetLastActivity()

	if queryErr != nil && queryErr != models.ErrNotFound {
		http.Error(w, queryErr.Error(), http.StatusInternalServerError)
		return
	}

	setCookie(w, last.StartTime, last.EndTime)
	component := activity_views.StopButton()
	component.Render(r.Context(), w)
}

func (app *application) finishActivity(w http.ResponseWriter, r *http.Request) {

	err := app.activitiesRepository.EndActivity()
	if err != nil {
		switch err {
		case models.ErrNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	last, queryErr := app.activitiesRepository.GetLastActivity()

	if queryErr != nil && queryErr != models.ErrNotFound {
		http.Error(w, queryErr.Error(), http.StatusInternalServerError)
		return
	}

	setCookie(w, last.StartTime, last.EndTime)

	component := activity_views.StartButton()
	component.Render(r.Context(), w)
}

func (app *application) GetDailyLog(w http.ResponseWriter, r *http.Request) {

	dateReq, err := time.Parse(time.DateOnly, r.PathValue("date"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log, err := app.activitiesRepository.GetDailyLog(dateReq)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tags, err := app.tagsRepository.GetActivityTags()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	partialLog := activity_views.DetailedLog(dateReq, log, tags)

	reqType := r.URL.Query().Get("partial")
	if reqType == "true" {
		partialLog.Render(r.Context(), w)
		return
	}

	page := activity_views.Log(partialLog)
	page.Render(r.Context(), w)
}

// GET modal that creates a new daily log
func (app *application) NewDailyLog(w http.ResponseWriter, r *http.Request) {

	dateReq, err := time.Parse(time.DateOnly, r.PathValue("date"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log, err := app.activitiesRepository.GetDailyLog(dateReq)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tags, err := app.tagsRepository.GetActivityTags()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := activity_views.DetailedLog(dateReq, log, tags)
	component.Render(r.Context(), w)
}

// POST form to create new log
func (app *application) CreateDailyLog(w http.ResponseWriter, r *http.Request) {

	var from24ClockToTime = func(date string, hourMinTime string) (time.Time, error) {

		dateTime, err := time.Parse(time.DateOnly, date)

		if err != nil {
			return time.Time{}, err
		}

		holder := strings.Split(hourMinTime, ":")
		entryHours, err := strconv.Atoi(holder[0])

		if err != nil {
			return time.Time{}, err
		}

		entryMinutes, err := strconv.Atoi(holder[1])

		if err != nil {
			return time.Time{}, err
		}

		dateTime = dateTime.Add(time.Hour * time.Duration(entryHours))
		dateTime = dateTime.Add(time.Minute * time.Duration(entryMinutes))

		return dateTime, nil
	}

	r.ParseForm()

	var a models.Activity
	var err error

	a.Title = r.FormValue("title")
	a.Description = r.FormValue("description")
	a.StartTime, err = from24ClockToTime(r.FormValue("date"), r.FormValue("start"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.EndTime, err = from24ClockToTime(r.FormValue("date"), r.FormValue("end"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = app.activitiesRepository.NewCompleteActivity(a)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/log/%s", r.FormValue("date")), http.StatusSeeOther)
}

func setCookie(w http.ResponseWriter, startTime time.Time, endTime time.Time) {

	startTimerCookie := http.Cookie{Name: "start"}

	// endTime != null, last activity already finished. Kill the cookie
	if !endTime.IsZero() {
		startTimerCookie.Value = ""
		startTimerCookie.Expires = time.Now()
	} else {
		startTimerCookie.Value = strconv.FormatInt(startTime.Unix(), 10)
		startTimerCookie.Expires = time.Now().Add(time.Hour * 24)
	}

	http.SetCookie(w, &startTimerCookie)
}

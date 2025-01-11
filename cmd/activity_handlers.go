package main

import (
	"net/http"
	"strconv"
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

	component := activity_views.DetailedLog(dateReq, log)
	component.Render(r.Context(), w)
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

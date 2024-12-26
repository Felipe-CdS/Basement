package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"nugu.dev/basement/pkg/models"
	"nugu.dev/basement/views/activity_views"
)

func (app *application) Activities(w http.ResponseWriter, r *http.Request) {

	last, queryErr := app.activitiesRepository.GetLastActivity()

	if queryErr != nil && queryErr != models.ErrNotFound {
		http.Error(w, queryErr.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(last)

	setCookie(w, last.StartTime, last.EndTime)

	// Create New Activity
	if r.Method == "POST" {
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
		component := activity_views.StopButton()
		component.Render(r.Context(), w)
		return
	}

	// Finish Open Activity
	if r.Method == "PATCH" {
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
		component := activity_views.StartButton()
		component.Render(r.Context(), w)
		return
	}

	var component templ.Component

	// if last activity doesnt exists or already ended, then delete cookie
	if queryErr == models.ErrNotFound || last.EndTime.IsZero() {
		component = activity_views.ActivityIndex(false) // start button
	} else {
		component = activity_views.ActivityIndex(true) // stop button
	}

	component.Render(r.Context(), w)
}

func setCookie(w http.ResponseWriter, startTime time.Time, endTime time.Time) {

	startTimerCookie := http.Cookie{Name: "start"}

	// endTime != null, last activity already finished. Kill the cookie
	if endTime.IsZero() {
		startTimerCookie.Value = ""
		startTimerCookie.Expires = time.Now()
	} else {
		startTimerCookie.Value = strconv.FormatInt(startTime.Unix(), 10)
		startTimerCookie.Expires = time.Now().Add(time.Hour * 24)
	}

	log.Println(startTimerCookie)

	http.SetCookie(w, &startTimerCookie)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"nugu.dev/basement/pkg/models"
	dashboard_view "nugu.dev/basement/views/dashboard"
	landing_view "nugu.dev/basement/views/landing"
	layouts_view "nugu.dev/basement/views/layouts"
)

func (app *application) Landing(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		_, err := app.activities.Insert(models.Study, time.Now())

		if err != nil {
			log.Fatalln(err)
		}
		w.Write([]byte("Starting Study..."))
		return
	}

	if r.Method == "PATCH" {
		duration, aType, err := app.activities.SetEnd(time.Now())

		if err != nil {
			log.Fatalln(err)
		}

		_, err = app.dayStats.Insert(time.Now(), duration, aType)

		if err != nil {
			log.Fatalln(err)
		}

		w.Write([]byte("Ending Study..."))
		return
	}

	list, err := app.dayStats.GetByYear(2024)

	if err != nil {
		log.Fatalln(err)
	}

	component := landing_view.LandingPageView(list)
	component.Render(r.Context(), w)
}

func (app *application) Dashboard(w http.ResponseWriter, r *http.Request) {

	authToken, err := r.Cookie("t")

	if err != nil || authToken.Value != app.AuthToken {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	list, err := app.dayStats.GetByYear(2024)

	if err != nil {
		log.Fatalln(err)
	}

	component := dashboard_view.Dashboard(list)
	component.Render(r.Context(), w)
}

func (app *application) AddStatTime(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" && r.Method != "PATCH" {
		w.Header().Set("Allow", "POST PATCH")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	yyyy, _ := strconv.Atoi(r.PostForm.Get("year"))
	mm, _ := strconv.Atoi(r.PostForm.Get("month"))
	dd, _ := strconv.Atoi(r.PostForm.Get("day"))

	date := time.Date(yyyy, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)

	if r.Method == "POST" {
		stat, err := app.dayStats.CreateEmpty(date)

		if err == models.ErrAlreadyExists {
			w.Header().Set("HX-Reswap", "innerHTML")
			http.Error(w, "Already Exists", http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		component := dashboard_view.SingleLine(stat)
		component.Render(r.Context(), w)
		return
	}

	if r.Method == "PATCH" {

		t := models.ActivityEnum(r.PostForm.Get("type"))
		h, _ := strconv.Atoi(r.PostForm.Get("hours"))
		m, _ := strconv.Atoi(r.PostForm.Get("minutes"))
		s, _ := strconv.Atoi(r.PostForm.Get("seconds"))

		totalTime := h*3600 + m*60 + s

		stat, err := app.dayStats.Insert(date, totalTime, t)

		if err != nil {
			w.Header().Set("HX-Reswap", "innerHTML")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		w.Header().Set("HX-Retarget", fmt.Sprintf("#row-%s", date.Format(time.DateOnly)))
		component := dashboard_view.SingleLine(stat)
		component.Render(r.Context(), w)
	}

}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseForm()

		if r.FormValue("p") != app.AuthToken {
			w.Write([]byte("Wrong"))
			return
		}

		var redirectURL string
		refererSplit := strings.Split(r.Header.Get("Referer"), "redirect=")

		if len(refererSplit) > 1 {
			//redirect query parameter exists
			redirectURL = refererSplit[1]
		} else {
			//redirect query parameter doesnt exist, go /home
			redirectURL = "/"
		}

		expiration := time.Now().Add(time.Hour)
		cookie := http.Cookie{Name: "t", Value: app.AuthToken, Expires: expiration}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, redirectURL, http.StatusFound)
		return
	}

	if r.Method == "GET" {
		component := layouts_view.Login()
		component.Render(r.Context(), w)
		return
	}
}

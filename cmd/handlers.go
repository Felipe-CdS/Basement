package main

import (
	"net/http"
	"strings"
	"time"

	layouts_view "nugu.dev/basement/views/layouts"
)

func (app *application) Landing(w http.ResponseWriter, r *http.Request) {

	// if r.Method == "POST" {
	// 	_, err := app.activities.Insert(models.Study, time.Now())
	//
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	w.Write([]byte("Starting Study..."))
	// 	return
	// }
	//
	// if r.Method == "PATCH" {
	// 	duration, aType, err := app.activities.SetEnd(time.Now())
	//
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	//
	// 	_, err = app.dayStats.Insert(time.Now(), duration, aType)
	//
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	//
	// 	w.Write([]byte("Ending Study..."))
	// 	return
	// }
	//
	// list, err := app.dayStats.GetByYear(2024)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	component := layouts_view.StaticHome()
	component.Render(r.Context(), w)
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

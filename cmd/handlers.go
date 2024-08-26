package main

import (
	"log"
	"net/http"
	"time"

	"nugu.dev/basement/pkg/models"
	"nugu.dev/basement/views"
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

	component := views.LandingPageView(list)
	component.Render(r.Context(), w)
}

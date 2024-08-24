package main

import (
	"net/http"
	"nugu.dev/basement/views"
)

func Landing(w http.ResponseWriter, r *http.Request) {

	component := views.LandingPageView()

	component.Render(r.Context(), w)
}

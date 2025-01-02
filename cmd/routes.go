package main

import (
	"net/http"

	landing_view "nugu.dev/basement/views/landing"
)

func (a *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.Landing)

	mux.HandleFunc("/activities", func(w http.ResponseWriter, r *http.Request) {

		if err := a.AuthMiddleware(w, r); err != nil {
			return
		}

		a.Activities(w, r)
	})

	mux.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {

		if err := a.AuthMiddleware(w, r); err != nil {
			return
		}

		a.Gallery(w, r)
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		// not logged in, render login page
		if err := a.AuthMiddleware(w, r); err != nil {
			a.Login(w, r)
			return
		}

		// logged in trying to see /login page, redirect to parameter or /home
		var redirectURL string
		if redirectURL = r.URL.Query().Get("redirect"); redirectURL == "" {
			redirectURL = "/"
		}

		http.Redirect(w, r, redirectURL, http.StatusFound)
	})

	mux.HandleFunc("/reads", func(w http.ResponseWriter, r *http.Request) {

		if err := a.HtmxMiddleware(w, r); err != nil {
			return
		}

		component := landing_view.Reads()
		component.Render(r.Context(), w)
	})

	mux.HandleFunc("/bookmarks", func(w http.ResponseWriter, r *http.Request) {

		if err := a.HtmxMiddleware(w, r); err != nil {
			return
		}

		component := landing_view.Bookmarks()
		component.Render(r.Context(), w)
	})

	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	return mux
}

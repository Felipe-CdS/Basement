package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) AuthMiddleware(w http.ResponseWriter, r *http.Request) error {

	authToken, err := r.Cookie("t")

	if err != nil || authToken.Value != app.AuthToken {

		if r.URL.Path == "/login" {
			return fmt.Errorf(http.StatusText(http.StatusPermanentRedirect))
		}

		redir := fmt.Sprintf("/login?redirect=%s", r.URL.Path)
		http.Redirect(w, r, redir, http.StatusPermanentRedirect)
		return fmt.Errorf(http.StatusText(http.StatusPermanentRedirect))
	}

	return nil
}

func (app *application) HtmxMiddleware(w http.ResponseWriter, r *http.Request) error {

	htmxReq, err := strconv.ParseBool(r.Header.Get("HX-Request"))

	if err != nil || !htmxReq {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return fmt.Errorf(http.StatusText(http.StatusPermanentRedirect))
	}

	return nil
}

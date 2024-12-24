package main

import "net/http"

func (a *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.Landing)
	mux.HandleFunc("/dashboard", a.Dashboard)

	mux.HandleFunc("/stat", a.AddStatTime)

	mux.HandleFunc("/gallery", a.Gallery)

	mux.HandleFunc("/login", a.Login)

	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	return mux
}

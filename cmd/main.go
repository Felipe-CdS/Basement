package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", Landing)
	fmt.Println("Server Working...")
	http.ListenAndServe(":3000", nil)
}

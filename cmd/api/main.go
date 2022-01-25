package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jrigoin/lenslocked/controllers"
)

const (
	GET  = "GET"
	POST = "POST"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Not Found")
}

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods(GET)
	r.Handle("/contact", staticC.Contact).Methods(GET)
	r.HandleFunc("/signup", usersC.New).Methods(GET)
	r.HandleFunc("/signup", usersC.Create).Methods(POST)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

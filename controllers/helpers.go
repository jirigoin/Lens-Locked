package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	// r.PostForm = map[string][]string  -> with standard library
	// fmt.Fprintln(w, r.PostForm["password"])
	// fmt.Fprintln(w, r.PostForm["email"])

	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}

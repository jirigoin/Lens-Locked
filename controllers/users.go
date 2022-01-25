package controllers

import (
	"fmt"
	"lenslocked/views"
	"net/http"

	"github.com/gorilla/schema"
)

// NewUsers is used to create a new Users controller.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	// r.PostForm = map[string][]string  -> with standard library
	// fmt.Fprintln(w, r.PostForm["password"])
	// fmt.Fprintln(w, r.PostForm["email"])
	dec := schema.NewDecoder()
	var form SignupForm
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}

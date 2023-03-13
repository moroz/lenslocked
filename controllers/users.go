package controllers

import (
	"fmt"
	"net/http"

	"github.com/moroz/lenslocked/models"
)

const COOKIE_KEY = "_lenslocked_session"

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

func signUserIn(w http.ResponseWriter, cookieKey string, user *models.User) {
	if user == nil {
		panic("Attempted to sign in with nil user")
	}
	token, _ := models.IssueTokenForUserID(user.ID)
	cookie := http.Cookie{
		Name:     cookieKey,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, r, nil)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	u.Templates.SignIn.Execute(w, r, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	signUserIn(w, COOKIE_KEY, user)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Authenticate(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid email/password combination", http.StatusUnprocessableEntity)
		return
	}
	signUserIn(w, COOKIE_KEY, user)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Users) Profile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	fmt.Fprintf(w, "Current user: %s\n", user.Email)
}

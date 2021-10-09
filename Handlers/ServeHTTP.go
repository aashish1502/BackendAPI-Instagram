package Handlers

import (
	"fmt"
	"net/http"
	"regexp"
)

var (
	listUserRe = regexp.MustCompile(`^\/users[\/]*$`)
	getUserRe  = regexp.MustCompile(`^\/users\/([a-zA-Z0-9]+)$`)
	createUser = regexp.MustCompile(`^\/users[\/]*$`)
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	fmt.Println(r.URL.Path)
	switch {

	case r.Method == http.MethodGet && listUserRe.MatchString(r.URL.Path):
		GET(w, r)
		return
	case r.Method == http.MethodGet && getUserRe.MatchString(r.URL.Path):
		GETByID(w, r)
		return

	case r.Method == http.MethodPost && createUser.MatchString(r.URL.Path):
		CreateUser(w, r)
		return
	}

}

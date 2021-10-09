package Handlers

import (
	"net/http"
	"regexp"
	"sync"
)

var (
	listUserRe      = regexp.MustCompile(`^\/users[\/]*$`)
	listPostRe      = regexp.MustCompile(`^\/posts\/([a-zA-Z0-9]+)$`)
	getpostbyUserRe = regexp.MustCompile(`^\/posts\/users\/([a-zA-Z0-9]+)$`)
	getUserRe       = regexp.MustCompile(`^\/users\/([a-zA-Z0-9]+)$`)
	createUser      = regexp.MustCompile(`^\/users[\/]*$`)
	createPost      = regexp.MustCompile(`^\/posts[\/]*$`)
)

var lock sync.Mutex

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	//fmt.Println(r.URL.Path)

	lock.Lock()
	defer lock.Unlock()
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
	case r.Method == http.MethodGet && listPostRe.MatchString(r.URL.Path):
		getPostsByID(w, r)
		return
	case r.Method == http.MethodGet && getpostbyUserRe.MatchString(r.URL.Path):
		getByUser(w, r)
		return
	case r.Method == http.MethodPost && createPost.MatchString(r.URL.Path):
		CreatePost(w, r)
		return
	}

}

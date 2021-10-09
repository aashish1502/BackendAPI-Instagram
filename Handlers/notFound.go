package Handlers

import "net/http"

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{error: "Not found"}`))
}

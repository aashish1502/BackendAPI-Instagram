package Handlers

import (
	"BackendAPI-Instagram/DatabaseConnector"
	"encoding/json"
	"fmt"
	"net/http"
)

func getPostsByID(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("This is called!!")
	cliet, ctx := DatabaseConnector.ConnectDatabase()
	defer cliet.Disconnect(ctx)
	matches := listPostRe.FindStringSubmatch(r.URL.Path)
	fmt.Println(matches)
	post := DatabaseConnector.FetchPosts(cliet, ctx, matches[1])

	if post == nil {
		notFound(w, r)
	} else {
		data, err := json.Marshal(post)
		if err != nil {
			internalServerError(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}

}

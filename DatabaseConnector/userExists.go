package DatabaseConnector

import (
	"BackendAPI-Instagram/Structures"
)

func UserExists(post Structures.Post) bool {

	users := GetUsersByEmail(post.Email)
	if users == nil {
		return false
	} else {
		return true
	}

}

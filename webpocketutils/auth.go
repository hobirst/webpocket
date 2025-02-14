package webpocket

import (
	"net/http"
)

func checkAuth(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}

	if username == authuser && password == authpass {
		return true
	}

	return false
}

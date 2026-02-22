package webpocket

import (
	"net/http"
)

func checkAuth(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}

	if username == Authuser && password == Authpass {
		return true
	}

	return false
}

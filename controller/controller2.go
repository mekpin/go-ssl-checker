package controller

import (
	"fmt"
	"net/http"
)

func HandlerLogin2(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	fmt.Println(username)
	fmt.Println(password)

}

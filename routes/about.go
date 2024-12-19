package routes

import (
	"fmt"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "About page")
	if err != nil {
		return
	}
}

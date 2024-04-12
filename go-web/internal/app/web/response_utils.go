package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TextResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(status)
	fmt.Fprintf(w, message)
}

func JsonResponse(w http.ResponseWriter, status int, body interface{}) {
	jsonstr, err := json.Marshal(body)
	if err != nil {
		TextResponse(w, 500, "JSON marshal error")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, string(jsonstr))
}

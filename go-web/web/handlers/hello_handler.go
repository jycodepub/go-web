package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "Hello GO Web!")
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/healthcheck", healthcheck)
	http.ListenAndServe(":8080", nil)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

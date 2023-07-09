package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8880")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authorized!"))
	})

	http.ListenAndServe(":8880", nil)
}

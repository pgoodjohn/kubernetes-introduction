package main

import (
	"fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Recorded Visit to /")
    })

    err := http.ListenAndServe(":8000", nil)
    log.Fatal(err)
}


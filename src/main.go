package main

import (
	"fmt"
    "log"
    "net/http"
    "os"
)

func main() {

    fmt.Println("Started Application, listening on port 8080")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Recorded Visit to /")
    })

    http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)

        fmt.Println("Bye Bye");

        os.Exit(1)
    })

    err := http.ListenAndServe(":8080", nil)
    log.Fatal(err)
}


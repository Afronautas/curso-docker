package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go!")
	})

	serverRunningMessage := fmt.Sprintf("Running server on port %d", port)
	fmt.Println(serverRunningMessage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/send", sendLines)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func sendLines(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
	}

	log.Println(string(body))

	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		_, err := http.Post("http://localhost:8000/add-item", "text/plain", strings.NewReader(line))
		if err != nil {
			http.Error(w, "Can't send to queue", http.StatusInternalServerError)
		}
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Received: " + string(body)))
}

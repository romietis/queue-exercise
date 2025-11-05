package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var queue []string

func main() {
	http.HandleFunc("/add-item", addItem)
	http.HandleFunc("/get-item", getItem)

	fmt.Println("Server started at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}

	queue = append(queue, string(body))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Item added to the queue"))
	log.Println("Item added to the queue")
}

func getItem(w http.ResponseWriter, r *http.Request) {
	if len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(item))
		log.Printf("Item retrieved from the queue: %s", item)
	} else {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Queue is empty"))
		log.Println("Queue is empty")
	}
}

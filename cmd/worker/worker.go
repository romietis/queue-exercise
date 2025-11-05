package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Worker struct {
	QueueBaseUrl string
}

var defaultWorker = Worker{
	QueueBaseUrl: "http/:/localhost:8000",
}

func main() {

	http.HandleFunc("/send", defaultWorker.sendLines)
	http.HandleFunc("/receive", defaultWorker.receiveLines)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func (worker Worker) sendLines(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
	}

	log.Println(string(body))

	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		_, err := http.Post(worker.QueueBaseUrl+"/add-item", "text/plain", strings.NewReader(line))
		if err != nil {
			http.Error(w, "Can't send to queue", http.StatusInternalServerError)
		}
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Received: " + string(body)))
}

func (worker Worker) receiveLines(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(worker.QueueBaseUrl + "/get-item")
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Queue is empty"))
		// http.Error(w, err.Error(), http.StatusNoContent)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	writeFile(string(body), "output.txt") // add error handling
	w.WriteHeader(http.StatusOK)
}

func writeFile(line string, fileName string) error {
	if line == "" {
		return fmt.Errorf("empty line")
	}
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(line + "\n"); err != nil {
		return err
	}
	return nil
}

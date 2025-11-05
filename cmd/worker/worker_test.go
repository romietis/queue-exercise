package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSendLines(t *testing.T) {
	body := "abc\nbca\ncba"

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Item added"))
	}))
	defer mockServer.Close()

	request := httptest.NewRequest(http.MethodPost, "/send", strings.NewReader(body))
	recorder := httptest.NewRecorder()

	sendLines(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, recorder.Code)
	}

	wantResponseBody := fmt.Sprintf("Received: %s", body)

	if recorder.Body.String() != wantResponseBody {
		t.Fatalf("expected: %s, got: %s", wantResponseBody, recorder.Body.String())
	}
}

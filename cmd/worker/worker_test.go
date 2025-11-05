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
	requestCount := 0

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Item added"))
	}))
	defer mockServer.Close()

	testWorker := Worker{
		QueueBaseUrl: mockServer.URL,
	}

	request := httptest.NewRequest(http.MethodPost, "/send", strings.NewReader(body))
	recorder := httptest.NewRecorder()

	testWorker.sendLines(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, recorder.Code)
	}

	wantResponseBody := fmt.Sprintf("Received: %s", body)

	if recorder.Body.String() != wantResponseBody {
		t.Fatalf("expected: %s, got: %s", wantResponseBody, recorder.Body.String())
	}

	wantRequestCount := 3
	if requestCount != wantRequestCount {
		t.Fatalf("expexted %d request count, got %d", wantRequestCount, requestCount)
	}
}

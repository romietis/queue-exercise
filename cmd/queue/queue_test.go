package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddItem(t *testing.T) {
	queue = []string{}
	body := "some text"
	request := httptest.NewRequest(http.MethodPost, "/add-item", strings.NewReader(body))

	recorder := httptest.NewRecorder()

	addItem(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	wantResponseBody := "Item added to the queue"
	if recorder.Body.String() != wantResponseBody {
		t.Fatalf("expected: %s, got: %s", wantResponseBody, recorder.Body.String())
	}

	if len(queue) != 1 {
		t.Fatalf("expected queue length 1, got %d", len(queue))
	}

	if queue[0] != body {
		t.Fatalf("expected queue[0] to be %q, got %q", body, queue[0])
	}
}

// func TestAddItemJson(t *testing.T) {
// 	queue = []string{}

// 	request := httptest.NewRequest(http.MethodPost, "/add", nil)
// 	recorder := httptest.NewRecorder()

// 	addItem(recorder, request)

// 	if recorder.Code != http.StatusBadRequest {
// 		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
// 	}
// }

func TestGetItem(t *testing.T) {
	queue = []string{"abc"}

	request := httptest.NewRequest(http.MethodGet, "/get-item", nil)
	recorder := httptest.NewRecorder()

	getItem(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	wantResponseBody := "abc"
	if recorder.Body.String() != wantResponseBody {
		t.Fatalf("expected: %s, got: %s", wantResponseBody, recorder.Body.String())
	}
}

func TestGetItemEmptyQueue(t *testing.T) {
	queue = []string{}

	request := httptest.NewRequest(http.MethodGet, "/get-item", nil)
	recorder := httptest.NewRecorder()

	getItem(recorder, request)

	if recorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, recorder.Code)
	}

	wantResponseBody := "Queue is empty"
	if recorder.Body.String() != wantResponseBody {
		t.Fatalf("expected: %s, got: %s", wantResponseBody, recorder.Body.String())
	}
}

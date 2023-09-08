package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// response recorder
	rr := httptest.NewRecorder()

	// call index handler
	IndexHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: handler returned %v", status)
	}

	expected := "Hello, Suitmedia!"
	if rr.Body.String() != expected {
		t.Errorf("Error: unexpedted body: got %v, want %v", rr.Body, expected)
	}
}

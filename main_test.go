package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	HelloServer(res, req)

	exp := "!...Hello World...!"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s gog %s", exp, act)
	}
}

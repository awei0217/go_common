package basic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var personResponse = []PersonNew{
	{
		Name:    "wahaha",
		Address: "shanghai",
		Age:     20,
	},
	{
		Name:    "lebaishi",
		Address: "shanghai",
		Age:     10,
	},
}

var personResponseBytes, _ = json.Marshal(personResponse)

func TestPublishWrongResponseStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(personResponseBytes)
		if r.Method != "GET" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}
		if r.URL.EscapedPath() != "/person" {
			t.Errorf("Expected request to '/person', got '%s'", r.URL.EscapedPath())
		}
		r.ParseForm()
		topic := r.Form.Get("addr")
		if topic != "shanghai" {
			t.Errorf("Expected request to have 'addr=shanghai', got: '%s'", topic)
		}
	}))

	defer ts.Close()
	api := ts.URL
	fmt.Println("url:", api)
	resp, _ := GetInfo(api)

	fmt.Println("reps:", resp)
}

package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestPOST(t *testing.T) {
	postUrl := "localhost:8080/books"
	// Want to post a book
	t.Run("post a book", func(t *testing.T) {
		var jsonStr = []byte(`{
			"Name": "Either/Or",
			"Year": 1843,
			"Author": "Soren Kierkegaard",
			"Summary": "Either/Or portrays two life views. Each life view is written and represented by a fictional pseudonymous author, with the prose of the work reflecting and depending on the life view being discussed.",
			"Publisher": "Penguin Classics",
			"PageCount": 640,
			"ReadPage": 0,
			"Reading": true,
			"Finished": false,
		}`)

		req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Errorf("Cannot post book. Got %v", req)
		}
	})

	// Want to post a missing author book
	t.Run("post a missing author book", func(t *testing.T) {
		var jsonStr = []byte(`{
			"Name": "Dalang dibalik G30SPKI",
			"Year": 1990,
			"Author": "",
			"Summary": "Dalang",
			"Publisher": "Majalah Obor",
			"PageCount": 240,
			"ReadPage": 0,
			"Reading": false,
			"Finished": false,
		}`)
		req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonStr))
		if err == nil {
			t.Errorf("Should shown an error. Got %v", req)
		}
	})
}

// func TestGET(t *testing.T)    {}
// func TestPUT(t *testing.T)    {}
// func TestDELETE(t *testing.T) {}

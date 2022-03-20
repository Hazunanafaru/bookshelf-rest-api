package main

import "testing"

func TestPOST( t*testing.T) {
	t.Run("post a book", func(t *testing.T) {
		url := "localhost:8080/books"

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
    	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Errorf("cannot post book")
		}

		
		}
	})
}
func TestGET( t*testing.T) {}
func TestPUT( t*testing.T) {}
func TestDELETE( t*testing.T) {}
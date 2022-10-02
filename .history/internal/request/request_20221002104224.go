package request

import (
	"fmt"
	"log"
	"net/http"
)

func Get(url string) error {
	// Build fileName from fullPath
	getURL, err := url.Parse(url)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	// Put content on file
	resp, err := client.Get(getURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Downloaded a file %s with size %d", fileName, size)
	return err
}

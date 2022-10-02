package request

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Get(u string) error {
	// Build fileName from fullPath
	getURL, err := url.Parse(u)
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
	b, _ := ioutil.ReadAll(resp.Body)

	size := len(b)
	fmt.Printf("Fetched the url of a file %s with size %d", getURL, size)
	return err
}

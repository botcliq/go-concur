package request

import (
	"fmt"
	"go-concur/internal/csvlog"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func Get(u string, n int) error {
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
	reqTime := time.Now()
	// Put content on file
	resp, err := client.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	respTime := time.Now()
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	latency := respTime.Sub(reqTime)
	size := len(b)
	fmt.Printf("Fetched the url for task %d, url %s with size %d and latency %f \n", n, getURL, size, latency.Seconds())

	csvlog.LogCsv("result.csv", []string{string(n), u, string(size), reqTime.String(), respTime.String(), latency.String()})
	return err
}

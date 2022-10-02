package request

import (
	"crypto/tls"
	"fmt"
	"go-concur/internal/csvlog"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func Get(r Apis, n string) error {
	// Build fileName from fullPath
	getURL, err := url.Parse(r.Url)
	if err != nil {
		log.Fatal(err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr,
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	req, err := http.NewRequest(r.Type, r.Url, nil)
	for _, v := range r.Headers {
		req.Header.Add(v.Attr, v.Value)
	}
	reqTime := time.Now()
	resp, err := client.Do(req)
	defer resp.Body.Close()
	// Put content on file
	if err != nil {
		log.Fatal(err)
	}
	respTime := time.Now()
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	latency := respTime.Sub(reqTime)
	size := len(b)
	fmt.Printf("Fetched the url for task %s, url %s with size %d and latency %f \n", n, getURL, size, latency.Seconds())

	csvlog.LogCsv("result.csv", []string{n, r.Url, string(size), reqTime.String(), respTime.String(), latency.String()})
	return err
}

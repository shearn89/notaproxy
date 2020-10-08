/*
Simple web proxy to mask an API key from a client.
*/

package main

import (
	"net/http"
	"net/http/httputil"

	log "github.com/sirupsen/logrus"
)

const (
	scheme = "https"
	host   = "www.bbc.co.uk"
	//apiKey = "foobar"
)

// set up logger
func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

// Actual handler function
func Proxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = scheme
		req.URL.Host = host
		req.Host = host
		//req.Header.Add("X-ApiKey", apiKey)
		log.WithFields(log.Fields{
			"scheme":      req.URL.Scheme,
			"host":        req.URL.Host,
			"host-header": req.Host,
		}).Debug("proxying request")
	}
	return &httputil.ReverseProxy{Director: director}
}

func main() {
	proxy := Proxy()
	log.WithFields(log.Fields{
		"target_scheme": scheme,
		"target_host":   host,
	}).Info("starting proxy instance")
	log.Fatal(http.ListenAndServe(":80", proxy))
}

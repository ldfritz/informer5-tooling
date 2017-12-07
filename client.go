package informer5

import (
	"crypto/tls"
	"net/http"
)

func InsecureClient() *http.Client {
	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: t}
}

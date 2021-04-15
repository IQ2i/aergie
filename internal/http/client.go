package http

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os"
)

func CreateClient() *http.Client {
	insecureTLS := false
	if os.Getenv("AE_ENV") == "dev" {
		insecureTLS = true
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureTLS},
	}

	return &http.Client{Transport: tr}
}

func CreateRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func ExecRequest(client *http.Client, req *http.Request) (string, error) {
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

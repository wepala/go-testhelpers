package testhelpers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

//This is meant to make it easy to record request and responses so that they can be used later as test fixtures
func NewClientRecorder(fileName string, verbose bool) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {

			reqf, err := os.Create("./fixtures/" + fileName + "_request")
			respf, err := os.Create("./fixtures/" + fileName + "_response")
			if err == nil {
				//record request
				requestBytes, _ := httputil.DumpRequestOut(req, verbose)
				reqf.Write(requestBytes)

				//setup regular transport that does it's thing
				response, _ := http.DefaultTransport.RoundTrip(req)

				//record response
				responseBytes, _ := httputil.DumpResponse(response, verbose)
				respf.Write(responseBytes)

				return response

			}
			defer func() {
				reqf.Close()
				respf.Close()
				if r := recover(); r != nil {
					fmt.Println("Recording failed with errors", r)
				}
			}()

			return nil

		}),
	}
}

package testhelpers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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

			reqf, err := os.Create(fileName + "_request")
			respf, err := os.Create(fileName + "_response")
			if err == nil {
				//record request
				requestBytes, _ := httputil.DumpRequestOut(req, verbose)
				reqf.Write(requestBytes)

				//setup regular transport that does it's thing
				response, callError := http.DefaultTransport.RoundTrip(req)

				if callError != nil {
					respf.Write([]byte(callError.Error()))
					return nil
				} else {
					//record response
					responseBytes, _ := httputil.DumpResponse(response, verbose)
					respf.Write(responseBytes)

					return response
				}

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

func LoadHTTPRequestFixture(data []byte) (*http.Request, error) {
	reader := bufio.NewReader(bytes.NewReader(data))
	resp, err := http.ReadRequest(reader)
	if err == io.EOF {
		return resp, nil
	}
	return resp, err
}

func LoadHTTPResponseFixture(data []byte, req *http.Request) (*http.Response, error) {
	reader := bufio.NewReader(bytes.NewReader(data))
	resp, err := http.ReadResponse(reader, req)
	if err != nil {
		return nil, err
	}
	//save response body
	b := new(bytes.Buffer)
	io.Copy(b, resp.Body)
	resp.Body.Close()
	resp.Body = ioutil.NopCloser(b)

	return resp, err
}

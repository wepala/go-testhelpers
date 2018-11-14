package testhelpers

import (
	"io/ioutil"
	"testing"
)

func TestLoadResponseFixture(t *testing.T) {
	responseData := []byte(`HTTP/2.0 200 OK
Content-Length: 231
Content-Type: application/json;charset=UTF-8
Date: Wed, 14 Nov 2018 16:33:59 GMT

{"data":{"title":null}}`)
	response, err := LoadHTTPResponseFixture(responseData, nil)
	if err != nil {
		t.Error(err)
	} else {
		if response.StatusCode != 200 {
			t.Errorf("Expected status code 200 got: '%d'", response.StatusCode)
		}

		if response.Header.Get("Content-Type") != "application/json;charset=UTF-8" {
			t.Errorf("Expected Header Content-Type to have value 'application/json;charset=UTF-8' got: '%s'", response.Header.Get("Content-Type"))
		}

		if response.Header.Get("Date") != "Wed, 14 Nov 2018 16:33:59 GMT" {
			t.Errorf("Expected Header Date to have value 'application/json;charset=UTF-8' got: '%s'", response.Header.Get("Date"))
		}

		if response.ContentLength != 231 {
			t.Errorf("Expected content length 26 got: '%d'", response.ContentLength)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Error(err)
		}
		bodyString := string(bodyBytes)
		if bodyString != `{"data":{"title":null}}` {
			t.Errorf("Expected body to be '%s' got: '%s'", `{"data":{"title":null}}`, bodyString)
		}
	}

}

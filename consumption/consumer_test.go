package consumption

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func getDummyClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestDoStuffWithRoundTripper(t *testing.T) {
	dummyClient := getDummyClient(func(req *http.Request) *http.Response {
		// Test request parameters
		equals(t, req.URL.String(), "http://example.com")
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	request := NewConsumer(dummyClient)
	URL, _ := url.Parse("http://example.com")
	httpCode, body, err := request.Consume(URL)
	ok(t, err)
	equals(t, 200, httpCode)
	equals(t, []byte("OK"), body)

}

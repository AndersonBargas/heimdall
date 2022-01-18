package notification

import (
	"net/url"
	"time"
)

type EmailChannel struct {
}

func NewEmailChannel() *EmailChannel {
	return &EmailChannel{}
}

func (ec *EmailChannel) Notify(endpoint *url.URL, httpCode int, body []byte, begin time.Time, end time.Time) error {
	println("Deu FOOOM!!")
	println(endpoint.String())
	println(httpCode)
	println(string(body))
	println(begin.String())
	println(end.String())
	return nil
}

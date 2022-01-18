package notification

import (
	"net/url"
	"time"
)

type Notificator struct {
	channels []Channel
}

type Channel interface {
	Notify(endpoint *url.URL, httpCode int, body []byte, begin time.Time, end time.Time) error
}

func NewNotificator(channels ...Channel) *Notificator {
	return &Notificator{
		channels: channels,
	}
}

func (n *Notificator) Notify(endpoint *url.URL, httpCode int, body []byte, begin time.Time, end time.Time) error {
	for _, channel := range n.channels {
		if err := channel.Notify(endpoint, httpCode, body, begin, end); err != nil {
			return err
		}
	}
	return nil
}

package orchestrator

import (
	"net/url"
	"time"
)

type Orchestrator struct {
	consumer Consumer
	endpoint *url.URL
	notifier Notifier
}

type Consumer interface {
	Consume(endpoint *url.URL) (httpCode int, body []byte, err error)
}

type Notifier interface {
	Notify(endpoint *url.URL, httpCode int, body []byte, begin time.Time, end time.Time) error
}

func NewOrchestrator(consumer Consumer, endpoint *url.URL, notifier Notifier) *Orchestrator {
	return &Orchestrator{consumer, endpoint, notifier}
}

func (o *Orchestrator) Do() error {
	begin := time.Now()

	httpCode, body, err := o.consumer.Consume(o.endpoint)
	end := time.Now()
	if err != nil {
		return o.notifier.Notify(o.endpoint, httpCode, body, begin, end)
	}

	if httpCode < 200 || httpCode > 299 {
		return o.notifier.Notify(o.endpoint, httpCode, body, begin, end)
	}

	return nil
}

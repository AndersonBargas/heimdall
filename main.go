package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AndersonBargas/heimdall/config"
	"github.com/AndersonBargas/heimdall/consumption"
	"github.com/AndersonBargas/heimdall/notification"
	"github.com/AndersonBargas/heimdall/orchestrator"
)

func main() {

	c := config.NewConfigFromFlags()
	httpClient := &http.Client{Timeout: time.Second * 18000}

	consumer := consumption.NewConsumer(httpClient)

	notificator := notification.NewNotificator(
		notification.NewEmailChannel(),
	)

	if err := orchestrator.NewOrchestrator(consumer, c.URL(), notificator).Do(); err != nil {
		log.Fatalf(err.Error())
	}

	// request := core.NewRequest(&http.Client{
	// 	Timeout: c.Timeout(),
	// }, c.URL())

	// _, _, err := request.Call()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

}

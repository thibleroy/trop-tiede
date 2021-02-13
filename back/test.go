package main

import (
	"github.com/getsentry/sentry-go"
	"log"
	"time"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "http://6baa3172eb714473ba8c66e53818fec7@localhost:9000/2",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("Done")
}

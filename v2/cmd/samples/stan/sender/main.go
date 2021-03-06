package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cestan "github.com/cloudevents/sdk-go/v2/protocol/stan"
)

func main() {
	s, err := cestan.NewSender("test-cluster", "test-client", "test-subject", cestan.StanOptions())
	if err != nil {
		log.Fatalf("failed to create protocol: %v", err)
	}

	defer s.Close(context.Background())

	c, err := cloudevents.NewClient(s, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	for i := 0; i < 10; i++ {
		e := cloudevents.NewEvent()
		e.SetType("com.cloudevents.sample.sent")
		e.SetSource("https://github.com/cloudevents/sdk-go/v2/cmd/samples/stan/sender")
		_ = e.SetData(cloudevents.ApplicationJSON, map[string]interface{}{
			"id":      i,
			"message": "Hello, World!",
		})

		result := c.Send(context.Background(), e)
		if !cloudevents.IsACK(result) {
			log.Printf("failed to send: %v", result)
		} else {
			log.Printf("sent: %d", i)
		}
	}
}

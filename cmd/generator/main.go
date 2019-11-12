package main

import (
	"context"
	"math/rand"
	"os"
	"time"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/n3wscott/d2k/pkg/api"
	"github.com/kelseyhightower/envconfig"
)


type envConfig struct {
	Sink string `envconfig:"SINK" default:"http://localhost:8080/"`
}

func main() {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("failed to process env var: %s", err)
		os.Exit(1)
	}

	ctx := cloudevents.ContextWithTarget(context.Background(), env.Sink)

	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	for true {
		// Sleep 1 second.
		time.Sleep(1*time.Second)

		// Send a random number [0-100)
		event := cloudevents.NewEvent("1.0")
		event.SetType("io.d2k8s.number")
		event.SetSource("https://github.com/n3wscott/d2k/cmd/generator")
		event.SetDataContentType(cloudevents.ApplicationJSON)
		if err := event.SetData(&api.Number{
			Number: rand.Intn(100),
		}); err != nil {
			log.Printf("failed to set data: %v", err)
			continue
		}

		if _, _, err := c.Send(ctx, event); err != nil {
			log.Printf("failed to send: %v", err)
		}
	}
}
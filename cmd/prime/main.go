package main

import (
	"context"
	"fmt"
	"log"
	"math"

	"github.com/n3wscott/d2k/pkg/api"

	cloudevents "github.com/cloudevents/sdk-go"
)


func main() {
	ctx := context.Background()

	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Printf("will listen on :8080\n")
	log.Fatalf("failed to start receiver: %s", c.StartReceiver(ctx, classify))
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func classify(ctx context.Context, event cloudevents.Event, r *cloudevents.EventResponse) {
	data := &api.Number{}
	if err := event.DataAs(data); err != nil {
		fmt.Printf("failed to get data as Number: %s\n", err.Error())
		return
	}

	resp := cloudevents.NewEvent("0.3")

	resp.SetSource("https://github.com/n3wscott/d2k/cmd/prime")
	resp.SetDataContentType(cloudevents.ApplicationJSON)
	if isPrime(data.Number) {
		resp.SetType("io.d2k8s.number.prime")
		data.Classification = "prime"
	} else {
		// Not prime, skip.
		r.RespondWith(204, nil)
	}
	if err := resp.SetData(data); err != nil {
		log.Printf("failed to set data: %v", err)
	}
	r.RespondWith(200, &resp)
}
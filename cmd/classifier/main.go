package main

import (
	"context"
	"fmt"
	"log"

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

func classify(ctx context.Context, event cloudevents.Event, r *cloudevents.EventResponse) {
	data := &api.Number{}
	if err := event.DataAs(data); err != nil {
		fmt.Printf("failed to get data as Number: %s\n", err.Error())
		return
	}

	resp := cloudevents.NewEvent("0.3")

	resp.SetSource("https://github.com/n3wscott/d2k/cmd/classifier")
	resp.SetDataContentType(cloudevents.ApplicationJSON)
	if data.Number % 2 == 0 {
		resp.SetType("io.d2k8s.number.even")
		data.Classification = "even"
		resp.SetExtension("classification", "even")
	} else {
		resp.SetType("io.d2k8s.number.odd")
		data.Classification = "odd"
		resp.SetExtension("classification", "odd")
	}
	if err := resp.SetData(data); err != nil {
		log.Printf("failed to set data: %v", err)
	}
	r.RespondWith(200, &resp)
}
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

func main() {
	fmt.Println("main started, with 5 seconds time out")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// if you uncomment the below line, notice that we get a 200 OK.
	// resp, err := http.Get("http://pubsub:8085")

	cl, err := pubsub.NewClient(ctx, "marwan-test")
	if err != nil {
		log.Fatalf("could not create client: %v", err)
	}

	t, err := cl.CreateTopic(ctx, "email")
	if err != nil {
		log.Fatalf("could not create topic: %v", err)
	}

	fmt.Printf("successfully created topic: %+v\n", t)
}

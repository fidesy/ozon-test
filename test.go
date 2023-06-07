package main

import (
	"context"
	"log"

	shortener "github.com/fidesy/ozon-test/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client for the URL service
	client := shortener.NewURLServiceClient(conn)

	// Make a request to the CreateURL method
	createURLRequest := &shortener.CreateShortURLRequest{
		OriginalUrl: "https://ozon.ruj",
	}
	createURLResponse, err := client.CreateShortURL(context.Background(), createURLRequest)
	if err != nil {
		log.Fatalf("CreateURL failed: %v", err)
	}

	// Process the response
	log.Printf("Short URL: %s", createURLResponse.ShortUrl)
}

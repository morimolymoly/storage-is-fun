package main

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

// YourName ... なまえだよ
type YourName struct {
	Value string
}

const (
	projectID string = "my-project"
	nameKeyID string = "YourName"
)

func main() {
	ctx := context.Background()

	log.Printf("Creating Client...")
	dsClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Can't create client: %s", err)
	}
	log.Printf("Created!\n")

	nameKey := datastore.NameKey(nameKeyID, "stringID", nil)
	yourname := new(YourName)
	yourname.Value = "MINO MONTA"
	log.Printf("Putting your name...")
	if _, err := dsClient.Put(ctx, nameKey, yourname); err != nil {
		log.Fatal("Can't push your name!")
	}
	log.Printf("Done!\n")

	log.Printf("Getting your name...")
	yourname = new(YourName)
	if err := dsClient.Get(ctx, nameKey, yourname); err != nil {
		log.Fatal("Can't get your name")
	}
	log.Printf("Done!\nHello，%s!", yourname.Value)

	if err := dsClient.Delete(ctx, nameKey); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

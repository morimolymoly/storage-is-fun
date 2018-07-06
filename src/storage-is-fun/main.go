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

	dsClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Can't create client: %s", err)
	}

	nameKey := datastore.NameKey(nameKeyID, "stringID", nil)
	yourname := new(YourName)

	if err := dsClient.Get(ctx, nameKey, yourname); err != nil {
		log.Printf("Can't get your name: %s", err)
	}

	oldName := yourname.Value
	yourname.Value = "MINO MONTA"

	if _, err := dsClient.Put(ctx, nameKey, yourname); err != nil {
		log.Fatal("Can't update your name!")
	}

	log.Printf("Updated value from %q to %q\n", oldName, yourname.Value)
}

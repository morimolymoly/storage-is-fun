package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
)

type YourName struct {
	Value string
}

func main() {
	ctx := context.Background()

	dsClient, err := datastore.NewClient(ctx, "my-project")
	if err != nil {
		log.Fatalf("Can't create client: %s", err)
	}

	nameKey := datastore.NameKey("YourName", "stringID", nil)
	yourname := new(YourName)

	if err := dsClient.Get(ctx, nameKey, yourname); err != nil {
		log.Printf("Can't get your name: %s", err)
	}

	oldName := yourname.Value
	yourname.Value = "MINO MONTA"

	if _, err := dsClient.Put(ctx, nameKey, yourname); err != nil {
		log.Fatal("Can't update your name!")
		return
	}

	fmt.Printf("Updated value from %q to %q\n", oldName, yourname.Value)
}

package client

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/go-sdk/client"
)

// ExampleDefault shows how to use the default client, which is lazily initialized.
// It will use the environment variables or the current Docker context to connect to the daemon.
func ExampleDefault() {
	cli := client.DefaultClient

	info, err := cli.Info(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(info.ServerVersion != "")

	// Output:
	// true
}

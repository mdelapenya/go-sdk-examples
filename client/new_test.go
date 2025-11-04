package client

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/go-sdk/client"
)

// ExampleNew shows how to create a new client.
func ExampleNew() {
	cli, err := client.New(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		err = cli.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	info, err := cli.Info(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(info.ServerVersion != "")

	// Output:
	// true
}

// ExampleNew_useUnderlyingClient shows how to use the underlying Docker client.
func ExampleNew_useUnderlyingClient() {
	cli, err := client.New(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	ping, err := cli.Ping(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(ping.OSType != "")

	// Output:
	// true
}

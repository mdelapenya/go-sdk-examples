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

// ExampleNew_withCustomHost_invalid shows how to create a new client with an invalid host.
func ExampleNew_withCustomHost_invalid() {
	cli, err := client.New(context.Background(), client.WithDockerHost("tcp://127.0.0.1:1234"))
	if err != nil {
		log.Println(err)
	}

	// Because the host is not valid, the client will be nil.
	fmt.Println(cli == nil)

	// Output:
	// true
}

// ExampleNew_withCustomHost_valid shows how to create a new client with a valid host.
func ExampleNew_withCustomHost_valid() {
	cli, err := client.New(context.Background(), client.WithDockerHost("unix:///var/run/docker.sock"))
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err = cli.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	fmt.Println(cli == nil)

	// Output:
	// false
}

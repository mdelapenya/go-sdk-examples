package client

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/go-sdk/client"
	dockercontext "github.com/docker/go-sdk/context"
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

// ExampleNew_withCustomHost shows how to create a new client with a valid host, obtained from the current context.
func ExampleNew_withCustomHost() {
	dockerHost, err := dockercontext.CurrentDockerHost()
	if err != nil {
		log.Println(err)
		return
	}

	cli, err := client.New(context.Background(), client.WithDockerHost(dockerHost))
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

	fmt.Println(cli == nil)

	// Output:
	// false
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

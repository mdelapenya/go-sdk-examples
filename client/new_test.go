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
	dockerHost, err := dockercontext.CurrentDockerHost()
	if err != nil {
		log.Println(err)
		return
	}

	cli, err := client.New(context.Background(), client.WithDockerHost(dockerHost))
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

// ExampleNew_useUnderlyingClient shows how to use the underlying Docker client.
func ExampleNew_useUnderlyingClient() {
	cli, err := client.New(context.Background())
	if err != nil {
		log.Println(err)
	}

	dockerClient, err := cli.Client()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dockerClient != nil)

	ping, err := dockerClient.Ping(context.Background())
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ping.OSType != "")

	// Output:
	// true
	// true
}

// ExampleDefault_useUnderlyingClient shows how to use the underlying Docker client.
func ExampleDefault_useUnderlyingClient() {
	dockerClient, err := client.DefaultClient.Client()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dockerClient != nil)

	ping, err := dockerClient.Ping(context.Background())
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ping.OSType != "")

	// Output:
	// true
	// true
}

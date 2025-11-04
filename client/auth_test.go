package client

import (
	"context"
	"fmt"

	"github.com/docker/go-sdk/client"
	"github.com/docker/go-sdk/client/registry"
)

func Example_authConfigForImage() {
	cli, err := client.New(context.Background())
	defer func() {
		if cli != nil {
			err = cli.Close()
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}

	imageRegistry, authConfig, err := cli.AuthConfigForImage("nginx:latest")
	fmt.Println(err)
	fmt.Println(imageRegistry)
	fmt.Println(authConfig.ServerAddress)

	// Output:
	// <nil>
	// docker.io
	// docker.io
}

func Example_authConfigForHostname() {
	cli, err := client.New(context.Background())
	defer func() {
		if cli != nil {
			err = cli.Close()
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = cli.AuthConfigForHostname(registry.IndexDockerIO)
	fmt.Println(err)

	// Output:
	// <nil>
}

func Example_authConfigForHostname_publicRegistry() {
	cli, err := client.New(context.Background())
	defer func() {
		if cli != nil {
			err = cli.Close()
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = cli.AuthConfigForHostname(registry.DockerRegistry)
	fmt.Println(err)

	// Output:
	// <nil>
}

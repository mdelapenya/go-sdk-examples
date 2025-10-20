package config

import (
	"fmt"

	"github.com/docker/go-sdk/config"
	"github.com/docker/go-sdk/config/auth"
)

func ExampleAuthConfigs() {
	authConfigs, err := config.AuthConfigs("nginx:latest")
	fmt.Println(err)
	fmt.Println(len(authConfigs))
	fmt.Println(authConfigs[auth.DockerRegistry].ServerAddress)

	// Output:
	// <nil>
	// 1
	// docker.io
}

func ExampleAuthConfigForHostname() {
	_, err := config.AuthConfigForHostname(auth.IndexDockerIO)
	fmt.Println(err)

	// Output:
	// <nil>
}

func ExampleAuthConfigForHostname_publicRegistry() {
	_, err := config.AuthConfigForHostname(auth.DockerRegistry)
	fmt.Println(err)

	// Output:
	// <nil>
}

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
	authConfig, err := config.AuthConfigForHostname(auth.IndexDockerIO)
	fmt.Println(err)
	fmt.Println(authConfig.Username != "")

	// Output:
	// <nil>
	// true
}

func ExampleAuthConfigForHostname_publicRegistry() {
	authConfig, err := config.AuthConfigForHostname(auth.DockerRegistry)
	fmt.Println(err)
	fmt.Println(authConfig.Username != "")

	// Output:
	// <nil>
	// true
}

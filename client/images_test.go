package client

import (
	"fmt"

	"github.com/docker/go-sdk/client/registry"
)

func Example_parseImagRef() {
	ref, err := registry.ParseImageRef("nginx:latest")
	if err != nil {
		fmt.Println("Error parsing image:", err)
		return
	}

	fmt.Println(ref.Registry)
	fmt.Println(ref.Repository)
	fmt.Println(ref.Tag)
	fmt.Println(ref.Digest)

	// Output:
	// docker.io
	// library/nginx
	// latest
	//
}

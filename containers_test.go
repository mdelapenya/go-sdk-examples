package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/docker/go-sdk/container"
)

func ExampleRun() {
	ctr, err := container.Run(
		context.Background(),
		container.WithImage("nginx:alpine"),
		container.WithExposedPorts("80/tcp"),
	)
	if err != nil {
		log.Println(err)
		return
	}

	port, err := ctr.MappedPort(context.Background(), "80/tcp")
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.Get("http://localhost:" + port.Port())
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(strings.Contains(string(body), "<title>Welcome to nginx!</title>"))

	err = container.Terminate(ctr)
	if err != nil {
		log.Println(err)
		return
	}

	// Output:
	// true

}

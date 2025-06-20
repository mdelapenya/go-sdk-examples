package containers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/docker/go-sdk/client"
	"github.com/docker/go-sdk/container"
	"github.com/docker/go-sdk/container/wait"
)

// ExampleRun_withLogger shows how to run a container with a custom logger.
// The logger is part of the Docker client configuration, so instead of using the
// default Docker client, we create a new one with a custom logger.
func ExampleRun_withLogger() {
	buf := bytes.NewBuffer(nil)
	logger := slog.New(slog.NewTextHandler(buf, nil))

	cli, err := client.New(context.Background(), client.WithLogger(logger))
	if err != nil {
		log.Println(err)
		return
	}

	ctr, err := container.Run(
		context.Background(),
		container.WithDockerClient(cli),
		container.WithImage("nginx:alpine"),
		container.WithExposedPorts("80/tcp"),
		container.WithWaitStrategy(wait.ForListeningPort("80/tcp").WithTimeout(time.Second*5)),
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

	log := buf.String()

	fmt.Println(strings.Contains(log, "Creating container"))
	fmt.Println(strings.Contains(log, "Container created"))
	fmt.Println(strings.Contains(log, "Starting container"))
	fmt.Println(strings.Contains(log, "Container started"))
	fmt.Println(strings.Contains(log, "Container is ready"))
	fmt.Println(strings.Contains(log, "Stopping container"))
	fmt.Println(strings.Contains(log, "Container stopped"))
	fmt.Println(strings.Contains(log, "Terminating container"))
	fmt.Println(strings.Contains(log, "Container terminated"))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
}

package containers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/docker/go-sdk/container"
	"github.com/docker/go-sdk/container/wait"
)

// ExampleRun shows how to run a container, exposing the default HTTP port to a random port
// and waiting for it to be ready.
func ExampleRun() {
	ctr, err := container.Run(
		context.Background(),
		container.WithImage("nginx:alpine"),
		container.WithExposedPorts("80/tcp"),
		container.WithWaitStrategy(wait.ForListeningPort("80/tcp").WithTimeout(time.Second*5)),
	)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		err = container.Terminate(ctr)
		if err != nil {
			log.Println(err)
		}
	}()

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

	// Output:
	// true

}

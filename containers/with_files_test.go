package containers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/docker/go-sdk/container"
	"github.com/docker/go-sdk/container/wait"
)

// ExampleRun_withFiles shows how to run a container with a file inside it.
// The file is copied after the container is created and before it is started.
func ExampleRun_withFiles() {
	content := []byte("<h1>Hello, Moby!</h1>")

	ctr, err := container.Run(
		context.Background(),
		container.WithImage("nginx:alpine"),
		container.WithExposedPorts("80/tcp"),
		container.WithFiles(container.File{
			Reader:        bytes.NewReader(content),
			ContainerPath: "/usr/share/nginx/html/index.html",
			Mode:          0o644,
		}),
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

	fmt.Println(string(body))

	// Output:
	// <h1>Hello, Moby!</h1>
}

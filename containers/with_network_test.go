package containers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/go-sdk/container"
	"github.com/docker/go-sdk/container/wait"
	"github.com/docker/go-sdk/network"
)

// ExampleRun_withNetwork shows how to run a container attached to a network,
// using an alias in that network.
func ExampleRun_withNetwork() {
	nw, err := network.New(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := nw.Terminate(context.Background()); err != nil {
			log.Println(err)
		}
	}()

	ctr, err := container.Run(
		context.Background(),
		container.WithImage("nginx:alpine"),
		container.WithExposedPorts("80/tcp"),
		container.WithNetwork([]string{"web-server"}, nw),
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

	fmt.Println(ctr.ID() != "")

	// Output:
	// true
}

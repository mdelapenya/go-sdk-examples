package networks

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/go-sdk/network"
)

func ExampleNew() {
	nw, err := network.New(context.Background(), network.WithName("my-network"))
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := nw.Terminate(context.Background()); err != nil {
			log.Println(err)
		}
	}()

	fmt.Println(nw.Name())

	// Output:
	// my-network
}

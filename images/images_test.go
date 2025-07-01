package images

import (
	"fmt"
	"log"
	"strings"

	"github.com/docker/go-sdk/image"
)

func Example_readImagesFromDockerfile() {
	dockerfile := `FROM alpine:latest
FROM busybox:latest
FROM ${MY_IMAGE}`

	nginx := "nginx:latest"

	imgs, err := image.ImagesFromReader(strings.NewReader(dockerfile), map[string]*string{
		"MY_IMAGE": &nginx,
	})
	if err != nil {
		log.Println("Error reading images from Dockerfile:", err)
		return
	}

	for _, img := range imgs {
		fmt.Println(img)
	}

	// Output:
	// alpine:latest
	// busybox:latest
	// nginx:latest
}

package images

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"log/slog"
	"strings"

	"github.com/docker/go-sdk/client"
	"github.com/docker/go-sdk/image"
)

func Example_build_fromContext() {
	buf := bytes.NewBuffer(nil)
	cli, err := client.New(context.Background(), client.WithLogger(slog.New(slog.NewTextHandler(buf, nil))))
	if err != nil {
		log.Println("Error creating client:", err)
		return
	}

	reader, err := image.ArchiveBuildContext("testdata/multiple", "Dockerfile")
	if err != nil {
		log.Println("Error creating build context:", err)
		return
	}

	tag, err := image.Build(
		context.Background(),
		reader,
		"my-image:context",
		image.WithBuildClient(cli),
	)
	if err != nil {
		log.Println("Error building image:", err)
		return
	}

	fmt.Println(tag)

	// Output:
	// my-image:context
}

func Example_buildFromDir_simple() {
	buf := bytes.NewBuffer(nil)
	cli, err := client.New(context.Background(), client.WithLogger(slog.New(slog.NewTextHandler(buf, nil))))
	if err != nil {
		log.Println("Error creating client:", err)
		return
	}

	tag, err := image.BuildFromDir(
		context.Background(),
		"testdata/simple",
		"Dockerfile",
		"my-image:simple",
		image.WithBuildClient(cli),
	)
	if err != nil {
		log.Println("Error building image:", err)
		return
	}

	fmt.Println(tag)

	// Output:
	// my-image:simple
}

func Example_buildFromDir_multiple() {
	buf := bytes.NewBuffer(nil)
	cli, err := client.New(context.Background(), client.WithLogger(slog.New(slog.NewTextHandler(buf, nil))))
	if err != nil {
		log.Println("Error creating client:", err)
		return
	}

	tag, err := image.BuildFromDir(
		context.Background(),
		"testdata/multiple",
		"Dockerfile",
		"my-image:multiple",
		image.WithBuildClient(cli),
	)
	if err != nil {
		log.Println("Error building image:", err)
		return
	}

	fmt.Println(tag)

	// Output:
	// my-image:multiple
}

func Example_imagesFromDockerfile() {
	images, err := image.ImagesFromDockerfile("testdata/complex/Dockerfile", nil)
	if err != nil {
		log.Println("Error reading images from Dockerfile:", err)
		return
	}

	for _, img := range images {
		fmt.Println(img)
	}

	// Output:
	// golang:1.23-alpine@sha256:f8113c4b13e2a8b3a168dceaee88ac27743cc84e959f43b9dbd2291e9c3f57a0
}

func Example_imagesFromReader() {
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

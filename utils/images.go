package utils

import (
	"context"
	"fmt"

	client "docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

func ListImages() {

	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}


	//List all images available locally
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("LIST IMAGES\n-----------------------")
	fmt.Println("Image ID | Repo Tags | Size")
	for _, image := range images {
		fmt.Printf("%s | %s | %d\n", image.ID, image.RepoTags, image.Size)
	}

}

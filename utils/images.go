package utils

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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

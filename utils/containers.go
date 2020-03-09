package utils

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ListContainers() {

	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	//Retrieve a list of containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Print("\n\n\n")
	fmt.Println("LIST CONTAINERS\n-----------------------")
	fmt.Println("Container Names | Image | Mounts")
	//Iterate through all containers and display each container's properties
	for _, container := range containers {
		fmt.Printf("%s | %s | %s\n", container.Names, container.Image, container.Mounts)
	}

}

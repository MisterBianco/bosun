package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/docker/docker/api/types"
	dockerclient "github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
)

func BuildImage(conf Configuration) {

	fmt.Println(conf.Path, conf.Dockerfile)
	ctx := context.Background()
	dockerClient, err := dockerclient.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	buildCtx, err := archive.TarWithOptions(path.Join(conf.Path, conf.Dockerfile, "../"), &archive.TarOptions{})
	if err != nil {
		log.Fatal(err)
	}

	imageBuildResponse, err := dockerClient.ImageBuild(
		ctx,
		buildCtx,
		types.ImageBuildOptions{
			Context: buildCtx,
			Remove: true,
			Tags: conf.Tags,
			BuildArgs: conf.Args})

	if err != nil {
		fmt.Printf("Failed to build docker image: %s\n", err)
		os.Exit(1)
	}
	defer imageBuildResponse.Body.Close()

	// _, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	// if err != nil {
	// 	log.Fatal(err, " :unable to read image build response")
	// }

	termFd, isTerm := term.GetFdInfo(os.Stdout)
	jsonmessage.DisplayJSONMessagesStream(imageBuildResponse.Body, os.Stdout, termFd, isTerm, nil)
}

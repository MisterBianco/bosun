package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/docker/docker/api/types"
	dockerclient "github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
)


func PushImage(conf Configuration) {
	ctx := context.Background()
	dockerClient, err := dockerclient.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	// sess, _ := session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// })
	// svc := ecr.New(sess)
	// ain := &ecr.GetAuthorizationTokenInput{}
	// aout, _ := svc.GetAuthorizationToken(ain)

	for _, imageName := range conf.Tags {
		imagePushResponse, err := dockerClient.ImagePush(
			ctx,
			imageName,
			types.ImagePushOptions{
				All: true})
				// RegistryAuth: *aout.AuthorizationData[0].AuthorizationToken})

		if err != nil {
			log.Fatal(err, " :unable to build docker image")
		}

		defer imagePushResponse.Close()

		io.Copy(os.Stdout, imagePushResponse)

		// defer imagePushResponse.Body.Close()

		// termFd, isTerm := term.GetFdInfo(os.Stderr)
		// jsonmessage.DisplayJSONMessagesStream(imagePushResponse.Body, os.Stderr, termFd, isTerm, nil)
	}
}

func CreateImage(conf Configuration, filePath string) {
	ctx := context.Background()
	dockerClient, err := dockerclient.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(filePath)
	buildCtx, err := archive.TarWithOptions(path.Join(filePath, "../"), &archive.TarOptions{})
	if err != nil {
		log.Fatal(err)
	}

	imageBuildResponse, err := dockerClient.ImageBuild(
		ctx,
		buildCtx,
		types.ImageBuildOptions{
			Context: buildCtx,
			// Dockerfile: filePath,
			Remove: true,
			Tags: conf.Tags,
			BuildArgs: conf.Args})

	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	defer imageBuildResponse.Body.Close()

	// _, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	// if err != nil {
	// 	log.Fatal(err, " :unable to read image build response")
	// }

	termFd, isTerm := term.GetFdInfo(os.Stderr)
	jsonmessage.DisplayJSONMessagesStream(imageBuildResponse.Body, os.Stderr, termFd, isTerm, nil)

	// Might work
	// func writeToLog(reader io.ReadCloser) error {
	// 	defer reader.Close()
	// 	rd := bufio.NewReader(reader)
	// 	for {
	// 		n, _, err := rd.ReadLine()
	// 		if err != nil && err == io.EOF {
	// 			break
	// 		} else if err != nil {
	// 			return err
	// 		}
	// 		log.Println(string(n))
	// 	}
	// 	return nil
	// }
}

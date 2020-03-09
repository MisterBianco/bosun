package utils

import (
	"archive/tar"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	dockerclient "github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
)

func CreateImage(conf Configuration, filePath string) {
	ctx := context.Background()
	dockerClient, err := dockerclient.NewEnvClient()
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
    tw := tar.NewWriter(buf)
    defer tw.Close()

	dockerFileReader, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err, "cant open dockerfile")
	}

	dockerFile, err := ioutil.ReadAll(dockerFileReader)
	if err != nil {
		log.Fatal(err, "cant read dockerfile")
	}

	tarHeader := &tar.Header{
		Name: filePath,
		Size: int64(len(dockerFile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Fatal(err, " :unable to write tar header")
	}
	_, err = tw.Write(dockerFile)
	if err != nil {
		log.Fatal(err, " :unable to write tar body")
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())
	imageBuildResponse, err := dockerClient.ImageBuild(
		ctx,
		dockerFileTarReader,
		types.ImageBuildOptions{
			Context: dockerFileTarReader,
			Dockerfile: filePath,
			Remove: true,
			Tags: conf.Tags,
			BuildArgs: conf.Args})

	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	defer imageBuildResponse.Body.Close()

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

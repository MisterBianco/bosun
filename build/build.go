package build

import (
	"fmt"
	"path"
	"strings"

	"github.com/misterbianco/bosun/utils"
)

type Build struct {

}

func (build *Build) BuildImage(args []string) {
	var images []string

	for _, argpath := range args {
		fmt.Printf("Checking for build context and configuration: %s\n", argpath)

		var conf utils.Configurations
		conf.GetConfigurations(path.Join(argpath, "bosun.yml"))

		for _, config := range conf.Configs {
			config.GenerateTags()

			// fmt.Println(config.Tags)

			dockerfile := config.Dockerfile
			if config.Dockerfile == "" {
				dockerfile = "Dockerfile"
			}

			fmt.Println(path.Join(argpath, dockerfile))

			utils.CreateImage(config, path.Join(argpath, dockerfile))

			for _, tag := range config.Tags {
				images = append(images, tag)
			}
		}
	}

	fmt.Println("\r\n------------------------------------")
	fmt.Println("docker push",strings.Join(images, " && docker push "))
}

package build

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/misterbianco/bosun/utils"
	"gopkg.in/yaml.v2"
)

type Build struct {
	Configurations utils.Configurations
	Quiet bool
	File string
	Images []string
}

func (self *Build) GetConfigurations(filePath string) *Build {
	configFile, err := os.Open(path.Join(filePath, self.File))

	if err != nil {
		fmt.Printf("Failed to open file: %s\n", path.Join(filePath, self.File))
		return self
	}

	defer configFile.Close()

	var decoder = yaml.NewDecoder(configFile)
	decoder.SetStrict(true)

	var yamlConfig utils.Configuration

	for {
		err := decoder.Decode(&yamlConfig)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		yamlConfig.Path = filePath
		yamlConfig.Tags = utils.GenerateTags(yamlConfig.Name, yamlConfig.Tags)
		if yamlConfig.Dockerfile == "" {
			yamlConfig.Dockerfile = "Dockerfile"
		}

		self.Configurations = append(self.Configurations, yamlConfig)
	}

	return self
}
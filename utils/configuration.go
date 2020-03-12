package utils

import (
	"fmt"
	"os"

	"github.com/blang/semver"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Name string `yaml:"name"`
	Tags []string `yaml:"tags"`
	Args map[string]*string `yaml:"args"`
	Labels map[string]*string `yaml:"labels"`
	Dockerfile string `yaml:"dockerfile"`
	// Labels are not recommended since they are injected as the last build step.
}

type Configurations struct {
	Configs []Configuration
}

func (c *Configurations) GetConfigurations(filePath string) *Configurations {

	conf, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("File %s couldn't be read: %s\n", filePath, err)
		os.Exit(-1)
	}
	defer conf.Close()

	// err = yaml.Unmarshal(conf, c)
	// if err != nil {
	// 	fmt.Printf("Unmarshalling failure: %s\n", err)
	// 	os.Exit(-1)
	// }

	var decoder = yaml.NewDecoder(conf)

	var thing Configuration
	for decoder.Decode(&thing) == nil {
		c.Configs = append(c.Configs, thing)
	}

	return c
}

func (c *Configuration) GenerateTags() *Configuration {
	result := []string{}

	for j := 0; j < len(c.Tags); j++ {
		v, err := semver.Parse(c.Tags[j])
		if err != nil {
			fmt.Println(err)
			result = append(result, c.Tags[j])
			continue
		}

		if v.Pre == nil {
			result = append(result, fmt.Sprintf("%s:%v",c.Name,v.Major))
			result = append(result, fmt.Sprintf("%s:%v.%v", c.Name,v.Major, v.Minor))
			result = append(result, fmt.Sprintf("%s:%v.%v.%v", c.Name,v.Major, v.Minor, v.Patch))
		} else {
			result = append(result, fmt.Sprintf("%s:%v-%s", c.Name, v.Major, v.Pre[0]))
			result = append(result, fmt.Sprintf("%s:%v.%v-%s", c.Name,v.Major, v.Minor, v.Pre[0]))
			result = append(result, fmt.Sprintf("%s:%v.%v.%v-%s", c.Name,v.Major, v.Minor, v.Patch, v.Pre[0]))
		}
	}
	result = append(result, fmt.Sprintf("%s:latest", c.Name))

	c.Tags = result
	return c
}

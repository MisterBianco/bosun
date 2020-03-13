package utils

type Configuration struct {
	Name string `yaml:"name"`
	Tags []string `yaml:"tags"`
	Args map[string]*string `yaml:"args"`
	Labels map[string]*string `yaml:"labels"`
	Dockerfile string `yaml:"dockerfile"`
	Path string
}

type Configurations []Configuration
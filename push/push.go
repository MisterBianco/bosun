package push

import (
	"fmt"
	"path"

	"github.com/misterbianco/bosun/utils"
)

type Push struct {

}

func (push *Push) PushImage(args []string) {
	for _, argpath := range args {
		fmt.Printf("Checking for build context and configuration: %s\n", argpath)

		var conf utils.Configurations
		conf.GetConfigurations(path.Join(argpath, "bosun.yml"))

		for _, config := range conf.Configs {
			config.GenerateTags()

			fmt.Println(config.Tags)

			utils.PushImage(config)
		}
	}
}

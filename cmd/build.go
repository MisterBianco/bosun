/*
Copyright © 2020 jarad dingman <jarad.dingman@wgu.edu>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"
	"path"

	"github.com/misterbianco/bosun/utils"

	"github.com/spf13/cobra"
)

var intro = `
______
| ___ \
| |_/ / ___  ___ _   _ _ __
| ___ \/ _ \/ __| | | | '_ \
| |_/ / (_) \__ \ |_| | | | |
\____/ \___/|___/\__,_|_| |_|

-------------------------------------------------
`

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Command to build a docker image and tag it",
	Long:  `The command to build the dockerfile on a given path`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(intro)

		for i := 0; i < len(args); i++ {
			fmt.Printf("Checking for build context and configuration: %s\n", args[i])

			var conf utils.Configuration
			conf.GetConfiguration(path.Join(args[i], "bosun.yml"))
			conf.GenerateTags()

			utils.CreateImage(conf, path.Join(args[i], "Dockerfile"))
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

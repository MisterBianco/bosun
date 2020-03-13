/*
Copyright 2020 Jayrad1996@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
DEALINGS IN THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/misterbianco/bosun/build"
	"github.com/misterbianco/bosun/utils"
	"github.com/spf13/cobra"
)

var quiet bool
var file string

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build your images",
	Long: `Build your images`,
	Run: func(cmd *cobra.Command, args []string) {
		app := build.Build{
			Quiet: quiet,
			File: file}

		for _, filePath := range args {
			configPath, err := filepath.Abs(filePath)
			if err != nil {
				fmt.Println("Failed to get absolute path to:", filePath)
				os.Exit(1)
			}
			app.GetConfigurations(configPath)
		}

		for _, config := range app.Configurations {
			utils.BuildImage(config)
			app.Images = append(app.Images, config.Tags...)
		}

		fmt.Println("\r\n------------------------------------")
		fmt.Println("docker push", strings.Join(app.Images, " && docker push "))
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.
	buildCmd.Flags().StringVarP(&file, "file", "f", "bosun.yml", "The config file to use.")
	buildCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Run in quiet mode")
}

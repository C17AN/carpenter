/*
Copyright © 2022 Chanmin, Kim <kimchanmin1@gmail.com>

*/
package cmd

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type buildMetadata struct {
	tag            string
	dockerfilePath string
	architecture   string
}

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build docker image with interactive window",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Select Your Architecture Type :")
		dockerfilePathPrompt := promptui.Prompt{
			Label:   "Input dockerfile path",
			Default: ".",
		}

		architectureTypeSelect := promptui.Select{
			Label: "Architecture",
			Items: []string{"ARM64", "AMD64", "ARM/v7"},
		}

		imageTagPrompt := promptui.Prompt{
			Label: "Target Image name",
		}

		isUsingCachePrompt := promptui.Select{
			Label: "Use layer cache on build",
			Items: []string{"yes", "no"},
		}

		dockerfilePath, _ := dockerfilePathPrompt.Run()
		_, architectureName, _ := architectureTypeSelect.Run()
		imageTagname, _ := imageTagPrompt.Run()
		_, isUsingCache, _ := isUsingCachePrompt.Run()
		fmt.Printf(`Image will created with those info : 
		Dockerfile Path: %q
		Target architecture: %q
		Target Image tag: %q
		Using Cache: %s`, dockerfilePath, architectureName, imageTagname, isUsingCache)

		command := fmt.Sprintf("docker build --platform=linux/%s -t %s %s", architectureName, imageTagname, dockerfilePath)
		confirmBuildPrompt := promptui.Prompt{
			Label:   "Proceed with these settings? [Y/n]",
			Default: "y",
		}

		isConfirmed, _ := confirmBuildPrompt.Run()
		if isConfirmed == "y" || isConfirmed == "yes" {
			execute := exec.Command("bash", "-c", command)

			// line 69 ~ 도커 빌드 로그 출력해주는 코드
			stdout, _ := execute.StdoutPipe()
			// start the command after having set up the pipe
			if err := execute.Start(); err != nil {
				fmt.Println("Error")
			}

			// read command's stdout line by line
			scanner := bufio.NewScanner(stdout)

			for scanner.Scan() {
				fmt.Println(scanner.Text()) // write each line to your fmt, or anything you need
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("error: %s", err)
			}
			_ = execute.Wait()
		} else {
			fmt.Printf("Aborting")
			for i := 0; i < 3; i++ {
				fmt.Printf("%s", strings.Repeat(".", 1))
				time.Sleep(500 * time.Millisecond)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2022 Chanmin, Kim <kimchanmin1@gmail.com>

*/
package cmd

import (
	"fmt"

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
	Short: "아키텍처",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Select Your Architecture Type :")

		dockerfilePathPrompt := promptui.Prompt{
			Label: "Input dockerfile path",
		}

		architectureTypeSelect := promptui.Select{
			Label: "Architecture",
			Items: []string{"ARM64", "AMD64"},
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

		command := fmt.Sprintf("--platform=%s -t %s %s", architectureName, imageTagname, dockerfilePath)
		confirmBuildPrompt := promptui.Prompt{
			Label: "Proceed with these settings?",
		}

		isConfirmed, _ := confirmBuildPrompt.Run()
		if isConfirmed == "y" {
			execute := exec.Command("docker", "build", command)
			stdout, err := execute.Output()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			// Print the output
			fmt.Println(string(stdout))
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

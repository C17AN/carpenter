/*
Copyright © 2022 Chanmin, Kim <kimchanmin1@gmail.com>

*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"carpenter/utils"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build docker image with interactive window",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// name := petname.Generate(*words, *separator)
		fmt.Println()
		fmt.Println("Select Your Architecture Type :")

		var defaultTagName = utils.GenerateUniqueTag()

		dockerfilePathPrompt := promptui.Prompt{
			Label:   "Input dockerfile path",
			Default: ".",
		}

		architectureTypeSelect := promptui.Select{
			Label: "Architecture",
			Items: []string{"ARM64", "AMD64", "ARM/v7", "wowwowwow"},
		}

		imageTagPrompt := promptui.Prompt{
			Label:    "Target Image tagname",
			Default:  defaultTagName,
			Validate: utils.ImageTagValidator,
		}

		imageVersionPrompt := promptui.Prompt{
			Label:   "Target Image version",
			Default: "latest",
		}

		isUsingCachePrompt := promptui.Select{
			Label: "Use layer cache on build",
			Items: []string{"yes", "no"},
		}

		dockerfilePath, err := dockerfilePathPrompt.Run()
		if err != nil {
			fmt.Println("Dead!")
		}
		_, architectureName, _ := architectureTypeSelect.Run()
		imageTagname, _ := imageTagPrompt.Run()
		imageVersion, _ := imageVersionPrompt.Run()
		_, isUsingCache, _ := isUsingCachePrompt.Run()

		fmt.Printf("\n--- Image will created with those info ---\n\n")
		fmt.Printf("\"Dockerfile Path\" : %q\n", dockerfilePath)
		fmt.Printf("\"Target architecture\" : %q\n", architectureName)
		fmt.Printf("\"Target Image tag\" : %q\n", imageTagname)
		fmt.Printf("\"Target Image version\" : %q\n", imageVersion)
		fmt.Printf("\"Using Cache\" : %s\n\n", isUsingCache)

		command := fmt.Sprintf("docker build --platform=linux/%s -t %s:%s %s", architectureName, imageTagname, imageVersion, dockerfilePath)
		confirmBuildPrompt := promptui.Prompt{
			Label:     "Proceed with these settings? [Y/n]",
			Default:   "y",
			IsConfirm: true,
		}

		isConfirmed, _ := confirmBuildPrompt.Run()

		if isConfirmed == "y" || isConfirmed == "yes" {
			execute := exec.Command("bash", "-c", command)

			// line 85 ~ 도커 빌드 로그 출력해주는 코드
			f, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			mwriter := io.MultiWriter(f, os.Stdout)
			execute.Stderr = mwriter
			execute.Stdout = mwriter

			err = execute.Run() //blocks until sub process is complete
			if err != nil {
				panic(err)
			}
		}
		// } else {
		// 	fmt.Printf("Aborting")
		// 	for i := 0; i < 3; i++ {
		// 		fmt.Printf("%s", strings.Repeat(".", 1))
		// 		time.Sleep(300 * time.Millisecond)
		// 	}
		// }
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/project-flogo/cli/api"
	"github.com/project-flogo/cli/util"
	"github.com/spf13/cobra"
)

var appName string
var flogoJsonPath string
var goPath = os.Getenv("GOPATH")
var examplePath = filepath.Join("github.com", "project-flogo", "cli-plugins", "examples", "app")

var CreateCmd = &cobra.Command{
	Use:              "create",
	Short:            "Create Command ",
	Long:             `This command helps you to work flogo contributions `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			appName = args[0]
		} else {
			fmt.Fprintf(os.Stderr, "Please provide the App name \n")
			os.Exit(1)
		}

		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error determining working directory: %v\n", err)
			os.Exit(1)
		}

		if !strings.HasPrefix(appName, "http") {
			//Install the example dir in GOPATH/src

			err := util.ExecCmd(exec.Command("go", "get", "github.com/project-flogo/cli-plugins/examples"), currentDir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error installing examples directory %v", err)
				os.Exit(1)
			}

			flogoJsonPath = filepath.Join(goPath, "src", examplePath, appName+".json")

			if _, err := os.Stat(flogoJsonPath); os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "Please provide the Valid App name")
				os.Exit(1)
			}
		}

		_, err = api.CreateProject(currentDir, appName, flogoJsonPath, "")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating project: %v\n", err)
			os.Exit(1)
		}
	},
}

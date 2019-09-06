package examples

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/project-flogo/cli-plugins/examples/commands"
	"github.com/project-flogo/cli/common"
	"github.com/project-flogo/cli/util"

	"github.com/spf13/cobra"
)

var exmpCmd = &cobra.Command{
	Use:   "example",
	Short: "Developer tool for basic work ",
	Long:  `This command helps you to work flogo contributions `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		if _, err := os.Stat(filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "project-flogo", "cli-plugins", "examples")); os.IsNotExist(err) {

			currentDir, err := os.Getwd()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error determining working directory: %v\n", err)
				os.Exit(1)
			}

			err = util.ExecCmd(exec.Command("go", "get", "github.com/project-flogo/cli-plugins/examples"), currentDir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error installing examples directory %v", err)
				os.Exit(1)
			}
		}

	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	exmpCmd.AddCommand(commands.CreateCmd)
	exmpCmd.AddCommand(commands.ListAppCmd)
	common.RegisterPlugin(exmpCmd)

}

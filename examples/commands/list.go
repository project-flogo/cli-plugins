package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"text/tabwriter"

	"github.com/project-flogo/cli/util"
	"github.com/spf13/cobra"
)

var ListAppCmd = &cobra.Command{
	Use:              "list",
	Short:            "List Contributions Command ",
	Long:             `This command helps you to find flogo contributions `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {

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

		var result []string
		if _, err := os.Stat(filepath.Join(goPath, "src", examplePath)); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error listing examples directory %v", err)
			os.Exit(1)
		}
		err = filepath.Walk(filepath.Join(goPath, "src", examplePath), func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				result = append(result, info.Name())
			}

			return nil
		})

		if err != nil {
			fmt.Errorf("%v", err)
			os.Exit(1)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "Name\tDescription")

		for key, val := range result {
			if key == 10 {
				break
			}

			fmt.Fprintf(w, "%v\t \n", val[:len(val)-5])
		}
		w.Flush()
	},
}

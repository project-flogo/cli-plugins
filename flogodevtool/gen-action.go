package flogodevtool

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var genAction = &cobra.Command{
	Use:   "gen-action",
	Short: "Generate activity scaffold",
	Long:  `This subcommand helps you generate activity-scaffold`,

	Run: func(cmd *cobra.Command, args []string) {
		var actionContrib string

		if len(args) < 1 {
			actionContrib = "action"
		} else {
			actionContrib = args[0]
		}

		err := os.Mkdir(actionContrib, os.ModePerm)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating dir: %v\n", err)
			os.Exit(1)
		}

		pwd, err := os.Getwd()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting current dir: %v\n", err)
			os.Exit(1)
		}

		err = copyFiles(filepath.Join(COREPATH, "action"), filepath.Join(pwd, actionContrib))

	},
}

func init() {
	descCmd.AddCommand(genAction)
}

package flogodevtool

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var genTrigger = &cobra.Command{
	Use:   "gen-trigger",
	Short: "Generate activity scaffold",
	Long:  `This subcommand helps you generate activity-scaffold`,

	Run: func(cmd *cobra.Command, args []string) {
		var triggerContrib string

		if len(args) < 1 {
			triggerContrib = "trigger"
		} else {
			triggerContrib = args[0]
		}
		err := os.Mkdir(triggerContrib, os.ModePerm)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating dir: %v\n", err)
			os.Exit(1)
		}

		pwd, err := os.Getwd()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting current dir: %v\n", err)
			os.Exit(1)
		}
		err = copyFiles(filepath.Join(COREPATH, "trigger"), filepath.Join(pwd, triggerContrib))

	},
}

func init() {
	descCmd.AddCommand(genTrigger)
}

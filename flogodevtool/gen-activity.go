package flogodevtool

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var genActivity = &cobra.Command{
	Use:   "gen-activity",
	Short: "Generate activity scaffold",
	Long:  `This subcommand helps you generate activity-scaffold`,

	Run: func(cmd *cobra.Command, args []string) {
		var activityContrib string

		if len(args) < 1 {
			activityContrib = "activity"
		} else {
			activityContrib = args[0]
		}

		err := os.Mkdir(activityContrib, os.ModePerm)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating dir: %v\n", err)
			os.Exit(1)
		}

		pwd, err := os.Getwd()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting current dir: %v\n", err)
			os.Exit(1)
		}

		err = copyFiles(filepath.Join(COREPATH, "activity"), filepath.Join(pwd, activityContrib))

	},
}

func init() {
	descCmd.AddCommand(genActivity)
}

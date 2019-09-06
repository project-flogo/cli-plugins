package examples

import (
	"github.com/project-flogo/cli-plugins/examples/commands"
	"github.com/project-flogo/cli/common"

	"github.com/spf13/cobra"
)

var exmpCmd = &cobra.Command{
	Use:              "example",
	Short:            "Developer tool for basic work ",
	Long:             `This command helps you to work flogo contributions `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	exmpCmd.AddCommand(commands.CreateCmd)
	exmpCmd.AddCommand(commands.ListAppCmd)
	common.RegisterPlugin(exmpCmd)

}

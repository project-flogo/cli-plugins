package devtool

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/project-flogo/cli/api"
	"github.com/project-flogo/cli/common"
	"github.com/project-flogo/cli/util"
	"github.com/spf13/cobra"
)

func init() {
	descCmd.AddCommand(updateMaster)
}

const (
	fJsonFile = "flogo.json"
	corePath  = "github.com/project-flogo/master"
)

var updateMaster = &cobra.Command{
	Use:   "upgrade-master",
	Short: "Update all contributions to master",
	Long:  `This subcommand helps you to upgrade all contributions to master branch`,
	Run: func(cmd *cobra.Command, args []string) {
		project := common.CurrentProject()

		imports, err := util.GetAppImports(filepath.Join(project.Dir(), fJsonFile), project.DepManager(), true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating all contributions: %v\n", err)
			os.Exit(1)
		}
		//Update each package in imports
		for _, imp := range imports.GetAllImports() {

			err = api.UpdatePkg(project, imp.GoImportPath()+"@master")

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error updating contribution/dependency: %v\n", err)
				os.Exit(1)
			}
		}

		err = api.UpdatePkg(project, corePath+"@master")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating contribution/dependency: %v\n", err)
			os.Exit(1)
		}
	},
}

package devtool

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/project-flogo/cli/common"

	"github.com/spf13/cobra"
)

var descCmd = &cobra.Command{
	Use:   "dev",
	Short: "Developer tool for basic work ",
	Long:  `This command helps you to work flogo contributions `,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var GOPATH string
var COREPATH string

func init() {

	common.RegisterPlugin(descCmd)

	GOPATH = os.Getenv("GOPATH")

	if _, err := os.Stat(filepath.Join(GOPATH, "pkg", "mod", "github.com", "project-flogo", "core@v0.9.0", "examples")); os.IsNotExist(err) {
		os.Setenv("GO111MODULE", "on")
		err = exec.Command("go", "get", "github.com/project-flogo/core@v0.9.0").Run()
		if err != nil {
			fmt.Printf("Error in getting 0.9.0 version of the core %v \n", err)
		}
	}
	COREPATH = filepath.Join(GOPATH, "pkg", "mod", "github.com", "project-flogo", "core@v0.9.0", "examples")
}

package flogodevtool

import (
	"encoding/json"
	"fmt"
	"go/parser"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"go/token"

	"github.com/project-flogo/core/data"
	"github.com/spf13/cobra"
)

var syncMetadata = &cobra.Command{
	Use:   "sync-metadata",
	Short: "sync descriptor json and metadata",
	Long:  `This subcommand creates descriptor json from metadata`,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()

		if err != nil {
			fmt.Errorf("Error in getting  current dir")
			os.Exit(1)
		}

		if _, err = os.Stat(filepath.Join(pwd, "metadata.go")); os.IsNotExist(err) {
			fmt.Errorf("Metadata file not present")
			os.Exit(1)
		}

		err = createDescriptorJSON(filepath.Join(pwd, "metadata.go"))

		if err != nil {
			fmt.Errorf("Error in creating file")
			os.Exit(1)
		}

	},
}

func init() {
	descCmd.AddCommand(syncMetadata)
}

func createDescriptorJSON(src string) error {
	var b []byte
	pwd, err := os.Getwd()

	srcData, err := ioutil.ReadFile(src)

	if err != nil {

		return err
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", srcData, 0)

	if err != nil {

		return err
	}

	fieldDetailsMap := make(map[string][]*fieldDetails)

	for key, _ := range f.Scope.Objects {

		fieldDetailsMap[key] = getKeyStructs(string(srcData), key)

	}

	if _, ok := fieldDetailsMap["Settings"]; ok {
		contribStruct := ContribStruct{}

		contribStruct.Type, err = getType(pwd)

		if err == nil {
			contribStruct.Name = path.Base(pwd) + "-" + contribStruct.Type[6:]
		}

		contribStruct.Settings = fieldDetailsMap["Settings"]
		contribStruct.Input = fieldDetailsMap["Input"]
		contribStruct.Output = fieldDetailsMap["Output"]
		contribStruct.Handler = make(map[string][]*fieldDetails)
		contribStruct.Handler["settings"] = fieldDetailsMap["HandlerSettings"]
		b, err = json.Marshal(contribStruct)

	} else {
		operationStruct := OperationStruct{}
		operationStruct.Type, err = getType(pwd)
		if err == nil {
			operationStruct.Name = path.Base(pwd) + "-" + operationStruct.Type[6:]
		}

		operationStruct.Params = fieldDetailsMap["Params"]
		operationStruct.Input = fieldDetailsMap["Input"]
		operationStruct.Output = fieldDetailsMap["Output"]

		b, err = json.Marshal(operationStruct)

	}

	err = ioutil.WriteFile("descriptor.json", []byte(jsonPrettyPrint(string(b))), os.ModePerm)
	if err != nil {
		return err
	}
	return nil

}

func getKeyStructs(src string, key string) []*fieldDetails {

	var result []*fieldDetails

	start := strings.Index(src, key+" struct {")
	offset := start + len(key) + 9

	end := strings.Index(src[offset:], "\n}")

	for _, val := range strings.Split(src[offset:offset+end], "\n") {
		var temp *fieldDetails
		if val != "" {
			fields := strings.Fields(val)

			temp = getFieldDetailStruct(fields[2])

			tempType, _ := data.ToTypeEnum(fields[1])

			if tempType.String() == "unknown" {
				temp.Type = "any"
			} else {
				temp.Type = tempType.String()
			}

			if temp != nil {
				result = append(result, temp)
			}

		}
	}

	return result
}

func getFieldDetailStruct(md string) *fieldDetails {

	components := deconstructTag(md[5:])

	temp := &fieldDetails{}

	if len(components) > 0 {

		temp.Name = strings.TrimFunc(components[0], removeSpecialChars)

		for i := 0; i < len(components); i++ {

			applyTagComponent(temp, components[i])
		}
	}

	return temp
}

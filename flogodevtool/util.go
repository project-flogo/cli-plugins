package flogodevtool

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type OperationStruct struct {
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Version     string          `json:"version"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Params      []*fieldDetails `json:"params"`
	Input       []*fieldDetails `json:"input"`
	Output      []*fieldDetails `json:"output"`
}

type ContribStruct struct {
	Name        string                     `json:"name"`
	Type        string                     `json:"type"`
	Version     string                     `json:"version"`
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	Settings    []*fieldDetails            `json:"settings"`
	Input       []*fieldDetails            `json:"input"`
	Output      []*fieldDetails            `json:"output"`
	Handler     map[string][]*fieldDetails `json:"handler"`
}

type fieldDetails struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Allowed  []string `json:"allowed"`
	Required bool     `json:"required"`
}

func applyTagComponent(details *fieldDetails, component string) {

	//process sets
	if strings.HasPrefix(component, "allowed(") {
		values := component[8 : len(component)-1]

		values = strings.TrimFunc(values, removeSpecialChars)
		details.Allowed = strings.Split(values, ",")
		return
	}

	//process flags
	switch component {
	case "required":
		details.Required = true
	}
}

func deconstructTag(str string) []string {

	var parts []string

	start := 0
	ignore := false
	for i := 0; i < len(str); i++ {

		switch str[i] {
		case '(':
			ignore = true
		case ')':
			ignore = false
		case ',':
			if !ignore {
				parts = append(parts, str[start:i])
				start = i + 1
			}
		}
	}

	parts = append(parts, str[start:])

	return parts
}

func removeSpecialChars(r rune) bool {

	if r == rune('"') || r == rune('`') || r == rune(')') {
		return true
	}
	return false
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func getType(src string) (string, error) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return "", err
	}

	for _, f := range files {

		switch f.Name() {
		case "trigger.go":
			return "flogo:trigger", nil
		case "activity.go":
			return "flogo:activity", nil
		case "operations.go":
			return "flogo:operations", nil
		}

	}
	return "", errors.New("Contrib Type not identified")
}

func copyFiles(src, dest string) error {
	_, err := os.Stat(src)
	if err != nil {
		return err
	}

	directory, _ := os.Open(src)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourceFile := filepath.Join(src, obj.Name())

		destinationFile := filepath.Join(dest, obj.Name())

		err = copy_file(sourceFile, destinationFile)
		if err != nil {
			fmt.Println(err)
		}

	}
	return nil
}

func copy_file(src string, dest string) (err error) {
	sourcFile, err := os.Open(src)
	if err != nil {
		return err
	}

	defer sourcFile.Close()

	destFile, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourcFile)

	if err == nil {
		sourceInfo, err := os.Stat(src)
		if err != nil {
			err = os.Chmod(dest, sourceInfo.Mode())
		}

	}

	return err
}

package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/mshirdel/scaffoldgo/internal"
	"github.com/spf13/cobra"
)

var _scaffoldCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new go project",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		projectRoot := filepath.Join(".", name)
		config := map[string]any{
			"ModuleName":   "github.com/mshirdel/byteland",
			"Port":         "9090",
			"RootCMD":      "sandbox",
			"RootCMDShort": "sandbox",
			"RunCMD":       "run",
			"RunCMDShort":  "run project",
		}

		files := map[string]string{
			filepath.Join(projectRoot, "", "go.mod"):             "templates/gomod",
			filepath.Join(projectRoot, "", "main.go"):            "templates/maingo",
			filepath.Join(projectRoot, "configs", "config.yaml"): "templates/configs/config.yaml",
			filepath.Join(projectRoot, "cmd", "root.go"):         "templates/cmd/rootgo",
			filepath.Join(projectRoot, "cmd", "run.go"):          "templates/cmd/rungo",
		}

		for path, template := range files {
			create(path, template, config)
		}

		return nil
	},
}

func create(destPath, templatePath string, values map[string]any) {
	err := internal.WriteTemplateFile(destPath, templatePath, values)
	if err != nil {
		fmt.Println(err)
	}
}

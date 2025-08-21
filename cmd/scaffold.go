package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/mshirdel/scaffoldgo/internal"
	"github.com/mshirdel/scaffoldgo/internal/assets"
	"github.com/spf13/cobra"
)

var (
	moduleName  string
	projectName string
	port        int
)

func init() {
	_scaffoldCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Module name like: gitbun.com/name/project_name ")
	_scaffoldCmd.Flags().StringVarP(&projectName, "name", "n", "", "Project name")
	_scaffoldCmd.Flags().IntVarP(&port, "port", "p", 8080, "Web server port")
}

var _scaffoldCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new go project",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectRoot := filepath.Join(".", projectName)
		config := map[string]any{
			"ModuleName":   moduleName,
			"Port":         port,
			"RootCMDShort": projectName,
			"RootCMD":      projectName,
			"RunCMD":       "run",
			"RunCMDShort":  "run project",
		}

		files, err := collectFiles(projectRoot)
		if err != nil {
			fmt.Println(err)
		}

		for path, template := range files {
			create(path, template, config)
		}

		return nil
	},
}

func collectFiles(projectRoot string) (map[string]string, error) {
	files := make(map[string]string)

	err := fs.WalkDir(assets.Templates, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel("templates", path)
		if err != nil {
			return err
		}

		dest := filepath.Join(projectRoot, relPath)

		files[dest] = path
		return nil
	})

	return files, err
}

func create(destPath, templatePath string, values map[string]any) {
	err := internal.WriteTemplateFile(destPath, templatePath, values)
	if err != nil {
		fmt.Println(err)
	}
}

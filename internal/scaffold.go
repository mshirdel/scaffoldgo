package internal

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/mshirdel/scaffoldgo/internal/assets"
)

func WriteTemplateFile(destPath string, templatePath string, params map[string]any) error {
	dir := filepath.Dir(destPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("make all dir: %w", err)
	}

	content, err := assets.Templates.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("read template file: %w", err)
	}

	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(content))
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	file, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("create destination file: %w", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, params); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	return nil
}

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	_rootCMD = cobra.Command{
		Use:   "{{.projectName}}",
		Short: "{{.projectName}}",
	}

	_configPath string
)

func init() {
	_rootCMD.AddCommand(_runCmd)
	_rootCMD.AddCommand(_serve)
}

func Execute() {
	if err := _rootCMD.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

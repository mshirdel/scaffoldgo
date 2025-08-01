package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	_rootCMD = cobra.Command{
		Use:   "sandbox",
		Short: "Sandbox",
	}

	_configPath string
)

func init() {
	_rootCMD.AddCommand(_scaffoldCmd)
}

func Execute() {
	if err := _rootCMD.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

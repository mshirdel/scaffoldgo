package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var _runCmd = &cobra.Command{
	Use:   "run",
	Short: "run project",
	RunE: func(cmd *cobra.Command, args []string) error {
		run()
		return nil
	},
}

func run() {
	fmt.Println("test is ok")
}

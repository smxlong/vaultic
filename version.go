package vaultic

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "0.0.1"

// versionCommand returns the version command.
func versionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(Version)
			return nil
		},
	}
	return cmd
}

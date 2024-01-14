package vaultic

import (
	"github.com/spf13/cobra"

	"github.com/smxlong/vaultic/server"
)

// Vaultic is the vaultic application.
type Vaultic struct {
}

// New returns a new Vaultic application.
func New() *Vaultic {
	return &Vaultic{}
}

// Command returns the main cobra.Command for the application.
func (v *Vaultic) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vaultic",
		Short: "Vaultic is a JWT auth server backed by Vault.",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := server.New()
			if err != nil {
				return err
			}
			return s.Run()
		},
	}
	cmd.AddCommand(versionCommand())
	return cmd
}

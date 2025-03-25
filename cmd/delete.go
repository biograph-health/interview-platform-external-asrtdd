package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

type deleteCmd struct {
	cmd *cobra.Command
}

func newDeleteCmd() *deleteCmd {
	root := &deleteCmd{}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Destroy provisioned infrastructure",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: implement delete functionality
			return errors.New("not implemented")
		},
	}

	root.cmd = cmd
	return root
}

package cmd

import (
	"errors"
	
	"github.com/spf13/cobra"
)

type createCmd struct {
	cmd *cobra.Command
}

func newCreateCmd() *createCmd {
	root := &createCmd{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Provision infrastructure based on configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: implement create functionality
			return errors.New("not implemented")
		},
	}

	root.cmd = cmd
	return root
}

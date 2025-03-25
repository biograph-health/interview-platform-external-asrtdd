package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type deleteCmd struct {
	cmd    *cobra.Command
	config string
}

func newDeleteCmd() *deleteCmd {
	root := &deleteCmd{}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete infrastructure",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: implement delete functionality
			fmt.Println("delete not implemented")

			return nil
		},
	}

	cmd.Flags().StringVarP(&root.config, "config", "f", "config.yaml", "Load configuration from file")
	_ = cmd.MarkFlagFilename("config", "yaml", "yml")

	root.cmd = cmd
	return root
}

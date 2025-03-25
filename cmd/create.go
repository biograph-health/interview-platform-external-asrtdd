package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type createCmd struct {
	cmd    *cobra.Command
	config string
}

func newCreateCmd() *createCmd {
	root := &createCmd{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create infrastructure",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: implement create functionality
			fmt.Println("create not implemented")

			return nil
		},
	}

	cmd.Flags().StringVarP(&root.config, "config", "f", "config.yaml", "Load configuration from file")
	_ = cmd.MarkFlagFilename("config", "yaml", "yml")

	root.cmd = cmd
	return root
}

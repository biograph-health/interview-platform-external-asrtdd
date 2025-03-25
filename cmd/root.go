package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

type rootCmd struct {
	cmd *cobra.Command
}

func Execute(args []string) {
	rootCmd := newRootCmd()
	rootCmd.cmd.SetArgs(args)
	if err := rootCmd.cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func newRootCmd() *rootCmd {
	root := &rootCmd{}
	cmd := &cobra.Command{
		Use:               "infra-cli",
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}
	cmd.AddCommand(
		newCreateCmd().cmd,
		newDeleteCmd().cmd,
	)
	root.cmd = cmd
	return root
}

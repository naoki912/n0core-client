package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newVMCommand())
}

func newVMCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vm",
		Short: "vm command",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newCreateVMRequestCmd(),
	)

	return cmd
}

func newCreateVMRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "CreateVMRequestCmd",
		RunE:  runCreateVMRequestCmd,
	}

	return cmd
}

func runCreateVMRequestCmd(cmd *cobra.Command, args []string) error {
	return NewN0coreClient().CreateVMRequest()
}

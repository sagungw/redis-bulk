package cli

import "github.com/spf13/cobra"

var (
	RootCmd = &cobra.Command{
		RunE: rootRunE,
	}
)

func rootRunE(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

func init() {
	RootCmd.AddCommand(SingleCmd, SeedCmd)
}

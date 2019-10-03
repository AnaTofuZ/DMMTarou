package cmd

import "github.com/spf13/cobra"

func NewDMMTarouCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "dmmtarou",
	}

	cmd.AddCommand()
	return cmd
}

// The init command. It writes the first arch.yaml, either from observed edges or
// from candidate modules suggested by the exploratory path.

package main

import "github.com/spf13/cobra"

func newInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Write an initial arch.yaml",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented
		},
	}

	cmd.Flags().Bool("observe", false, "generate an allow-list from currently observed edges")
	// generate reads community detection output, which is seed dependent. It is a
	// starting point for a human to edit, never a checked in result.
	cmd.Flags().Bool("generate", false, "suggest module boundaries from community detection")

	return cmd
}

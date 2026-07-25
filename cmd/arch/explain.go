// The explain command. It answers a single question about the artifact on the
// terminal, using the same queries the MCP server exposes.

package main

import "github.com/spf13/cobra"

func newExplainCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "explain <symbol>",
		Short: "Explain the callers, boundary crossings and blast radius of a symbol",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented
		},
	}

	cmd.Flags().String("sha", "", "commit SHA to read, defaults to HEAD")
	cmd.Flags().Bool("json", false, "emit JSON instead of text")

	return cmd
}

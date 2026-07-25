// The mcp command. It runs the MCP server as a subprocess over stdio, so nothing
// but MCP messages may be written to stdout. Logging goes to stderr.

package main

import "github.com/spf13/cobra"

func newMCPCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Serve the artifact to a model over MCP on stdio",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented
		},
	}

	cmd.Flags().String("sha", "", "commit SHA to serve, defaults to HEAD")

	return cmd
}

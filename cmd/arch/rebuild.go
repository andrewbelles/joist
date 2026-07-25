// The rebuild command. Rebuild tiers are incremental by default, full on a trunk
// merge, and explicit when a lockfile, build config, toolchain or indexer version
// changes, since those invalidate keys that content hashing alone does not catch.

package main

import "github.com/spf13/cobra"

func newRebuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rebuild",
		Short: "Rebuild the graph for the current commit",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented
		},
	}

	cmd.Flags().Bool("rebuild-all", false, "ignore the cache and rebuild every document")

	return cmd
}

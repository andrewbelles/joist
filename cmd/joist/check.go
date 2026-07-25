// The check command. It is the CI entry point and the only command that decides
// an exit code from findings.

package main

import "github.com/spf13/cobra"

func newCheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check the working tree against the declared architecture",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented
		},
	}

	// inherit-parent keeps CI green when a commit does not build. Without it a
	// broken build reads as an empty graph, which reads as every edge deleted.
	cmd.Flags().Bool("inherit-parent", false, "reuse the parent artifact when this commit does not build")
	cmd.Flags().String("baseline", "", "artifact or ref to diff against, defaults to the merge base")
	cmd.Flags().Bool("warn-only", false, "report violations without failing")

	return cmd
}

// The viewer command. It serves the embedded SPA on loopback.

package main

import (
	"fmt"

	"github.com/andrewbelles/call-graph-tooling/internal/viewer"
	"github.com/spf13/cobra"
)

func newViewerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "viewer",
		Short: "Serve the read only viewer on loopback",
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cmd.Flags().GetString("addr")
			if err != nil {
				return err
			}
			sha, err := cmd.Flags().GetString("sha")
			if err != nil {
				return err
			}

			srv := viewer.New(addr, sha)
			fmt.Fprintf(cmd.ErrOrStderr(), "arch viewer on http://%s (spa=%t)\n", srv.Addr, viewer.Built)
			return srv.ListenAndServe()
		},
	}

	cmd.Flags().String("addr", viewer.DefaultAddr, "loopback address to bind")
	cmd.Flags().String("sha", "", "commit SHA to serve, defaults to HEAD")

	return cmd
}

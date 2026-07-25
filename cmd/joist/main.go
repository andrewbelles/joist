// Command joist analyzes a build level call graph and checks it against a
// declared architecture. This package is wiring only: flag definitions, command
// tree assembly, and exit codes. All behavior lives under internal.

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// version is set at link time by the build. It is part of the cache key, so a
// release must not ship the dev default.
var version = "dev"

// errNotImplemented is returned by every command in the scaffold.
var errNotImplemented = errors.New("not implemented")

func main() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "joist:", err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "joist",
		Short:         "Build level call graph and architecture conformance",
		Version:       version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.PersistentFlags().String("config", "arch.yaml", "path to the declared architecture")
	root.PersistentFlags().String("cache", ".joist/cache", "path to the local artifact cache")

	root.AddCommand(
		newCheckCmd(),
		newInitCmd(),
		newExplainCmd(),
		newRebuildCmd(),
		newMCPCmd(),
		newViewerCmd(),
	)
	return root
}

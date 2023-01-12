package mrf

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCMD = &cobra.Command{
	Use:     "mrf",
	Short:   "Music Record Fetcher - a CLI to get information about the records of your favorites artists",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: There was an error executing mrf '%s'", err)
		os.Exit(1)
	}
}

package mrf

import (
	"fmt"
	"os"

	"github.com/sevens7xix/mrf/internal/mrf"
	"github.com/spf13/cobra"
)

var (
	version       = "0.0.1"
	bestTracks    bool
	printResponse bool
)

var rootCMD = &cobra.Command{
	Use:     "mrf",
	Short:   "Music Record Fetcher - a CLI to get information about the records of your favorites artists",
	Version: version,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mrf.GetTracks(bestTracks, printResponse, args)
	},
}

func Execute() {
	rootCMD.PersistentFlags().BoolVarP(&bestTracks, "best", "b", false, "To get the best tracks from the desired artist")
	rootCMD.PersistentFlags().BoolVarP(&printResponse, "print", "p", false, "To print the fetch result on a file")

	if err := rootCMD.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: There was an error executing mrf '%s'", err)
		os.Exit(1)
	}
}

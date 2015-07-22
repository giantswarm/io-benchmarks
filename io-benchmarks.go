package main

import (
	"github.com/giantswarm/io-benchmarks/utils"
	"github.com/spf13/cobra"
)

var (
	ioBenchmarksCmd = &cobra.Command{
		Use:   "io-benchmarks",
		Short: "CLI control panel for Quobyte",
		Long:  "CLI cntrol panel for Quobyte",
		Run:   ioBenchmarksRun,
	}
)

func init() {
	ioBenchmarksCmd.PersistentFlags().BoolVarP(&utils.GlobalFlags.Debug, "debug", "d", false, "Print debug output")
	ioBenchmarksCmd.PersistentFlags().BoolVarP(&utils.GlobalFlags.Verbose, "verbose", "v", false, "Print verbose output")
}

func ioBenchmarksRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func main() {
	ioBenchmarksCmd.AddCommand(versionCmd)
	ioBenchmarksCmd.Execute()
}

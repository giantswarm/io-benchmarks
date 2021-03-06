package main

import (
	"github.com/giantswarm/io-benchmarks/command"
	"github.com/giantswarm/io-benchmarks/utils"
	"github.com/spf13/cobra"
)

var (
	ioBenchmarksCmd = &cobra.Command{
		Use:   "io-benchmarks",
		Short: "CLI tool to run I/O benchmarks based on fio",
		Long:  "CLI tool to run I/O benchmarks based on fio",
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
	ioBenchmarksCmd.AddCommand(command.RunCmd)
	ioBenchmarksCmd.Execute()
}

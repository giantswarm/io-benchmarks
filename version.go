package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectVersion = "dev"

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show cli version",
		Long:  "Show cli version",
		Run:   versionRun,
	}
)

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Println("io-benchmarks version", projectVersion)
}

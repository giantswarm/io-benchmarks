package command

import (
	"fmt"

	"github.com/giantswarm/io-benchmarks/fio"
	"github.com/giantswarm/io-benchmarks/utils"

	"github.com/spf13/cobra"
)

var (
	RunCmd = &cobra.Command{
		Use:   "run",
		Short: "Run benchmark tests",
		Long:  "Run specific parts of or the whole benchmark suite",
		Run:   runRun,
	}

	runTestCmd = &cobra.Command{
		Use:   "test [test name]",
		Short: "Run a specific test of the test suite",
		Long:  "Run a specific test of the test suite",
		Run:   runTestRun,
	}

	runFlags struct {
		TestsDirectory   string
		WorkingDirectory string
		OutputDirectory  string
		SummaryFilename  string

		GenerateBandwithStats bool
		GenerateIOPSStats     bool
		GenerateLatencyStats  bool
	}
)

func init() {
	RunCmd.PersistentFlags().StringVar(&runFlags.TestsDirectory, "tests-directory", "./tests", "Directory to search for test files")
	RunCmd.PersistentFlags().StringVar(&runFlags.WorkingDirectory, "working-directory", "./.io-benchmark", "Directory to perform benchmarks in")
	RunCmd.PersistentFlags().StringVar(&runFlags.OutputDirectory, "output-directory", "./io-benchmark-results", "Directory to store results to")
	RunCmd.PersistentFlags().StringVar(&runFlags.SummaryFilename, "summary-filename", "summary.log", "Filename of the test run summary")

	RunCmd.PersistentFlags().BoolVar(&runFlags.GenerateBandwithStats, "generate-bandwith-stats", true, "Generate bandwith stats for plots")
	RunCmd.PersistentFlags().BoolVar(&runFlags.GenerateIOPSStats, "generate-iops-stats", true, "Generate IOPS stats for plots")
	RunCmd.PersistentFlags().BoolVar(&runFlags.GenerateLatencyStats, "generate-latency-stats", true, "Generate latency stats for plots")

	RunCmd.AddCommand(runTestCmd)
}

func runRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func runTestRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 || args[0] == "" {
		utils.ExitStderrf("You have to provide a test name as argument.")
	}
	test := args[0]

	fioConf := fio.FioConfiguration{
		JobDirectory:         runFlags.TestsDirectory,
		WorkingDirectory:     runFlags.WorkingDirectory,
		OutputFilename:       runFlags.SummaryFilename,
		LogsDirectory:        runFlags.OutputDirectory,
		LogsPrefix:           "benchmark",
		GenerateBandwithLogs: runFlags.GenerateBandwithStats,
		GenerateIOPSLogs:     runFlags.GenerateIOPSStats,
		GenerateLatencyLogs:  runFlags.GenerateLatencyStats,
	}
	fioRunner, err := fio.NewFioRunner(fioConf)

	if err != nil {
		utils.ExitStderr(err)
	}

	if output, err := fioRunner.RunTest(test); err != nil {
		utils.ExitStderr(err)
	} else {
		fmt.Print(output)
	}

	fio2gnuplotConf := fio.Fio2GNUPlotConfiguration{
		LogsDirectory: runFlags.OutputDirectory,
		LogsPatterns: []string{
			"*_bw.*.log",
			"*_iops.*.log",
			"*_lat.*.log",
			"*_clat.*.log",
		},
	}
	fio2gnuplotRunner, err := fio.NewFio2GNUPlotRunner(fio2gnuplotConf)

	if err := fio2gnuplotRunner.RunPlots(); err != nil {
		utils.ExitStderr(err)
	}
}

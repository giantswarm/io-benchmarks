package command

import (
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
		DirectMode       bool
		TestsDirectory   string
		WorkingDirectory string
		OutputDirectory  string

		GenerateBandwithStats bool
		GenerateIOPSStats     bool
		GenerateLatencyStats  bool
	}
)

func init() {
	RunCmd.PersistentFlags().BoolVar(&runFlags.DirectMode, "direct-mode", true, "Use direct mode to bypass Kernel I/O buffers")
	RunCmd.PersistentFlags().StringVar(&runFlags.TestsDirectory, "tests-directory", "./tests", "Directory to search for test files")
	RunCmd.PersistentFlags().StringVar(&runFlags.WorkingDirectory, "working-directory", "./.io-benchmark", "Directory to perform benchmarks in")
	RunCmd.PersistentFlags().StringVar(&runFlags.OutputDirectory, "output-directory", "./io-benchmark-results", "Directory to store results to")

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

	c := fio.Configuration{
		JobDirectory:         runFlags.TestsDirectory,
		WorkingDirectory:     runFlags.WorkingDirectory,
		DirectMode:           runFlags.DirectMode,
		LogsDirectory:        runFlags.OutputDirectory,
		GenerateBandwithLogs: runFlags.GenerateBandwithStats,
		GenerateIOPSLogs:     runFlags.GenerateIOPSStats,
		GenerateLatencyLogs:  runFlags.GenerateLatencyStats,
	}
	fio, err := fio.NewFioRunner(c)

	if err != nil {
		utils.ExitStderr(err)
	}

	if err := fio.RunTest(test); err != nil {
		utils.ExitStderr(err)
	}
}

package fio

import (
	"os/exec"

	"github.com/giantswarm/io-benchmarks/utils"
)

func executableExists(executable string) bool {
	cmd := exec.Command("which", executable)

	if err := cmd.Run(); err != nil {
		utils.DebugError(err)
		return false
	}
	return true
}

func fioExists() bool {
	return executableExists("fio")
}

func fio2GNUPlotExists() bool {
	return executableExists("fio2gnuplot")
}

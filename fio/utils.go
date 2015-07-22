package fio

import (
	"os/exec"

	"github.com/giantswarm/io-benchmarks/utils"
)

func fioExists() bool {
	cmd := exec.Command("which", "fio")

	if err := cmd.Run(); err != nil {
		utils.DebugError(err)
		return false
	}
	return true
}

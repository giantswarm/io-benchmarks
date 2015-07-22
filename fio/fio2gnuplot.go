package fio

import (
	"path/filepath"

	"github.com/giantswarm/io-benchmarks/exec"

	"github.com/juju/errgo"
)

type Fio2GNUPlotConfiguration struct {
	LogsDirectory string
}

type Fio2GNUPlotRunner struct {
	conf Fio2GNUPlotConfiguration
}

func NewFio2GNUPlotRunner(c Fio2GNUPlotConfiguration) (Fio2GNUPlotRunner, error) {
	var err error

	if !fio2GNUPlotExists() {
		return Fio2GNUPlotRunner{}, errgo.Newf("Cannot locate fio2gnuplot. Looks like it is not installed on your system.")
	}

	if c.LogsDirectory, err = filepath.Abs(c.LogsDirectory); err != nil {
		return Fio2GNUPlotRunner{}, errgo.Mask(err)
	}

	return Fio2GNUPlotRunner{
		conf: c,
	}, nil
}

func (r Fio2GNUPlotRunner) RunPlots() error {
	patterns := []string{
		"*_bw.*.log",
		"*_iops.*.log",
		"*_lat.*.log",
		"*_clat.*.log",
	}

	for _, pattern := range patterns {
		if err := r.runPlotWithPattern(pattern); err != nil {
			return errgo.Mask(err)
		}
	}

	return nil
}

func (r Fio2GNUPlotRunner) runPlotWithPattern(pattern string) error {
	var cmdArguments []string

	cmdArguments = append(cmdArguments, "-p")
	cmdArguments = append(cmdArguments, pattern)
	cmdArguments = append(cmdArguments, "-g")

	return errgo.Mask(exec.RunCommand("fio2gnuplot", cmdArguments, r.conf.LogsDirectory))
}

package fio

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/giantswarm/io-benchmarks/exec"

	"github.com/juju/errgo"
)

type Configuration struct {
	JobDirectory     string
	WorkingDirectory string
	DirectMode       bool

	LogsDirectory        string
	GenerateBandwithLogs bool
	GenerateIOPSLogs     bool
	GenerateLatencyLogs  bool
}

type FioRunner struct {
	conf Configuration
}

func NewFioRunner(c Configuration) (FioRunner, error) {
	var err error

	if !fioExists() {
		return FioRunner{}, errgo.Newf("Cannot locate fio. Looks like it is not installed on your system.")
	}

	if c.JobDirectory, err = filepath.Abs(c.JobDirectory); err != nil {
		return FioRunner{}, errgo.Mask(err)
	}

	if c.WorkingDirectory, err = filepath.Abs(c.WorkingDirectory); err != nil {
		return FioRunner{}, errgo.Mask(err)
	}

	return FioRunner{
		conf: c,
	}, nil
}

func (r FioRunner) RunTest(test string) error {
	var cmdArguments []string

	if r.conf.DirectMode {
		cmdArguments = append(cmdArguments, "--direct=1")
	}

	if r.conf.WorkingDirectory != "" {
		cmdArguments = append(cmdArguments, "--directory="+r.conf.WorkingDirectory)
	}

	if r.conf.GenerateBandwithLogs {
		cmdArguments = append(cmdArguments, fmt.Sprintf("--write_bw_log=%s/benchmark", r.conf.LogsDirectory))
	}

	if r.conf.GenerateIOPSLogs {
		cmdArguments = append(cmdArguments, fmt.Sprintf("--write_iops_log=%s/benchmark", r.conf.LogsDirectory))
	}

	if r.conf.GenerateLatencyLogs {
		cmdArguments = append(cmdArguments, fmt.Sprintf("--write_lat_log=%s/benchmark", r.conf.LogsDirectory))
	}

	testfilePath := fmt.Sprintf("%s/%s", r.conf.JobDirectory, test)
	cmdArguments = append(cmdArguments, testfilePath)

	if err := r.createWorkingDirectory(); err != nil {
		return errgo.Mask(err)
	}

	runErr := exec.RunCommand("fio", cmdArguments, "")

	if err := r.removeWorkingDirectory(); err != nil {
		return errgo.Mask(err)
	}

	return errgo.Mask(runErr)
}

func (r FioRunner) createWorkingDirectory() error {
	f, err := os.Open(r.conf.WorkingDirectory)

	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(r.conf.WorkingDirectory, 0755)
			return nil
		} else {
			return errgo.Mask(err)
		}
	}

	if fi, err := f.Stat(); err != nil {
		return errgo.Mask(err)
	} else if !fi.IsDir() {
		return errgo.Newf("Working directory '%s' exists but appears to be a file.", r.conf.WorkingDirectory)
	}

	if fis, err := f.Readdir(0); err != nil {
		return errgo.Mask(err)
	} else if len(fis) > 0 {
		return errgo.Newf("Working directory '%s' appears to be not empty.", r.conf.WorkingDirectory)
	}

	return nil
}

func (r FioRunner) removeWorkingDirectory() error {
	fi, err := os.Stat(r.conf.WorkingDirectory)

	if err != nil {
		return errgo.Mask(err)
	}

	if !fi.IsDir() {
		return errgo.Newf("Working directory '%s' exists but appears to be a file.", r.conf.WorkingDirectory)
	} else {
		return os.RemoveAll(r.conf.WorkingDirectory)
	}

	return nil
}

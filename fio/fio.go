package fio

import (
	"fmt"
	"os"

	"github.com/giantswarm/io-benchmarks/exec"

	"github.com/juju/errgo"
)

type Configuration struct {
	JobDirectory     string
	WorkingDirectory string
	DirectMode       bool
}

type FioRunner struct {
	conf Configuration
}

func NewFioRunner(c Configuration) (FioRunner, error) {
	if !fioExists() {
		return FioRunner{}, errgo.Newf("Cannot locate fio. Looks like it is not installed on your system.")
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

	testfilePath := fmt.Sprintf("%s/%s", r.conf.JobDirectory, test)
	cmdArguments = append(cmdArguments, testfilePath)

	if err := r.createWorkingDirectory(); err != nil {
		return errgo.Mask(err)
	}

	return exec.RunCommand("fio", cmdArguments, r.conf.WorkingDirectory)
}

func (r FioRunner) createWorkingDirectory() error {
	fi, err := os.Stat(r.conf.WorkingDirectory)

	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(r.conf.WorkingDirectory, 0755)
			return nil
		} else {
			return errgo.Mask(err)
		}
	}

	if !fi.IsDir() {
		return errgo.Newf("Working directory '%s' exists and appears to be a file.", r.conf.WorkingDirectory)
	}

	return nil
}

package fio

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/giantswarm/io-benchmarks/exec"

	"github.com/juju/errgo"
)

type FioConfiguration struct {
	JobDirectory     string
	WorkingDirectory string

	LogsDirectory        string
	LogsPrefix           string
	OutputFilename       string
	GenerateBandwithLogs bool
	GenerateIOPSLogs     bool
	GenerateLatencyLogs  bool
}

type FioRunner struct {
	conf FioConfiguration
}

func NewFioRunner(c FioConfiguration) (FioRunner, error) {
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

	if c.LogsDirectory, err = filepath.Abs(c.LogsDirectory); err != nil {
		return FioRunner{}, errgo.Mask(err)
	}

	return FioRunner{
		conf: c,
	}, nil
}

func (r FioRunner) RunTest(test string) (string, error) {
	var cmdArguments []string

	if r.conf.WorkingDirectory != "" {
		cmdArguments = append(cmdArguments, "--directory="+r.conf.WorkingDirectory)
	}

	if r.conf.GenerateBandwithLogs {
		cmdArguments = append(cmdArguments, fmt.Sprintf("--write_bw_log=%s/%s", r.conf.LogsDirectory, r.conf.LogsPrefix))
	}

	if r.conf.GenerateIOPSLogs {
		cmdArguments = append(cmdArguments, fmt.Sprintf("--write_iops_log=%s/%s", r.conf.LogsDirectory, r.conf.LogsPrefix))
	}

	if r.conf.GenerateLatencyLogs {
		cmdArguments = append(cmdArguments, fmt.Sprintf("--write_lat_log=%s/%s", r.conf.LogsDirectory, r.conf.LogsPrefix))
	}

	cmdArguments = append(cmdArguments, fmt.Sprintf("--output=%s/%s", r.conf.LogsDirectory, r.conf.OutputFilename))

	testfilePath := fmt.Sprintf("%s/%s", r.conf.JobDirectory, test)
	cmdArguments = append(cmdArguments, testfilePath)

	if err := r.createWorkingDirectory(); err != nil {
		return "", errgo.Mask(err)
	}

	runErr := exec.RunCommand("fio", cmdArguments, "")

	if err := r.removeWorkingDirectory(); err != nil {
		return "", errgo.Mask(err)
	}

	if runErr != nil {
		return "", errgo.Mask(runErr)
	}

	output, err := r.getOutputLog()

	if err != nil {
		return "", errgo.Mask(err)
	}

	return output, nil
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

func (r FioRunner) getOutputLog() (string, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", r.conf.LogsDirectory, r.conf.OutputFilename))
	if err != nil {
		errgo.Mask(err)
	}

	return string(content[:]), nil
}

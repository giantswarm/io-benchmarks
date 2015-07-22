package exec

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/giantswarm/io-benchmarks/utils"
)

type CommandExecutionError struct {
	Executable string
	Arguments  []string
	Stdout     bytes.Buffer
	Stderr     bytes.Buffer
}

func (e CommandExecutionError) Error() string {
	return fmt.Sprintf("command '%s' returned error: %s", e.Executable, e.Stderr.String())
}

func RunCommand(executable string, arguments []string, workingDirectory string) error {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(executable, arguments...)
	cmd.Stderr = &stderr

	if utils.GlobalFlags.Verbose {
		cmd.Stdout = io.MultiWriter(&stdout, os.Stdout)
	} else {
		cmd.Stdout = os.Stdout
	}

	cmd.Dir = workingDirectory

	utils.Verbosef("running command in directory '%s': %s %s", cmd.Dir, executable, strings.Join(arguments, " "))
	err := cmd.Run()

	if err != nil {
		return CommandExecutionError{
			Executable: executable,
			Arguments:  arguments,
			Stdout:     stdout,
			Stderr:     stderr,
		}
	}

	return nil

}

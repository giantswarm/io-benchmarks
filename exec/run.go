package exec

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/giantswarm/io-benchmarks/utils"

	"github.com/juju/errgo"
)

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

	err := cmd.Run()

	if err != nil {
		errMsg := fmt.Sprintf("The following command failed:\n\n\t%s", strings.Join(cmd.Args, " "))

		if stdout.Len() > 0 {
			errMsg += fmt.Sprintf("\n\nOuput of stdout was:\n\n\t%s", stdout.String())
		}

		if stderr.Len() > 0 {
			errMsg += fmt.Sprintf("\n\nOutput of stderr was:\n\n\t%s", stderr.String())
		}

		utils.Stderrf(errMsg)
		return errgo.New("Running fio did not succeed")
	}

	return nil

}

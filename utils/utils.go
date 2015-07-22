package utils

import (
	"bufio"
	"fmt"
	"os"
)

var (
	GlobalFlags struct {
		Debug   bool
		Verbose bool
	}
)

// ExitError exits the process with 1 as return code
func ExitError() {
	os.Exit(1)
}

// ExitSuccess exits the process with 0 as return code
func ExitSuccess() {
	os.Exit(0)
}

// Exit process with code 1 and print output to Stderr, using the error message
// of err. If --debug is set to true, a helpful error will be printed.
func ExitStderr(err error) {
	debug(err)
	stderrf(err.Error())
	ExitError()
}

// Exit process with code 1 and print output to Stderr, using all objects of v
// in format f. If --debug is set to true, helpful errors will be printed if
// given.
func ExitStderrf(f string, v ...interface{}) {
	debug(v...)
	stderrf(f, v...)
	ExitError()
}

// Stderrf prints output to Stderr, using all objects of v in format f. If
// --debug is set to true, helpful errors will be printed if given.
func Stderrf(f string, v ...interface{}) {
	debug(v...)
	stderrf(f, v...)
}

// Exit process with code 0 and print output to Stdout, using all objects of v
// in format f.
func ExitStdoutf(f string, v ...interface{}) {
	stdoutf(f, v...)
	ExitSuccess()
}

// Print verbose output for the user to Stdout, using all objects of v in
// format f.
func Verbosef(f string, v ...interface{}) {
	if GlobalFlags.Verbose {
		stdoutf(f, v...)
	}
}

// Print output interesting for the user to Stdout, using all objects of v in
// format f.
func Stdoutf(f string, v ...interface{}) {
	stdoutf(f, v...)
}

// Confirm repeatedly asks the user a question until he confirms it with the string "yes"
func Confirm(question string) error {
	for {
		fmt.Printf("%s Enter 'yes': ", question)
		bio := bufio.NewReader(os.Stdin)
		line, _, err := bio.ReadLine()
		if err != nil {
			return err
		}

		if string(line) == "yes" {
			return nil
		}
		fmt.Println("Please enter 'yes' to confirm.")
	}
}

// DebugError prints debug output of error values when debugging mode is turned on
func DebugError(e error) {
	debug(e)
}

////////////////////////////////////////////////////////////////////////////////
// private

func debug(v ...interface{}) {
	if GlobalFlags.Debug {
		for _, obj := range v {
			if _, isErr := obj.(error); isErr {
				fmt.Printf("DEBUG: %#v", obj)
				fmt.Println()
			}
		}
	}
}

func stderrf(f string, v ...interface{}) {
	if !GlobalFlags.Debug {
		printf(os.Stderr, f, v...)
	}
}

func stdoutf(f string, v ...interface{}) {
	printf(os.Stdout, f, v...)
}

func printf(file *os.File, f string, v ...interface{}) {
	if f == "" {
		return
	}

	fmt.Fprintf(file, f, v...)
	fmt.Println()
}

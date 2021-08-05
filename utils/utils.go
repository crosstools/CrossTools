// Package utils provides small utility functions to make coding for CrossTools
// easier.
package utils

import (
	"fmt"
	"os"
)

// Check looks over the err to check if it is not nil.
// If it is not nil, the err panics.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// CliError formats using the default formats for its operands and writes to os.Stderr.
// NAME+" error: " is written to os.Stderr before CliError prints the arguments.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func CliError(a ...interface{}) (n int, err error) {
	return fmt.Fprint(os.Stderr, NAME, "error:", a)
}

// CliErrorln formats using the default formats for its operands and writes to os.Stderr.
// NAME+" error: " is written to os.Stderr before CliErrorln prints the arguments.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func CliErrorln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stderr, NAME, "error:", a)
}

// CliErrorf formats according to a format specifier and writes to w.
// NAME+" error: " is written to os.Stderr before CliErrorf prints the arguments.
// It returns the number of bytes written and any write error encountered.
func CliErrorf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stderr, NAME+" error: "+format, a...)
}

// CommandNotImplementedError prints a descriptive error message of the command not being implemented.
// It panics if there is an err.
func CommandNotImplementedError(command string) {
	if _, err := CliErrorf("Command '%s' is not yet implemented", command); err != nil {
		panic(err)
	}
}

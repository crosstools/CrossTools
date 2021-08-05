// Package commands provides the functions for the commands
// in the CLI.
package commands

import (
	"os"

	"github.com/crosstools/crosstools/utils"
)

// Remove removes a file or directory, according to the arg (name).
// arg is the name of the file or directory.
// If arg == "", no file/directory will be removed.
func Remove(arg string) {
	if arg != "" {
		err := os.Remove(arg)
		utils.Check(err)
	} else {
		utils.CliErrorln("Filename must be included for remove command") // If this function name changes,
		// CHANGE THIS STRING ERROR TOO
	}
}

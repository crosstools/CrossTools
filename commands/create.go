// Package commands provides the functions for the commands
// in the CLI.
package commands

import (
	"os"

	"github.com/crosstools/crosstools/utils"
)

// Create makes a file or directory, according to the arg (name).
// arg is the name of the file or directory.
// If arg == "", no file/directory will be made.
// If directory == true, directory will be made.
// If directory == false, a file will be made.
func Create(arg string, directory bool) {
	if arg != "" {
		if !directory {
			_, err := os.Create(arg)
			utils.Check(err)
		} else {
			utils.Check(os.Mkdir(arg, 0755))
		}
	}
}

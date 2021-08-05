package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/crosstools/crosstools/commands"
	"github.com/crosstools/crosstools/utils"
	"github.com/fatih/color"
	"github.com/pborman/getopt/v2"
)

// Deprecated: Prints to os.Stderr with NAME + " error: " and the error itself. Use utils.CliErrorln instead!
func cliErrorln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stderr, utils.NAME, "error:", a)
}

// Flags
var (
	versionFlag   = getopt.BoolLong("version", 'v', "", "Get the current version of "+utils.NAME)
	helpFlag      = getopt.BoolLong("help", 'h', "", "Display help/usage")
	directoryFlag = getopt.BoolLong("dir", 'd', "", "Directory instead of filename")
)

func Usage() {
	usage := `
Commands of %s:
  install
        Install crosstools into system
  update
        Update crosstools in system
  create <filename>
        Create a new file with <filename>
  -d create <directory>
        Create a new directory (using the -d or --dir flag) with <directory>
  remove <filename/directory>
        Removes the <filename/directory>

Deprecated commands:
  newfile <filename>
        Create a new file with the <filename>
`
	getopt.PrintUsage(os.Stderr)
	fmt.Fprintf(os.Stderr, usage, os.Args[0])
	color.Set(color.FgRed, color.Bold)
	defer color.Unset()
	fmt.Fprintln(os.Stderr,
		"\nNote:\n  If you're going to use flags, please use them before using the command. e.g., \"crosstools -d create hello\" instead of \"crosstools create -d hello\"")
}

func main() {
	// Set custom usage
	getopt.SetUsage(Usage)

	// Parse the program arguments
	getopt.ParseV2()
	// Get the remaining positional parameters
	// args := getopt.Args()

	// Deprecated: check panics the err if err is not nil. Please use utils.Check instead!
	check := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	if *versionFlag {
		fmt.Printf("%s version %s %s\n", utils.NAME, utils.VERSION, runtime.GOOS)
		return
	}

	if *helpFlag || len(os.Args) == 1 {
		getopt.Usage()
		return
	}

	switch getopt.Arg(0) {
	case "install", "update":
		utils.CommandNotImplementedError(getopt.Arg(0))
	case "create":
		commands.Create(getopt.Arg(1), *directoryFlag)
	case "remove":
		commands.Remove(getopt.Arg(1))

	// Deprecated commands
	case "newfile":
		filename := getopt.Arg(1)

		if filename != "" {
			_, err := os.Create(filename)
			check(err)
		} else {
			cliErrorln("Filename must be included for", getopt.Arg(0), "command")
		}

	default:
		utils.CliErrorf("Unknown command '%s'\n", getopt.Arg(0))
	}

}

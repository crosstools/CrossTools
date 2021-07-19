package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/pborman/getopt/v2"
)

// Global information
const (
	NAME    = "crosstools"
	VERSION = "1.0.0"
)

// Flags
var (
	versionFlag   = getopt.BoolLong("version", 'v', "", "Get the current version of "+NAME)
	helpFlag      = getopt.BoolLong("help", 'h', "", "Display help/usage")
	directoryFlag = getopt.BoolLong("dir", 'd', "", "Directory instead of filename")
)

// func cliError(a ...interface{}) (n int, err error) {
// 	return fmt.Fprint(os.Stderr, NAME, "error:", a)
// }

func cliErrorln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stderr, NAME, "error:", a)
}

func cliErrorf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stderr, NAME+" error: "+format, a...)
}

func commandNotImplementedError(command string) {
	cliErrorf("Command '%s' is not yet implemented", command)
}

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

	check := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	if *versionFlag {
		fmt.Printf("%s version %s %s\n", NAME, VERSION, runtime.GOOS)
		return
	}

	if *helpFlag || len(os.Args) == 1 {
		getopt.Usage()
		return
	}

	switch getopt.Arg(0) {
	case "install", "update":
		commandNotImplementedError(getopt.Arg(0))
	case "create":
		arg := getopt.Arg(1)

		if arg != "" {
			if !*directoryFlag {
				_, err := os.Create(arg)
				check(err)
			} else {
				check(os.Mkdir(arg, 0755))
			}
		}
	case "remove":
		arg := getopt.Arg(1)

		if arg != "" {
			err := os.Remove(arg)
			check(err)
		} else {
			cliErrorln("File name must be included for", getopt.Arg(0), "command")
		}

	// Deprecated commands
	case "newfile":
		filename := getopt.Arg(1)

		if filename != "" {
			_, err := os.Create(filename)
			check(err)
		} else {
			cliErrorln("File name must be included for", getopt.Arg(0), "command")
		}

	default:
		cliErrorf("Unknown command '%s'\n", getopt.Arg(0))
	}

}

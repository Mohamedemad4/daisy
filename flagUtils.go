package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

const (
	MODE_OR  = "or"
	MODE_NOT = "not"
	MODE_AND = "and"
)

func printUsage() {
	magenta := color.New(color.FgMagenta)
	blue := color.New(color.FgBlue)

	magenta.Fprintf(flag.CommandLine.Output(), "daisy: && chain commands across terminal sessions \n\n")
	blue.Fprintf(flag.CommandLine.Output(), "In first terminal: dchn cmd1 apt update \n")
	blue.Fprintf(flag.CommandLine.Output(), "In second terminal: dchn cmd1 apt upgrade \n\n")
	flag.PrintDefaults()
	magenta.Fprintf(flag.CommandLine.Output(), "\nman dchn for more details \n")
	os.Exit(2)
}

// check if the program exists in path
func checkIfInPath(name string) bool {
	for _, path_dir := range strings.Split(os.Getenv("PATH"), ":") {
		progs, err := ioutil.ReadDir(path_dir)

		if err != nil {
			color.Red("unable to read from path at %s", path_dir)
		}

		for _, prog := range progs {
			if prog.Name() == name {
				return true
			}
		}
	}
	return false

}

func handleFlags() {

	CFG_DIR_PATH = filepath.Join(os.Getenv("HOME"), ".dchn")

	flag.Usage = printUsage

	flag.StringVar(&chain_mode, "m", MODE_OR, "Mode: or,xor,and,not (see readme)")
	flag.StringVar(&afterCmd, "after", "nil", "if you have already executed a command without using dchn.\n you can tell dchn to run after that command by specifing it here")

	flag.Parse()

	if len(flag.Args()) < 2 {
		printUsage()
	}

	if !(chain_mode == MODE_OR || chain_mode == MODE_AND || chain_mode == MODE_NOT) {
		printUsage()
	}

	cmdID = flag.Args()[0]
	cmd = flag.Args()[1:]

	if !checkIfInPath(cmd[0]) {
		color.Red("Unable to find program:%s in path, are you sure you specified a cmdID?\n\n", color.GreenString(cmd[0]))
		printUsage()
	}

	logger.Debugf("mode: " + chain_mode)
	logger.Debugf("cmdID: " + cmdID)
	logger.Debugf("cmd: " + strings.Join(cmd, " "))
	logger.Debugf("after cmd: " + afterCmd)
}

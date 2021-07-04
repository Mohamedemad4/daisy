package main

import (
	"flag"
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
}

func handleFlags() {

	CFG_DIR_PATH = filepath.Join(os.Getenv("HOME"), ".dchn")

	flag.Usage = printUsage

	flag.StringVar(&chain_mode, "m", MODE_OR, "Mode: or,xor,and,not (see readme)")
	flag.StringVar(&afterCmd, "after", "nil", "if you have already executed a command without using dchn.\n you can tell dchn to run after that command by specifing it here")

	flag.Parse()

	if len(flag.Args()) < 2 {
		printUsage()
		os.Exit(2)
	}

	if !(chain_mode == MODE_OR || chain_mode == MODE_AND || chain_mode == MODE_NOT) {
		printUsage()
		os.Exit(2)
	}

	cmdID = flag.Args()[0]
	cmd = flag.Args()[1:]

	logger.Debugf("mode: " + chain_mode)
	logger.Debugf("cmdID: " + cmdID)
	logger.Debugf("cmd: " + strings.Join(cmd, " "))
	logger.Debugf("after cmd: " + afterCmd)
}

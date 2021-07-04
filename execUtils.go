package main

import (
	"github.com/fatih/color"
)

func waitForExec(s CommandConfig) {
	// uk the drill
	// while loops and shit
	color.Red("NO MORE MR NICE GUY")
}

func executeCommand() {
	//todo
	color.Red("exec comand")
}

// Check whether or not we should be executed depending on the state of our parent
func evalExitCode(s CommandConfig) bool {
	switch chain_mode {

	case MODE_OR: //execute whatever anyways
		return true

	case MODE_NOT: // only execute if parent failed

		if s.ExitCode > 0 {
			return true
		}
		return false

	case MODE_AND: // only execute if parent

		if s.ExitCode > 0 {
			return false
		}
		return true

	default:
		return true
	}
}

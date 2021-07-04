package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func waitForExec(s CommandConfig) {
	cmdIDFile := filepath.Join(CFG_DIR_PATH, s.CmdID+".json")

	for {
		jsonContents := readCmdIDFile(cmdIDFile)
		if jsonContents.CommandState == DONE {
			break
		} else {
			time.Sleep(time.Second * 1)
		}
	}

}

func executeCommand() int {
	command := exec.Command(cmd[0])
	command.Args = cmd[0:]
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	err := command.Run()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
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

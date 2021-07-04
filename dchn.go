package main

import (
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/hhkbp2/go-logging"
)

var logger logging.Logger
var chain_mode string
var afterCmd string
var cmdID string
var cmd []string
var CFG_DIR_PATH string

func main() {
	// init logger
	logger = logging.GetLogger("root")
	handler := logging.NewStdoutHandler()
	logger.AddHandler(handler)
	logger.SetLevel(logging.LevelDebug)

	handleFlags()

	switch parentState, jsonContents := getCommandParentState(); parentState {

	case DONE:
		// eval ExitCode and mode
		exec_cmd := evalExitCode(jsonContents)
		if exec_cmd {
			writeCmdState(EXECUTING, NO_EXIT_YET)
			exit_code := executeCommand()
			writeCmdState(DONE, exit_code)
		} else {

			color.New(color.FgBlue).Fprintf(color.Output, "parent command commandID: %s\n %s \nExited with code: %s \nNot Executing",
				jsonContents.CmdID,
				color.GreenString(strings.Join(jsonContents.Cmd, " ")),
				color.RedString(strconv.Itoa(jsonContents.ExitCode)),
			)

		}

	case NO_PARENT:
		color.New(color.FgBlue).Fprintf(color.Output, "Executing\n %s\n", color.GreenString(strings.Join(cmd, " ")))

		writeCmdState(EXECUTING, NO_EXIT_YET)
		exit_code := executeCommand()
		writeCmdState(DONE, exit_code)

	case WAITING, EXECUTING:

		color.New(color.FgBlue).Fprintf(color.Output, "waiting for commandID: %s\n %s\nto execute first\n",
			jsonContents.CmdID,
			color.GreenString(strings.Join(jsonContents.Cmd, " ")),
		)

		waitForExec(jsonContents)

		color.Green("Executing....")
		exit_code := executeCommand()

		writeCmdState(DONE, exit_code)

	}
	cleanUP()
}

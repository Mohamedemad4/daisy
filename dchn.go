package main

import (
	"fmt"
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
			writeCmdIDFile(EXECUTING)
			executeCommand()
			writeCmdIDFile(DONE)
		} else {

			color.New(color.FgBlue).Fprintf(color.Output, "parent command commandID: %s\n %s \nExited with code: %s \nNot Executing",
				jsonContents.CmdID,
				color.GreenString(strings.Join(jsonContents.Cmd, " ")),
				color.RedString(strconv.Itoa(jsonContents.ExitCode)),
			)

		}

	case NO_PARENT:
		fmt.Println("no parents")
		color.New(color.FgBlue).Fprintf(color.Output, "Executing\n %s\n", color.GreenString(strings.Join(cmd, " ")))

		writeCmdIDFile(EXECUTING)
		executeCommand()
		writeCmdIDFile(DONE)

	case WAITING, EXECUTING:

		color.New(color.FgBlue).Fprintf(color.Output, "waiting for commandID: %s\n %s\nto execute first\n",
			jsonContents.CmdID,
			color.GreenString(strings.Join(jsonContents.Cmd, " ")),
		)

		waitForExec(jsonContents)
		executeCommand()
		writeCmdIDFile(DONE)

	}
	cleanUP()
}

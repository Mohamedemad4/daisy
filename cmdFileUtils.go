package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
)

type CommandConfig struct {
	CmdID        string
	Cmd          []string
	CmdStart     int64
	CommandState int
	ExitCode     int
}

const (
	WAITING = iota
	EXECUTING
	DONE
	NO_PARENT
	NO_EXIT_YET
)

// checks to see if we are the parent command or if we should wait for some parent to stop executing first
func getCommandParentState() (int, CommandConfig) {

	files := listCmdIDFiles()

	for _, file := range files {
		currCmdFile := filepath.Join(CFG_DIR_PATH, file.Name())
		logger.Debugf(currCmdFile)

		if file.Name() == cmdID+".json" {
			var jsonContents CommandConfig
			dat, err := ioutil.ReadFile(currCmdFile)
			if err != nil {
				logger.Fatal(err)
			}

			json.Unmarshal([]byte(dat), &jsonContents)
			logger.Debugf("%+v", jsonContents)

			return jsonContents.CommandState, jsonContents
		}
	}
	return NO_PARENT, CommandConfig{}
}

// Writes the cmdID.json file
// if it exists do only modify the state
// and or append exit codes
func writeCmdState(CommandState int, ExitCode int) {
	cmdIDFile := filepath.Join(CFG_DIR_PATH, cmdID+".json")
	if fileExists(cmdIDFile) {
		jsonContents := readCmdIDFile(cmdIDFile)
		jsonContents.CommandState = CommandState
		jsonContents.ExitCode = ExitCode
		writeCmdIDFileRaw(cmdIDFile, jsonContents)

	} else {
		jsonContents := CommandConfig{
			Cmd:          cmd,
			CmdID:        cmdID,
			CmdStart:     time.Now().Unix(),
			CommandState: CommandState,
			ExitCode:     ExitCode,
		}
		writeCmdIDFileRaw(cmdIDFile, jsonContents)
	}
}

func listCmdIDFiles() []os.FileInfo {
	files, err := ioutil.ReadDir(CFG_DIR_PATH)
	if err != nil {
		color.Red("unable to open \nthis is probably your first time running this...")

		if _, ok := err.(*os.PathError); ok { //what does this mean?
			err := os.Mkdir(CFG_DIR_PATH, 0777)
			if err != nil {
				logger.Fatal(err)
			}
			color.Blue("created ~/.dchn directory")
		} else {
			logger.Fatal(err)
		}
	}
	return files
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func readCmdIDFile(cmdIDFile string) CommandConfig {
	var jsonContents CommandConfig
	dat, err := ioutil.ReadFile(cmdIDFile)
	if err != nil {
		logger.Fatal(err)
		return CommandConfig{}
	}

	json.Unmarshal([]byte(dat), &jsonContents)
	return jsonContents
}

func writeCmdIDFileRaw(cmdIDFile string, jsonContents CommandConfig) {
	content, _ := json.Marshal(jsonContents)

	err := ioutil.WriteFile(cmdIDFile, content, 0644)

	if err != nil {
		logger.Fatal(err)
	}
}

func cleanUP() {
	color.Red("cleaning up our parents")
}

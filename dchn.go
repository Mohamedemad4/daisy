package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/hhkbp2/go-logging"
)

var logger logging.Logger
var mode string
var afterCmd string
var cmdID string
var cmd []string
var CFG_DIR_PATH string

func readCmdIDFiles() {
	files, err := ioutil.ReadDir(CFG_DIR_PATH)
	if err != nil {
		color.Red("unable to open \nthis is probably your first time running this...")

		if _, ok := err.(*os.PathError); ok { //what does this mean?
			err := os.Mkdir(os.Getenv("HOME")+"/.dchn", 0777)
			if err != nil {
				log.Fatal(err)
			}
			color.Blue("created ~/.dchn directory")
		} else {
			logger.Fatal(err)
		}
	}

	for _, file := range files {
		currCmdFile := filepath.Join(CFG_DIR_PATH, file.Name())
		if file.Name() == cmdID+".json" {
			color.Blue("waiting for cmd to execute first")
			waitForExec()
		}

		logger.Debugf(currCmdFile)

	}

}
func main() {
	// init logger
	logger = logging.GetLogger("root")
	handler := logging.NewStdoutHandler()
	logger.AddHandler(handler)
	logger.SetLevel(logging.LevelDebug)

	handleFlags()
	readCmdIDFiles()
}

func waitForExec() {
	//todo
	color.Red("NO MORE MR NICE GUY")
}

//
//func executeCommand() {
//	//todo
//	color.Red("exec")
//}
//
//func cleanUP() {
//	//todo also clean orphans with your parent
//}

package main

import (
	"os/exec"
	"bytes"
)

var source = make(chan []byte)

func init() {
	go sourceHandler()
}

func sourceHandler() {
	for {
		logger.Debug("Start take snapshot")
		command := exec.Command("avconv",
			"-f", "x11grab",
			"-s", "1920:1080",
			"-i", ":0.0",
			"-s", "1280:720",
			"-f", "image2",
			"-frames", "1",
			"-",
		)

		errorBuffer := new(bytes.Buffer)

		command.Stderr = errorBuffer

		data, cmdErr := command.Output()
		logger.Debug("End take snapshot")

		if nil != cmdErr {
			logger.Error(cmdErr.Error())
			logger.Error("Error buffer: %s", errorBuffer)

			continue
		}

		logger.Debug("Start push source")
		source <- data
		logger.Debug("End push source")
	}
}

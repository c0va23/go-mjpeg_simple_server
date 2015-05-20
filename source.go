package main

import (
	"os/exec"
	"bytes"
	"log"
)

var source = make(chan []byte)

func init() {
	go sourceHandler()
}

func sourceHandler() {
	for {
		log.Printf("Start take snapshot")
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
		log.Printf("End take snapshot")

		if nil != cmdErr {
			log.Printf(cmdErr.Error())
			log.Printf("Error buffer: %s", errorBuffer)

			continue
		}

		log.Printf("Start push source")
		source <- data
		log.Printf("End push source")
	}
}

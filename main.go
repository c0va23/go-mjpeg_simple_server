package main

import (
	"net/http"
	"log"
)

const listen = "localhost:40000"

func main() {
	serveMux := http.NewServeMux()

	fileSystem := http.Dir("public")
	serveMux.Handle("/", http.FileServer(fileSystem))
	serveMux.HandleFunc("/jpeg", jpeg)
	serveMux.HandleFunc("/mjpeg", mjpeg)

	log.Printf("Start listen on %s\n", listen)
	if err := http.ListenAndServe(listen, serveMux); nil != err {
		log.Printf("Error: %s", err.Error())
	}
}

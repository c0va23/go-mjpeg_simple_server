package main

import (
	"net/http"
	logging "github.com/op/go-logging"
)

var logger = logging.MustGetLogger("SERVER")

const listen = "localhost:40000"

func main() {
	serveMux := http.NewServeMux()

	fileSystem := http.Dir("public")
	serveMux.Handle("/", http.FileServer(fileSystem))
	serveMux.HandleFunc("/jpeg", jpeg)
	serveMux.HandleFunc("/mjpeg", mjpeg)

	logger.Info("Start listen on %s", listen)
	if err := http.ListenAndServe(listen, serveMux); nil != err {
		logger.Error(err.Error())
	}
}

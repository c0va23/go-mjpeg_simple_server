package main

import (
	"net/http"
)

func jpeg(responseWriter http.ResponseWriter, request *http.Request) {
	logger.Info("Start request %s", request.URL)

	logger.Info("Wait source")
	snapshot := <- source
	logger.Info("Write snapshot")

	responseWriter.Header().Add("Content-Type", "image/jpeg")
	responseWriter.Write(snapshot)
	logger.Info("Success request")
}

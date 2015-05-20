package main

import (
	"fmt"
	"net/http"
	"mime/multipart"
	"net/textproto"
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

func mjpeg(responseWriter http.ResponseWriter, request *http.Request) {
	logger.Info("Start request %s", request.URL)

	mimeWriter := multipart.NewWriter(responseWriter)

	logger.Debug("Boundary: %s", mimeWriter.Boundary())

	contentType := fmt.Sprintf("multipart/x-mixed-replace;boundary=%s", mimeWriter.Boundary())
	responseWriter.Header().Add("Content-Type", contentType)

	for {
		partHeader := make(textproto.MIMEHeader)
		partHeader.Add("Content-Type", "image/jpge")

		partWriter, partErr := mimeWriter.CreatePart(partHeader)
		if nil != partErr {
			logger.Error(partErr.Error())
			break
		}

		snapshot := <- source
		if _, writeErr := partWriter.Write(snapshot); nil != writeErr {
			logger.Error(writeErr.Error())
		}
	}

	logger.Info("Success request")
}

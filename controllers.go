package main

import (
	"fmt"
	"time"
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
		frameStartTime := time.Now()
		partHeader := make(textproto.MIMEHeader)
		partHeader.Add("Content-Type", "image/jpeg")

		partWriter, partErr := mimeWriter.CreatePart(partHeader)
		if nil != partErr {
			logger.Error(partErr.Error())
			break
		}

		snapshot := <- source
		if _, writeErr := partWriter.Write(snapshot); nil != writeErr {
			logger.Error(writeErr.Error())
		}
		frameEndTime := time.Now()

		frameDuration := frameEndTime.Sub(frameStartTime)
		fps := float64(time.Second) / float64(frameDuration)
		logger.Info("Frame time: %s (%.2f)", frameDuration, fps)
	}

	logger.Info("Success request")
}

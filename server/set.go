package server

import (
	"ClickHitsCount/infrastructure/logs"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setServerLog() (*os.File, error) {
	logFile, err := os.Create("./log/restful_server.log")
	if err != nil {
		logs.Warning.Println(err)
		return nil, err
	} else {
		gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
		return logFile, nil
	}
}

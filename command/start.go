package command

import (
	"ClickHitsCount/infrastructure/configs"
	"ClickHitsCount/infrastructure/database"
	"ClickHitsCount/infrastructure/logs"
	"ClickHitsCount/infrastructure/trace"
	"ClickHitsCount/server"
	"log"
	"time"
)

func Start() {
	// -----------------------------------------------[Channel]
	logChannel := make(chan int, 1)
	traceChannel := make(chan int, 1)
	databaseChannel := make(chan int, 1)
	// -----------------------------------------------[Init]
	configs.InitializationViper()
	go logs.InitializationLog(logChannel)
	<-logChannel
	go trace.InitializationTrace(traceChannel)
	<-traceChannel
	database.InitializationDatabase(databaseChannel)
	<-databaseChannel
	// -----------------------------------------------[service]
	server.Start()
	// -----------------------------------------------[stop]
	databaseChannel <- 1
	traceChannel <- 1
	logChannel <- 1
	log.Println("Closing System,please wait 10 second.")
	time.Sleep(10 * time.Second)
	log.Println("Closed System.")
	panic(nil)
}

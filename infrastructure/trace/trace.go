package trace

import (
	"ClickHitsCount/infrastructure/logs"
	"os"
	"runtime/trace"
)

func InitializationTrace(stop chan int) {
	if file, err := os.Create(".//log//trace.out"); err != nil {
		logs.Error.Panic(err)
	} else if err = trace.Start(file); err != nil {
		logs.Error.Panic(err)
	} else {
		defer file.Close()
	}
	stop <- 1
	<-stop
}

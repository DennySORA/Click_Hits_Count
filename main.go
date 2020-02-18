package main

import (
	"ClickHitsCount/command"
	"log"
	"runtime/debug"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(string(debug.Stack()))
		}
	}()
	command.Start()
}

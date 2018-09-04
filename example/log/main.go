package main

import (
	"log"
	"os"

	butil "github.com/bonfy/butil-golang"
)

func main() {
	var logPath = "log.log"
	logFile, _ := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer logFile.Close()
	butil.InitLogFile(logFile)

	log.Println("Hello Test Log")
}

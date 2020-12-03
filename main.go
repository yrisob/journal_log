package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/yrisob/journal_log/webserver"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Старт пакета ppas_logger")
	if err := webserver.Start(); err != nil {
		panic(err)
	}
}

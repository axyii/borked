package utils

import (
	"fmt"
	"log"
	"os"
)

//Logger is exported
var Logger *log.Logger

func init() {

	Logfile, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	Logger = log.New(Logfile, "Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}

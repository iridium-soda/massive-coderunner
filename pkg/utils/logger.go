package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*Provide logger interface*/

func InitLogger() {
	//New a logger object
	//logger := log.New(os.Stdout, "", log.Lshortfile|log.Ldate|log.Ltime)

	os.MkdirAll("log", os.ModePerm)                     //Create a dir before
	log.SetFlags(log.Lshortfile | log.Ldate | log.LUTC) //Set Config
	timeStamp := time.Now()
	timeString := timeStamp.Format("2006-01-02_15-04-05") //Convert to str
	logFileName := "./log/" + timeString + ".log"
	logFile, err := os.Create(logFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(logFile) //Set Output File
}

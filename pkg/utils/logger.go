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
	log.SetFlags(log.Lshortfile | log.Ldate | log.LUTC) //Set Config
	timeStamp := time.Now()
	timeString := timeStamp.Format("2006-01-02_15:04:05") //Convert to str
	logFile, err := os.OpenFile("./log/[RunLog]"+timeString+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile) //Set Output File
}

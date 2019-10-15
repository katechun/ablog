package helpers

import (
	"log"
	"os"
	"time"
)
var Logger *log.Logger

func InitLog() {
	Logger = WriteLog()
}

func WriteLog() *log.Logger{

	logpath:=*Logfile+"ablog_"+time.Now().Format("20060102")+".log"

    file,err := os.OpenFile(logpath,os.O_CREATE|os.O_RDWR|os.O_APPEND,0666)

    if err != nil {
    	log.Fatal("Log write Error!")
	}
	logger := log.New(file,"",log.Llongfile)

	return logger
}

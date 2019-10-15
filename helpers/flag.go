package helpers

import (
	"flag"
)

var (
	Configfile *string
	Logfile *string
)

func InitConfig() {
	//命令行参数
	Configfile = flag.String("c","conf/conf.ini","config file name")
	Logfile = flag.String("l","logs/","log file path")
	flag.Parse()
}

package helpers

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)


func GetConf() *goconfig.ConfigFile{
	cfg,err := goconfig.LoadConfigFile(*Configfile)

	if err != nil {
		fmt.Println("解析配置文件错误")
	}

	return cfg

}


func GetOneValue(val1,val2 string) string{

	//读取配置文件信息
	cfg := GetConf()
	value,err :=cfg.GetValue(val1,val2)
	if err != nil {
		fmt.Println("读取配置文件值错误：",err)
	}
	return value

}

func GetListValue(val string) map[string]string{

	//读取配置文件信息
	cfg := GetConf()
	value,err :=cfg.GetSection(val)
	if err != nil {
		fmt.Println("读取配置文件值错误：",err)
	}
	return value
}
package config

import (
	"github.com/Unknwon/goconfig"
	"log"
)

var Config *goconfig.ConfigFile

type MysqlSetting struct {
	UserName string
	Password string
	IP       string
	DataName string
}

func InitConfig() { //读取配置文件
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Println("读取配置文件失败[config.ini]", err)
		return
	}
	Config = cfg
}

func GetPort() string { //获取端口
	value, err := Config.GetValue("dev", "port")

	if err != nil {
		log.Fatalf("无法读取键-值(%s):%s", value, err)
	}
	return value
}

func GetSqlSetting() *MysqlSetting {
	sec, _ := Config.GetSection("mysql")
	myconfig := &MysqlSetting{} //定义一个结构体
	myconfig.UserName = sec["name"]
	myconfig.Password = sec["password"]
	myconfig.IP = sec["ip"]
	myconfig.DataName = sec["database"]
	return myconfig
}

func GetFileServer() string {
	value, err := Config.GetValue("file", "FileServer")
	if err != nil {
		log.Fatalf("无法读取键-值(%s):%s", value, err)
	}
	return value
}

func GetLocalPath() string {
	value, err := Config.GetValue("file", "LocalPath")
	if err != nil {
		log.Fatalf("无法读取键-值(%s):%s", value, err)
	}
	return value
}

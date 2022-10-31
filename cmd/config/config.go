package config

import (
	"github.com/yoshimi-I/Go_RESTAPI/utils"
	"gopkg.in/go-ini/ini.v1"
	"log"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

//configから値を受け取ってConfigに格納する
//あらかじめgo get "gopkg.in/go-ini/ini.v1"でインストール
var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}
func LoadConfig() {
	cfg, err := ini.Load("cmd/config/config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		cfg.Section("web").Key("port").MustString("8080"),
		cfg.Section("db").Key("driver").String(),
		cfg.Section("db").Key("name").String(),
		cfg.Section("web").Key("logfile").String(),
	}
}

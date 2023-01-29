package configs

import (
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = new(AppConfig)

type AppConfig struct {
	*Server
	*Pgsql
}
type Pgsql struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

type Server struct {
	Name string
	Port string
}



func ReadConfig(config string) error {
	viper.SetConfigFile(config)
	err := viper.ReadInConfig()
	if err != nil {
		hlog.Errorf("读取配置文件异常：%s", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		hlog.Warnf(in.String())
		hlog.Warnf("配置文件被修改了~~~~")
	})

	//解析配置文件
	if err := viper.Unmarshal(&Config); err != nil {
		hlog.Errorf("配置解析错误~~~~：%s", err)
		return err
	}
	marshal, _ := json.Marshal(Config)
	println(string(marshal))
	return nil
}

package main

import (
	"github.com/phuhao00/gddd/chat/internal/adapter"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote" // 必须导入，才能加载远程 key/value 配置
)

func main() {
	a := adapter.NewAdapter()
	err := viper.AddRemoteProvider("consul", "localhost:8500", "user/config") // 连接远程 consul 服务
	if err != nil {
		panic(err)
	}
	viper.SetConfigType("YAML") // 显式设置文件格式文 YAML
	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
	a.Run()
}

package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// Init 加载配置
func Init() (err error) {
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	err = viper.ReadInConfig()           // 读取配置信息
	if err != nil {                      // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})

	return
}

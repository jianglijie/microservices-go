package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
	"utils"
)

var (
	cfg  *tomlConfig
	once sync.Once
)

type tomlConfig struct {
	Title         string
	ConfigVersion int `toml:"config_version"`
	Mysql         mysqlConfig
	Redis         redisConfig
}

type mysqlConfig struct {
	Server          string
	Port            string
	User            string
	Password        string
	DefaultDatabase string `toml:"default_database"`
}

type redisConfig struct {
	Server   string
	Port     string
	Password string
}

func Config() *tomlConfig {
	once.Do(func() {
		viper.SetConfigName("config")   // 设置配置文件名 (不带后缀)
		viper.AddConfigPath(".")        // 第一个搜索路径
		viper.AddConfigPath("/Users/jianglj/work/go_ms")        // 第一个搜索路径
		err := viper.ReadInConfig()     // 读取配置数据
		if err != nil {
			content := fmt.Sprintf("init config error: %s", err)
			fields := make(map[string]interface{})
			fields["type"] = "config"
			fields["ope"] = "init"
			utils.LogError(content, fields)
		}
		viper.Unmarshal(&cfg)        // 将配置信息绑定到结构体上
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println(e.Op)
			fmt.Println(e.String())
			fmt.Println("Config file changed:", e.Name)
			err := viper.ReadInConfig()     // 读取配置数据
			if err != nil {
				content := fmt.Sprintf("update config error: %s", err)
				fields := make(map[string]interface{})
				fields["type"] = "config"
				fields["ope"] = "update"
				utils.LogError(content, fields)
			}
			viper.Unmarshal(&cfg)        // 将配置信息绑定到结构体上
		})
		fmt.Println(2333)
	})
	fmt.Println(222)
	return cfg
}

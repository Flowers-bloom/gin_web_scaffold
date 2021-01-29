package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func InitLocalConf() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := dir + "\\config\\config.json"
	viper.SetConfigFile(path)

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 热加载
	Watcher()
}

func Watcher() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("config file changed: %s\n", in.Name)
	})
}

func InArrayList(module string, moduleList []string) bool {
	for _, s := range moduleList {
		if s == module {
			return true
		}
	}
	return false
}
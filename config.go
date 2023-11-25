package main

import (
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Config struct {
	DownloadPath    string `json:"download_path"`
	CachePath       string `json:"cache_path"`
	VideoListPath   string `json:"videolist_path"`
	DownloadThreads int    `json:"download_threads"`
}

// 获取设置内容
func GetConfig() Config {
	for {
		// 判断设置文件是否已经存在
		if !IsFileExists("./config.json") {
			// 文件不存在
			cfg := bulidConfig()
			err := SaveJsonFile("./config.json", &cfg)
			if err != nil {
				// runtime.LogError(a.ctx, "写入设置文件失败："+err.Error())
				fmt.Printf("写入设置文件失败：%s", err)
			}
		} else {
			// 文件已存在
			var cfg Config
			err := LoadJsonFile("./config.json", &cfg)
			if err != nil {
				// runtime.LogError(a.ctx, "读取设置文件失败："+err.Error())
				fmt.Printf("读取设置文件失败：%s", err)
			}
			return cfg
		}
	}
}

// 写入设置
func (a *App) SaveConfig(cfg Config) {
	err := SaveJsonFile("./config.json", cfg)
	if err != nil {
		runtime.LogError(a.ctx, "写入设置文件失败："+err.Error())
	}
}

// 创建默认设置结构体
func bulidConfig() *Config {
	return &Config{
		DownloadPath:    "./Download",
		CachePath:       "./Cache",
		VideoListPath:   "./Cache/video_list.json",
		DownloadThreads: 5,
	}
}

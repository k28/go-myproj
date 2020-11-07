package myproj

import (
	"os"
	"path/filepath"
	"runtime"
)

// ConfigDirPath アプリの設定情報を保存するパスを返します
func ConfigDirPath(appName string) string {
	var configDir string
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		configDir = os.Getenv("APPDATA")
	} else {
		configDir = filepath.Join(home, ".config")
	}

	path := filepath.Join(configDir, appName)
	return path
}

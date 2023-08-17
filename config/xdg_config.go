package config

import (
	"fmt"
	"os"
)

func findXDGConfigDir() string {
	if isXDGConfigDirSet() {
		return os.Getenv("XDG_CONFIG_HOME")
	} else {
	return ""
	}
}

func isXDGConfigDirSet() bool {
	return os.Getenv("XDG_CONFIG_HOME") != ""
}

func createXDGConfigDir() {
	if err := os.MkdirAll(os.Getenv("XDG_CONFIG_HOME"), 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

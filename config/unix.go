package config

import "os"

func isUnixHomeSet() bool {
	return os.Getenv("HOME") != ""
}

func findUnixConfigDir() string {
	if isUnixHomeSet() {
		return os.Getenv("HOME") + "/.config"
	}
	return ""
}

func createUnixConfigDir() {
	if err := os.MkdirAll(os.Getenv("HOME") + "/.config", 0755); err != nil {
		os.Exit(1)
	}
}
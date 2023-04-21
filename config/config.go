package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Version	int
	Name				string
	BinDir 			string
	ProjectDir 	string
}

func LoadConfig() Config {
		// Check if config file exists, if not create it
	checkConfigFile()

	config_file := os.Getenv("HOME") + "/.config/gops.toml"
	checkFile(config_file)
	file, _ := os.ReadFile(config_file)
	var cfg Config

	err := toml.Unmarshal([]byte(file), &cfg)
	if err != nil {
		fmt.Println("Error while parsing config file", err)
		panic(err)
	}
	return cfg
}

func checkFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func checkConfigDir() {
	if err := os.MkdirAll(os.Getenv("HOME") + "/.config", 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkConfigFile() {
	checkConfigDir()

	file := os.Getenv("HOME") + "/.config/gops.toml"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("Config file does not exist")
		file, err := os.Create(file)
		if err != nil {
				log.Fatal(err)
		}
		defer file.Close()
	}	else {
		currentTime := time.Now().Local()
		err = os.Chtimes(file, currentTime, currentTime)
		if err != nil {
				fmt.Println(err)
		}
	}
}
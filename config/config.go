package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Version    int
	Name       string
	BinDir     string
	ProjectDir string
}

func LoadConfig() Config {
	var cfg Config
	config_file := findConfigFile()
	checkConfigFile(config_file)

	file, _ := os.ReadFile(config_file)

	err := toml.Unmarshal([]byte(file), &cfg)
	if err != nil {
		fmt.Println("Error while parsing config file", err)
		panic(err)
	}
	return cfg
}

func findConfigFile() string {
	var file string
	if isXDGConfigDirSet() {
		createXDGConfigDir()
		file = findXDGConfigDir() + "/gops.toml"
	} else if isUnixHomeSet() {
		createUnixConfigDir()
		file = findUnixConfigDir() + "/gops.toml"
	} else {
		fmt.Println("Error while finding config directory")
		os.Exit(1)
	}
	return file
}

func checkConfigFile(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("Config file does not exist")
		file, err := os.Create(file)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	} else {
		currentTime := time.Now().Local()
		err = os.Chtimes(file, currentTime, currentTime)
		if err != nil {
			fmt.Println(err)
		}
	}
}

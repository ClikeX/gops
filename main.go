package main

import (
	"os"

	"github.com/ClikeX/gops/config"
	"github.com/ClikeX/gops/utils"
)

func main() {
	cfg := config.LoadConfig()

	// Check if correct arguments are provided
	utils.CheckArgs(os.Args)

	// Parse arguments
	git_url, name, hostname := utils.ParseArgs(os.Args)

	// Create directory structure
	utils.CreateDirStructure(cfg.ProjectDir + "/" + hostname)

	// Clone repo
	utils.CloneRepo(git_url, cfg.ProjectDir + "/" + name)
}


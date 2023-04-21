package utils

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	giturls "github.com/whilp/git-urls"
)

func CheckArgs(args []string) {
	if len(args) == 1 {
		fmt.Println("No arguments provided, gops requires one argument")
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Println("Too many arguments provided, gops only accepts one argument")
		os.Exit(1)
	}
}

func ParseArgs(args []string) (*url.URL, string, string) {
	git_url, _ := giturls.Parse(args[1])
	repo := strings.Split(git_url.Path, ".")[0]
	hostname := git_url.Host

	return git_url, repo, hostname
}
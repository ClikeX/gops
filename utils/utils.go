package utils

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
)

func CreateDirStructure(path string) {
	// Create directory
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CloneRepo(repo *url.URL, path string) {
	fmt.Printf("Cloning %s into %s \n", repo, path)

	out, err := exec.Command("git", "clone", repo.String(), path).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}

func parseURL(uri string) (*url.URL, error) {
	parsedUrl, err := url.Parse(uri)
	return parsedUrl, err
}
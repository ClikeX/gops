package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	git "github.com/go-git/go-git/v5"
	giturls "github.com/whilp/git-urls"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No arguments provided, gops requires one argument")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments provided, gops only accepts one argument")
		os.Exit(1)
	}

	uri, _ := giturls.Parse(os.Args[1])
	repo := strings.Split(uri.Path, ".")[0]
	path := os.Getenv("HOME") + "/Code/" + uri.Host

	fmt.Printf("Cloning %s into %s \n", repo, path)

	// Create directory
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, gerr := git.PlainClone(path + "/" + repo, false, &git.CloneOptions{
			URL:      os.Args[1],
			Progress: os.Stdout,
	})

	if gerr != nil {
		fmt.Println(gerr)
		os.Exit(1)
	}
}

func ParseURL(uri string) (*url.URL, error) {
	parsedUrl, err := url.Parse(uri)
	return parsedUrl, err
}

package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	git "github.com/go-git/go-git/v5"
	cli "github.com/urfave/cli/v2"
	giturls "github.com/whilp/git-urls"
)

func main() {
	app := &cli.App{
		Name:  "gops",
		Usage: "A simple tool to manage your git repositories",
		Commands: []*cli.Command{
			{
				Name:    "clone",
				Aliases: []string{"c"},
				Usage:   "Clone a repository",
				Action: func(c *cli.Context) error {
					uri, _ := giturls.Parse(c.Args().First())
					repo := strings.Split(uri.Path, ".")[0]
					path := os.Getenv("HOME") + "/Code/" + uri.Host

					fmt.Printf("Cloning %s into %s \n", repo, path)

					err := os.MkdirAll(path, 0755)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					_, gerr := git.PlainClone(path+"/"+repo, false, &git.CloneOptions{
						URL:      c.Args().First(),
						Progress: os.Stdout,
					})

					if gerr != nil {
						fmt.Println(gerr)
						os.Exit(1)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func ParseURL(uri string) (*url.URL, error) {
	parsedUrl, err := url.Parse(uri)
	return parsedUrl, err
}

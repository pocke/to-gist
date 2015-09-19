package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"github.com/mitchellh/go-homedir"
	"github.com/pocke/gha"
)

func main() {
	err := CLIMain(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func CLIMain(args []string) error {
	c, err := FlagParse(args)
	if err != nil {
		return err
	}

	token, err := getToken()
	if err != nil {
		return err
	}

	return CreateGist(c, token)
}

func CreateGist(cli *CLI, token gha.RoundTripper) error {
	fname := github.GistFilename(cli.FileName)
	files := make(map[github.GistFilename]github.GistFile)
	files[fname] = github.GistFile{Content: &cli.FileContent}

	c := github.NewClient(&http.Client{
		Transport: token,
	})

	gist, _, err := c.Gists.Create(&github.Gist{
		Files:  files,
		Public: &cli.Public,
	})
	if err != nil {
		return err
	}

	fmt.Println(*gist.HTMLURL)
	return nil
}

func getToken() (gha.RoundTripper, error) {
	keyFile, err := homedir.Expand("~/.config/to-gist.githubkey")
	if err != nil {
		return "", err
	}
	token, err := gha.CLI(keyFile, &gha.Request{
		Note:   "to-gist",
		Scopes: []string{"gist"},
	})
	if err != nil {
		return "", err
	}
	return gha.RoundTripper(token), nil
}

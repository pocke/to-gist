package main

import (
	"fmt"
	"io/ioutil"
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
	if len(args) == 1 {
		return fmt.Errorf("file name is required")
	}

	token, err := getToken()
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	return CreateGist(args[1], string(content), token)
}

func CreateGist(name, content string, token gha.RoundTripper) error {
	fname := github.GistFilename(name)
	files := make(map[github.GistFilename]github.GistFile)
	files[fname] = github.GistFile{Content: &content}

	c := github.NewClient(&http.Client{
		Transport: token,
	})

	gist, _, err := c.Gists.Create(&github.Gist{
		Files: files,
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

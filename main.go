package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
)

func main() {
	err := CLIMain(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func CLIMain(args []string) error {
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	if len(args) == 1 {
		return fmt.Errorf("file name is required")
	}

	return CreateGist(args[1], string(content))
}

func CreateGist(name, content string) error {
	fname := github.GistFilename(name)
	files := make(map[github.GistFilename]github.GistFile)
	files[fname] = github.GistFile{Content: &content}

	c := github.NewClient(nil)
	gist, _, err := c.Gists.Create(&github.Gist{
		Files: files,
	})
	if err != nil {
		return err
	}

	fmt.Println(*gist.HTMLURL)
	return nil
}

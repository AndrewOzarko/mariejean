package main

import (
	"github.com/andrewozarko/mariejean/cmd/cli"
	"github.com/andrewozarko/mariejean/pkg/cmdRunner"
	"log"
	"os"
	"os/exec"
)

func init() {
	log.SetFlags(0)

	_, err := exec.LookPath("git")
	if err != nil {
		log.Fatalf("Before using this app you need to install: git")
	}

	_, err = exec.LookPath("composer")
	if err != nil {
		log.Fatalf("Before using this app you need to install: composer")
	}
}

func main() {
	mj := &cmdRunner.MarieJean{
		Commands: map[string]cmdRunner.AnyCommand{
			"version": &cli.VersionCommand{},
			"create":  &cli.CreateCommand{},
			"update":  &cli.UpdateCommand{},
		},
	}
	mj.RunCli(os.Args)
}

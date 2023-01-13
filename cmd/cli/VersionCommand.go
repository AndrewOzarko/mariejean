package cli

import (
	"github.com/andrewozarko/mariejean/pkg/cmdRunner"
	"log"
)

type VersionCommand struct {
}

var (
	Version string
)

func (cc *VersionCommand) Run(args cmdRunner.PreparedArgs) {
	log.Printf("mariejane: %v", Version)
}

package cli

import (
	"github.com/andrewozarko/mariejean/pkg/cmdRunner"
	"log"
)

type UpdateCommand struct {
}

func (cc *UpdateCommand) Run(args cmdRunner.PreparedArgs) {
	log.Println("in progress")
}

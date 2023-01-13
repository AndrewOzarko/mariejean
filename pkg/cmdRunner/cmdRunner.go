package cmdRunner

import (
	"log"
	"strings"
)

const OPTION_PREFIX = "--"
const OPTION_PREFIX_SIZE = 2

type PreparedArgs struct {
	Args    []string
	Options map[string]string
}

type AnyCommand interface {
	any
	Run(args PreparedArgs)
}

type MarieJean struct {
	Commands map[string]AnyCommand
}

func prepareArgs(args []string) PreparedArgs {
	newArgs := &PreparedArgs{
		Options: map[string]string{},
		Args:    []string{},
	}
	for i, a := range args {
		if len(a) > OPTION_PREFIX_SIZE && strings.HasPrefix(a, OPTION_PREFIX) {
			keyName := string([]rune(a[OPTION_PREFIX_SIZE:]))
			if len(args) > i+1 {
				newArgs.Options[keyName] = args[i+1]
			} else {
				newArgs.Options[keyName] = ""
			}
		}
		prevOne := ""
		if i != 0 {
			prevOne = args[i-1]
		}
		if !strings.HasPrefix(prevOne, OPTION_PREFIX) && !strings.HasPrefix(a, OPTION_PREFIX) {
			newArgs.Args = append(newArgs.Args, a)
		}
	}
	return *newArgs
}

func (mj *MarieJean) RunCli(args []string) {
	currentArg := ""
	if len(args) > 1 {
		currentArg = args[1]
	}
	if _, ok := mj.Commands[args[1]]; !ok {
		log.Fatalf("%v is undefined command.", args[1])
	}
	for name, value := range mj.Commands {
		if currentArg == name {
			value.Run(prepareArgs(args[1:]))
			break
		}
	}
}

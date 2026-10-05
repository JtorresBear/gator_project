package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JtorresBear/gator_project/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	st8 := state{
		cfg: &cfg,
	}

	cmds := commands{
		cmdMap: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	fmt.Println(cmds.cmdMap)

	userInputs := os.Args
	if len(userInputs) < 2 {
		log.Fatal("too few user inputs")
	}
	handlerName := userInputs[0]
	handlerArgs := userInputs[1:]
	cmdHandle := command{
		name:      handlerName,
		arguments: handlerArgs,
	}

	err = cmds.run(&st8, cmdHandle)
	if err != nil {
		log.Fatal(err)
	}
}

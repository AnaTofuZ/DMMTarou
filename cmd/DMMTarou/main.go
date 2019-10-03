package main

import (
	"flag"
	"log"
	"os"

	"github.com/AnaTofuZ/DMMTarou/cmd/DMMTarou/cmd"
)

func main() {
	log.SetFlags(0)
	err := cmd.NewDMMTarouCmd().Execute()
	if err != nil && err != flag.ErrHelp {
		log.Println(err)
		exitCode := 1
		if ecoder, ok := err.(interface{ ExitCode() int }); ok {
			exitCode = ecoder.ExitCode()
		}
		os.Exit(exitCode)
	}
}

package cmd

import (
	color "cs50-romain/gokey/pkg/colors"
	"fmt"
	"log"
	"os"
)

type Command struct {
	Main	string
	Name	string
}

func Execute(args []string) {
	if args[0] == "exit" {
		os.Exit(0)
	}

	command := buildCommand(args)
	
	if command.Main != "new" && command.Main != "copy" && command.Main != "show" {
		log.Fatal("[ERROR] Invalid command;")
	} else {
		if command.Main == "new" {
			color.PrintBold("Creating new creds for " + command.Name, color.ElectricPink)
		} else if command.Main == "show" {
			color.PrintBold("Fetching creds for " + command.Name, color.ElectricPink)
		} else {
			color.PrintBold("Copying creds for " + command.Name, color.ElectricPink)
		}
		fmt.Println()
	}
}

func buildCommand(args []string) *Command {
	var command Command

	if args[0] != "go" {
		log.Fatal("[ERROR] Invalid command; go needs to be first argument")
	}

	if len(args) < 3 {
		log.Fatal("[ERROR] Invalid command; Not enought arguments")
	} else {
		command.Main = args[1]
		command.Name = args[2]
	}
	return &command
}

package main

import (
	"bufio"
	color "cs50-romain/gokey/pkg/colors"
	"log"
	"os"
	"strings"

	"cs50-romain/gokey/cmd"
)

const asciilogo = `
   ___                      
  / _ \___   /\ /\___ _   _ 
 / /_\/ _ \ / //_/ _ \ | | |
/ /_\\ (_) / __ \  __/ |_| |
\____/\___/\/  \/\___|\__, |
                      |___/ 
`

const possible_cmds = `
GoKey is a simple password manager CLI tool. Manage your passwords directly from the terminal. TBH security is a lil lax.
Use: go <command> <args>
Examples:
	go new <name> OR go new
	go copy <name>
	go show <name>
`

func main() {
	color.PrintBold(string(asciilogo), color.ElectricPink)
	color.PrintBold(string(possible_cmds), color.LimeGreen)

	reader := bufio.NewReader(os.Stdin)
	for {
		color.PrintBold(">", color.RoyalPurple)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("[ERROR] Error reading user input; main.go -> ", err)
		}
		input = strings.TrimSuffix(input, "\n")
		inputarr := strings.Split(input, " ")
		cmd.Execute(inputarr)
	}
}

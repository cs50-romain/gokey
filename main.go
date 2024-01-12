package main

import (
	"bufio"
	color "cs50-romain/gokey/pkg/colors"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"cs50-romain/gokey/cmd"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

type Config struct {
	DB_File			string `json:"db_file"`
	Hashed_Master_Password	string `json:"hashed_password"`
	Hash_Used		string `json:"hash_used"`
}

const master_password = "BuffB@by101"

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
	// TESTING ONLY; HASHING master_password
	// THIS IS ONLY FOR SETUP SO IF config.json DOES NOT EXIST
	master_pass, err := bcrypt.GenerateFromPassword([]byte(master_password), 14)
	config := Config{
		DB_File: "gopass.db",
		Hashed_Master_Password: "",
		Hash_Used: "bcrypt",
	}

	b, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		log.Fatal("[ERROR] Error Marshaling golang type; main.go -> ", err)
	}
	os.WriteFile("config.json", b, 0600)

	color.PrintBold(string(asciilogo), color.ElectricPink)
	color.PrintBold(string(possible_cmds), color.LimeGreen)

	color.PrintBold("Please enter master password: ", color.TealGreen)
	input, err := term.ReadPassword(int(syscall.Stdin)) 
	if err != nil {
		log.Fatal("[ERROR] Error reading user input; main.go -> ", err)
	}

	err = bcrypt.CompareHashAndPassword(master_pass, input)
	if err != nil {
		fmt.Println("Invalid master password")
		os.Exit(0)
	}
	fmt.Println()
	
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

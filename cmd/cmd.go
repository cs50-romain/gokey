package cmd

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	crypto "cs50-romain/gokey/internal/crypto"
	color "cs50-romain/gokey/pkg/colors"

	"golang.org/x/term"
)

type Creds struct {
	Entries []struct {
		Name              string    `json:"name"`
		Username          string    `json:"username"`
		EncryptedPassword string    `json:"encrypted_password"`
	} `json:"entries"`
}

type Command struct {
	Main	string
	Name	string
}

func Execute(args []string, key []byte) {
	if args[0] == "exit" {
		os.Exit(0)
	}

	command := buildCommand(args)
	
	if command.Main != "new" && command.Main != "copy" && command.Main != "show" {
		log.Fatal("[ERROR] Invalid command;")
	} else {
		if command.Main == "new" {
			storeCreds(args[2], key)
			color.PrintBold("Creating new creds for " + command.Name, color.ElectricPink)
			fmt.Println()
		} else if command.Main == "show" {
			color.PrintBold("Fetching creds for " + command.Name, color.ElectricPink)
			fmt.Println()
			showCreds(args[2], key)
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

func showCreds(name string, key []byte) {
	var creds Creds
	b, err := os.ReadFile("gokey.json")
	if err != nil {
		log.Fatal("[ERROR] Error reading gokey.json; cmd.go -> ", err)
	}

	err = json.Unmarshal(b, &creds)
	if err != nil {
		log.Fatal("[ERROR] Error unmarshaling gokey.json; cmd.go -> ", err)
	}

	var password string
	var username string
	for _, entry := range creds.Entries {
		if entry.Name == name {
			username = entry.Username
			password = entry.EncryptedPassword
		}
	}
	base64_password, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		log.Fatal("[ERROR] Error decoding base64 password; cmd.go -> ", err)
	}

	b, err = crypto.Decrypt(base64_password, key)
	if err != nil {
		log.Fatal("[ERROR] Error decrypting; cmd.go -> ", err)
	}

	fmt.Printf("Username: %s\nPassword: %s\n", username, string(b))
}

func storeCreds(name string, key []byte) {
	var creds Creds
	b, err := os.ReadFile("gokey.json")
	if err != nil {
		log.Fatal("[ERROR] Error reading gokey.json; cmd.go -> ", err)
	}
	err = json.Unmarshal(b, &creds)
	if err != nil {
		log.Fatal("[ERROR] Error unmarshaling config.json; cmd.go -> ", err)
	}

	reader := bufio.NewReader(os.Stdin)
	color.PrintBold("Enter username:", color.LimeGreen)
	fmt.Println()
	color.PrintBold(">", color.RoyalPurple)
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("[ERROR] Error reading user input; cmd.go -> ", err)
	}
	username = strings.TrimSuffix(username, "\n")

	color.PrintBold("Enter password:", color.LimeGreen)
	fmt.Println()
	color.PrintBold(">", color.RoyalPurple)
	input, err := term.ReadPassword(int(syscall.Stdin)) 
	if err != nil {
		log.Fatal("[ERROR] Error reading user input; cmd.go -> ", err)
	}
		
	encrypted_pass, err := crypto.Encrypt(input, key)
	if err != nil {
		log.Fatal("[ERROR] Error encrypting; cmd.go -> ", err)
	}
	fmt.Printf("Encrypted Pass: %x\n", encrypted_pass)

	base64_string := base64.StdEncoding.EncodeToString(encrypted_pass)

	creds.Entries = append(creds.Entries, struct{Name string "json:\"name\""; Username string "json:\"username\""; EncryptedPassword string "json:\"encrypted_password\""}{name, username, base64_string})

	b, err = json.Marshal(creds)

	os.WriteFile("gokey.json", b, 444)
}

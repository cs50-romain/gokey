package cmd

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	crypto "cs50-romain/gokey/internal/crypto"
	clipboard "cs50-romain/gokey/pkg/clipboard"
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

	command, err := buildCommand(args)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	
	if command.Main != "new" && command.Main != "copy" && command.Main != "show" && command.Main != "list" {
		color.PrintBold("[ERROR] Invalid command;", color.Red)
		fmt.Println()
	} else {
		if command.Main == "new" {
			err := storeCreds(args[2], key)
			if err != nil {
				log.Printf(err.Error())
				return 
			}
			color.PrintBold("Creating new creds for " + command.Name, color.ElectricPink)
			fmt.Println()
		} else if command.Main == "show" {
			color.PrintBold("Fetching creds for " + command.Name, color.ElectricPink)
			fmt.Println()
			showCreds(args[2], key)
			if err != nil {
				log.Printf(err.Error())
				return 
			}
		} else if command.Main == "copy" {
			color.PrintBold("Copying creds for " + command.Name, color.ElectricPink)
			copyCreds(args[2], key)
			if err != nil {
				fmt.Println()
				log.Printf(err.Error())
				return 
			}
		} else {
			color.PrintBold("List of saved credentials: ", color.ElectricPink)
			fmt.Println()
			listCreds()
			if err != nil {
				log.Printf(err.Error())
				return 
			}
		}
		fmt.Println()
	}
}

func buildCommand(args []string) (*Command, error) {
	var command Command

	if args[0] != "go" {
		return nil, errors.New("ERROR: Invalid command; go should be first argument")
	}

	if len(args) < 2 {
		return nil, errors.New("ERROR: Not enough arguments")
	} else {
		command.Main = args[1]
		if command.Main == "new" || command.Main == "show" || command.Main == "copy" {
			if len(args) < 3 {
				return nil, errors.New("ERROR: Not enough arguments")
			}
			command.Name = args[2]
		}
	}
	return &command, nil
}

func listCreds() error {
	var creds Creds
	b, err := os.ReadFile("gokey.json")
	if err != nil {
		return errors.New("[ERROR] Error reading gokey.json; cmd.go -> " + err.Error())
	}

	err = json.Unmarshal(b, &creds)
	if err != nil {
		return errors.New("[ERROR] Error unmarshaling gokey.json; cmd.go -> " + err.Error())
	}

	for _, entry := range creds.Entries {
		fmt.Println("Name: ", entry.Name)
	}
	return nil
}

func showCreds(name string, key []byte) {
	username, password, err := getCreds(name, key)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Printf("Username: %s\nPassword: %s\n", username, password)
}

func copyCreds(name string, key []byte) {
	_, password, err := getCreds(name, key)	
	if err != nil {
		log.Printf(err.Error())
		return
	}
	clipboard.CopytoClipboard(password)
}

func getCreds(name string, key []byte) (string, string, error) {
	var creds Creds
	b, err := os.ReadFile("gokey.json")
	if err != nil {
		return "", "", errors.New("[ERROR] Error reading gokey.json; cmd.go -> " + err.Error())
	}

	err = json.Unmarshal(b, &creds)
	if err != nil {
		return "", "", errors.New("ERROR: could not unmarshal gokey.json; cmd.go -> " + err.Error())
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
		return "", "", errors.New("ERROR: could not decode base64 password; cmd.go -> " + err.Error())
	}

	b, err = crypto.Decrypt(base64_password, key)
	if err != nil {
		return "", "", errors.New("ERROR: Decryption issue; cmd.go -> " + err.Error())
	}

	return username, string(b), nil
}

func storeCreds(name string, key []byte) error {
	var creds Creds
	b, err := os.ReadFile("gokey.json")
	if err != nil {
		return errors.New("ERROR: Error reading gokey.json; cmd.go -> " + err.Error())
	}
	err = json.Unmarshal(b, &creds)
	if err != nil {
		return errors.New("ERROR: Error unmarshaling config.json; cmd.go -> " + err.Error())
	}

	reader := bufio.NewReader(os.Stdin)
	color.PrintBold("Enter username:", color.LimeGreen)
	fmt.Println()
	color.PrintBold(">", color.RoyalPurple)
	username, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("ERROR Error reading user input; cmd.go -> " + err.Error())
	}
	username = strings.TrimSuffix(username, "\n")

	color.PrintBold("Enter password:", color.LimeGreen)
	fmt.Println()
	color.PrintBold(">", color.RoyalPurple)
	input, err := term.ReadPassword(int(syscall.Stdin)) 
	if err != nil {
		return errors.New("ERROR: Error reading user input; cmd.go -> " + err.Error())
	}
		
	encrypted_pass, err := crypto.Encrypt(input, key)
	if err != nil {
		return errors.New("ERROR: Error encrypting; cmd.go -> " + err.Error())
	}
	fmt.Printf("Encrypted Pass: %x\n", encrypted_pass)

	base64_string := base64.StdEncoding.EncodeToString(encrypted_pass)

	creds.Entries = append(creds.Entries, struct{Name string "json:\"name\""; Username string "json:\"username\""; EncryptedPassword string "json:\"encrypted_password\""}{name, username, base64_string})

	b, err = json.Marshal(creds)

	os.WriteFile("gokey.json", b, 444)
	return nil
}

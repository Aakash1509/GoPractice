package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Usage: Enter enough arguments <username> <password> <ip>")
	}
	username := os.Args[1]
	password := os.Args[2]
	ip := os.Args[3]

	//SSH connection configuration
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		log.Fatal("Failed to connect to SSH server:", err)
	}
	defer client.Close() // Close the connection

	// Create a SSH session for command
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create SSH session:", err)
	}
	defer session.Close() // Close the session when we're done

	output, err := session.CombinedOutput("hostname")
	if err != nil {
		log.Fatal("Failed to run hostname command:", err)
	}

	// Print the hostname of the connected machine
	fmt.Println("Connected to host:", string(output))
}

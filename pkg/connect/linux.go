package connect

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func SSH() {
	// SSH Connection for Linux
	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", serverIP), sshConfig)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Successfully connected to Linux server: %s\n", serverIP)
}

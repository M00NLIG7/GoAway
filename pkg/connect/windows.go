package connect

import (
	"os"

	"fmt"

	"github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"
	"github.com/hirochachacha/go-smb2"
	"github.com/stacktitan/smb/smb"
	"golang.org/x/crypto/ssh"
)

func main() {
	options := smb.Options{
		Host:        "192.168.1.1", // Replace with the IP address of the SMB server
		Port:        445,           // Default SMB port is 445
		User:        "username",    // Replace with the username to authenticate with
		Password:    "password",    // Replace with the password to authenticate with
		Domain:      "",            // Replace with the domain to authenticate with, if applicable
		ClientName:  "client",      // Replace with a name for the client
		Encrypt:     true,          // Set to true to encrypt the connection
		Signing:     true,          // Set to true to sign the connection
		UseNTLMv2:   true,          // Set to true to use NTLMv2 authentication
		DisableFQDN: true,          // Set to true to disable fully qualified domain name (FQDN) resolution
	}
	session, err := smb.NewSession(options, false)
	if err != nil {
		fmt.Printf("Failed to create SMB session: %s\n", err)
		os.Exit(1)
	}
	defer session.Close()

	cmd := "whoami"
	output, err := session.Exec(cmd)
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(output))
}

func CnC() {
	// Input parameters
	serverIP := "x.x.x.x"      // Replace with server IP address
	username := "yourusername" // Replace with your username
	password := "yourpassword" // Replace with your password

	// Try SSH connection first
	sshClient, err := connectSSH(serverIP, username, password)
	if err == nil {
		defer sshClient.Close()
		fmt.Printf("Successfully connected to %s using SSH\n", serverIP)
		// Execute SSH command here if needed
		return
	}
	fmt.Printf("Failed to connect to %s using SSH: %v\n", serverIP, err)

	// Try WinRM connection second
	winRMShell, err := connectWinRM(serverIP, username, password)
	if err == nil {
		defer winRMShell.Exit()
		fmt.Printf("Successfully connected to %s using WinRM\n", serverIP)
		// Execute WinRM command here if needed
		return
	}
	fmt.Printf("Failed to connect to %s using WinRM: %v\n", serverIP, err)

	// Try SMB connection last
	smbClient, err := connectSMB(serverIP, username, password)
	if err == nil {
		defer smbClient.Logoff()
		fmt.Printf("Successfully connected to %s using SMB\n", serverIP)
		// Execute SMB command here if needed
		return
	}
	fmt.Printf("Failed to connect to %s using SMB: %v\n", serverIP, err)

	// All connection methods failed
	fmt.Printf("Unable to connect to %s using any available method\n", serverIP)
}

func connectSSH(serverIP, username, password string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", serverIP+":22", config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func connectWinRM(serverIP, username, password string) (*powershell.Shell, error) {
	connInfo := &backend.WinRMConnectionInfo{
		Hostname: serverIP,
		Port:     5985,
		Username: username,
		Password: password,
	}
	shell, err := powershell.NewWinRMShell(connInfo)
	if err != nil {
		return nil, err
	}

	return shell, nil
}

func connectSMB(serverIP, username, password string) (*smb2.Client, error) {
	options := smb2.Options{
		Host:        serverIP,
		User:        username,
		Password:    password,
		Domain:      "",
		Share:       "",
		Port:        445,
		Encrypt:     true,
	}

	client, err := smb2.NewClient(options)
	if err != nil {
		return nil, err
	}

	return client, nil
}
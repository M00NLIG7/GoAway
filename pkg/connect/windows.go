package connect

import (
	"fmt"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/masterzen/winrm"
	"golang.org/x/crypto/ssh"
)

func main() {
	var connType int
	fmt.Println("Please select a connection type:")
	fmt.Println("1. SSH")
	fmt.Println("2. FTP")
	fmt.Println("3. WinRM")
	_, err := fmt.Scanln(&connType)
	if err != nil {
		fmt.Println("Invalid input:", err)
		os.Exit(1)
	}

	switch connType {
	case 1:
		// SSH
		sshSession, err := sshConnect()
		if err != nil {
			fmt.Println("SSH connection error:", err)
			os.Exit(1)
		}
		defer sshSession.Close()
		fmt.Println("SSH connection successful. You can now execute commands on the remote machine.")
	case 2:
		// FTP
		ftpConn, err := ftpConnect()
		if err != nil {
			fmt.Println("FTP connection error:", err)
			os.Exit(1)
		}
		defer ftpConn.Quit()
		fmt.Println("FTP connection successful. You can now interact with the remote file system.")
	case 3:
		// WinRM
		winrmClient, err := winrmConnect()
		if err != nil {
			fmt.Println("WinRM connection error:", err)
			os.Exit(1)
		}
		defer func() {
			if cerr := winrmClient.Close(); cerr != nil {
				fmt.Println("Error closing WinRM connection:", cerr)
			}
		}()
		fmt.Println("WinRM connection successful. You can now execute PowerShell commands on the remote machine.")
	default:
		fmt.Println("Invalid connection type.")
		os.Exit(1)
	}
}

func sshConnect() (*ssh.Session, error) {
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
	}
	client, err := ssh.Dial("tcp", "remote-machine:22", config)
	if err != nil {
		return nil, err
	}
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func ftpConnect() (*ftp.ServerConn, error) {
	conn, err := ftp.Dial("remote-machine:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}
	err = conn.Login("username", "password")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func winrmConnect() (*winrm.Client, error) {
	endpoint := winrm.NewEndpoint("remote-machine", 5985, false, false, nil, nil, nil, time.Minute)
	client, err := winrm.NewClient(endpoint, "username", "password")
	if err != nil {
		return nil, err
	}
	return client, nil
}

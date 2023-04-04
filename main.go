// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/masterzen/winrm"
// 	"golang.org/x/crypto/ssh"
// )

// func main() {
//     var session *ssh.Session
//     var client *winrm.Client
//     var err error

//     username := "root"
//     password := "password123"
//     server := "192.168.1.27"

//     // Try to establish an SSH connection
//     session, err = sshConnect(username, password, server)
//     if err == nil {
//         defer session.Close()
//         handleSSHSession(session)
//         return
//     }
//     fmt.Println("Failed to connect via SSH:", err)

//     // Try to establish a WinRM connection
//     client, err = winrmConnect(username, password, server)
//     if err == nil {
//         // defer client.
//         handleWinRMClient(client)
//         return
//     }
//     fmt.Println("Failed to connect via WinRM:", err)
// }

// func sshConnect(username, password, server string) (*ssh.Session, error) {
//     config := &ssh.ClientConfig{
//         User: username,
//         Auth: []ssh.AuthMethod{
//             ssh.Password(password),
//         },
//         HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//     }
//     client, err := ssh.Dial("tcp", server+":22", config)
//     if err != nil {
//         return nil, err
//     }
//     session, err := client.NewSession()
//     if err != nil {
//         return nil, err
//     }
//     return session, nil
// }

// func winrmConnect(username, password, server string) (*winrm.Client, error) {
// 	endpoint := winrm.NewEndpoint(server, 5985, false, false, nil, nil, nil, time.Minute)
// 	client, err := winrm.NewClient(endpoint, username, password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return client, nil
// }

// func handleSSHSession(session *ssh.Session) {
//     fmt.Println("SSH session established.")
//     defer fmt.Println("SSH session closed.")
//     session.Stdout = os.Stdout
//     session.Stderr = os.Stderr
//     session.Stdin = os.Stdin
//     session.Shell()
//     session.Wait()
// }

// func handleWinRMClient(client *winrm.Client) {
//     fmt.Println("WinRM connection established.")
//     defer fmt.Println("WinRM connection closed.")
//     shell, err := client.CreateShell()
//     if err != nil {
//         fmt.Println("Error creating shell:", err)
//         return
//     }
//     defer shell.Close()
//     reader := bufio.NewReader(os.Stdin)
//     for {
//         fmt.Print("PS> ")
//         cmd, _ := reader.ReadString('\n')
//         cmd = strings.TrimSpace(cmd)
//         if strings.ToLower(cmd) == "exit" || strings.ToLower(cmd) == "quit" {
//             break
//         }
//         cmd = "powershell -Command " + cmd
//         command, err := shell.Execute(cmd)
//         if err != nil {
//             fmt.Println("Error executing command:", err)
//             continue
//         }
//         fmt.Println(command.Stdout)
//         fmt.Println(command.Stdout)
//     }
// }




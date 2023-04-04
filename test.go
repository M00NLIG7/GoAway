package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/M00NLIG7/GoAway/assets"
)

func main() {
	// Get the Psexec.exe binary from the assets package
	bin, err := assets.Asset("Psexec.exe")
	if err != nil {
		panic(err)
	}

	// Create a command to execute the Psexec.exe binary from memory
	cmd := exec.Command("", "")
	cmd.Path = "cmd.exe"
	cmd.Stdin = bytes.NewReader(bin)

	// Add Psexec.exe arguments
	cmd.Args = append(cmd.Args, "/c", "Psexec.exe", "\\\\REMOTE_HOSTNAME", "-u", "USERNAME", "-p", "PASSWORD", "COMMAND")

	// Set the command's environment variables
	cmd.Env = os.Environ()

	// Set the command's working directory to the current directory
	cmd.Dir = "."

	// Redirect the command's standard input, output, and error to the current process's
	// standard input, output, and error
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("Done.")
}

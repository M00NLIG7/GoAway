package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/M00NLIG7/GoAway/assets"
)


func main() {
	// Retrieve the contents of Psexec.exe from the embedded assets
	data, err := assets.Asset("psexec.exe")
	if err != nil {
		panic(err)
	}

// Create a temporary file to write the binary data to
	f, err := ioutil.TempFile("", "psexec")
	if err != nil {
		log.Fatalf("failed to create temporary file: %s", err)
	}
	defer os.Remove(f.Name())

	// Write the binary data to the temporary file
	_, err = f.Write(data)
	if err != nil {
		log.Fatalf("failed to write binary data to temporary file: %s", err)
	}

	// Make the temporary file executable
	err = os.Chmod(f.Name(), 0700)
	if err != nil {
		log.Fatalf("failed to make temporary file executable: %s", err)
	}

	// PSEXEC.EXE \\LocalComputerIPAddress -u DOMAIN\my-user -p mypass CMD

	// Execute psexec.exe from the temporary file
	cmd := exec.Command(f.Name(), "\\192.168.1.21", "-u", "Administrator", "-p", "Password123", "CMD")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("failed to execute psexec.exe: %s", err)
	}

	// Print the output of the command
	log.Printf("psexec.exe output: %s", string(out))
}

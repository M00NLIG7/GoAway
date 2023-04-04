package main

import (
	"fmt"
	"io/ioutil"
	"myproject/assets"
	"os"
	"os/exec"
)


func main() {
	// Retrieve the contents of Psexec.exe from the embedded assets
	data, err := assets.Must
	if err != nil {
		panic(err)
	}

	// Write the contents of Psexec.exe to a temporary file
	tmpfile, err := ioutil.TempFile("", "psexec.*.exe")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up the temporary file when we're done
	_, err = tmpfile.Write(data)
	if err != nil {
		panic(err)
	}
	err = tmpfile.Close()
	if err != nil {
		panic(err)
	}

	// Execute the temporary file with the specified command-line arguments
	cmd := exec.Command(tmpfile.Name(), "\\\\REMOTE_HOSTNAME", "-u", "USERNAME", "-p", "PASSWORD", "COMMAND")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}

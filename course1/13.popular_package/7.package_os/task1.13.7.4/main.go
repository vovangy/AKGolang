package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecBin(binPath string, args ...string) string {
	cmd := exec.Command(binPath, args...)

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()

	if err != nil {
		return fmt.Sprintf("Error executing %s: %s", binPath, errBuf.String())
	}

	return outBuf.String()
}

func main() {
	fmt.Println("Listing files:")
	fmt.Println(ExecBin("ls", "-la"))

	fmt.Println("Attempting to execute a nonexistent binary:")
	fmt.Println(ExecBin("nonexistent-binary"))
}

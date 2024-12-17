package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Starting Go Dependency Vulnerability Scan...")
	cmd := exec.Command("govulncheck", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error during vulnerability scan:", err)
		os.Exit(1)
	}
	fmt.Println("Vulnerability scan completed successfully!")
}

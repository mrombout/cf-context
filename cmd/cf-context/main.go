package main

import (
	"os/exec"
	"os"
	"fmt"
	"crypto/rand"
	"log"
)

func main() {
	contextName := generateRandomId()
	if len(os.Args) > 1 {
		contextName = os.Args[1]
	}

	favouriteShell := os.Getenv("SHELL")
	if favouriteShell == "" {
		favouriteShell = "bash"
	}
	cmd := exec.Command(favouriteShell)

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("unable to determine user home: %v\n", err)
		os.Exit(1)
	}

	cfHomeDirectory := fmt.Sprintf("%s/.cf/%s/", userHomeDir, contextName)
	cfHomeEnvironmentVariable := fmt.Sprintf("CF_HOME=%s", cfHomeDirectory)

	contextEnvironment := append(os.Environ(), cfHomeEnvironmentVariable)
	cmd.Env = contextEnvironment

	fmt.Printf("switching context to %q...\n", contextName)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	_ = cmd.Run()

	fmt.Printf("cleaning up context %q...\n", contextName)
	err = os.RemoveAll(cfHomeDirectory)
	if err != nil {
		fmt.Printf("failed to remove context directory %q: %v\n", cfHomeDirectory, err.Error())
		os.Exit(2)
	}
}

func generateRandomId() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
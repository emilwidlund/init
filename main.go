package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/emilwidlund/init/utils"
)

func main() {
	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	if commandName == "" {
		fmt.Println("Provide a command to initialize")
		os.Exit(1)
	}

	script := resolveCommand(commandName, commandArgs)

	runCommand(script)

	// manifest := manifest.LoadManifest(".initrc")

	os.Exit(0)
}

func resolveCommand(commandName string, commandArgs []string) string {
	commands := map[string]string{
		"file": "touch",
		"dir":  "mkdir",
		"git":  "git init",
		"npm":  "npm init",

		// URL's
		"design":      open("https://www.figma.com/file/new"),
		"repository":  open("https://github.com/new"),
		"pledge":      open("https://polar.sh/new"),
		"tweet":       open("https://twitter.com/intent/post"),
		"domain":      open("https://vercel.com/domains"),
		"repl":        open("https://repl.it/new"),
		"codepen":     open("https://codepen.io/pen"),
		"codesandbox": open("https://codesandbox.io/new"),
		"gpt":         open("https://chat.openai.com"),
		"shader":      open("https://www.shadertoy.com/new"),
	}

	commandScript, ok := commands[commandName]

	if !ok {
		fmt.Printf("Command %s not found\n", commandName)

		fmt.Println("Available commands:")
		for key := range commands {
			fmt.Println("- ", key)
		}

		os.Exit(1)
	}

	command := strings.Join(append([]string{commandScript}, commandArgs...), " ")

	return command
}

func runCommand(script string) {
	currentPath, wdErr := os.Getwd()
	utils.Check(wdErr)

	cmd := exec.Command("bash", "-c", script)
	cmd.Dir = currentPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// opens the specified URL in the default browser of the user.
func open(url string) string {
	var cmd string

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}

	return cmd + " " + url
}

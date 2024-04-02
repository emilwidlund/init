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
	moduleName := os.Args[1]
	moduleArgs := os.Args[2:]

	if moduleName == "" {
		fmt.Println("Provide a module to initialize")
		os.Exit(1)
	}

	script := resolveModule(moduleName, moduleArgs)

	runModule(script)

	// manifest := manifest.LoadManifest(".initrc")

	os.Exit(0)
}

func resolveModule(moduleName string, moduleArgs []string) string {
	modules := map[string]string{
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
		"vercel":      open("https://vercel.com/new"),
		"repl":        open("https://repl.it/new"),
		"codepen":     open("https://codepen.io/pen"),
		"codesandbox": open("https://codesandbox.io/new"),
		"gpt":         open("https://chat.openai.com"),
		"shader":      open("https://www.shadertoy.com/new"),
	}

	moduleScript, ok := modules[moduleName]

	if !ok {
		fmt.Printf("Module %s not found\n", moduleName)
		os.Exit(1)
	}

	module := strings.Join(append([]string{moduleScript}, moduleArgs...), " ")

	return module
}

func runModule(script string) {
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

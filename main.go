package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var rootCmd = &cobra.Command{
	Use:   "go-check",
	Short: "A CLI tool that checks, formats, and builds Go code, without generating an executable",
	Run:   check,
}

func check(cmd *cobra.Command, args []string) {
	// Run "go vet" on the current directory.
	vetCmd := exec.Command("go", "vet")
	vetCmd.Stderr = os.Stderr
	if err := vetCmd.Run(); err != nil {
		fmt.Println("go vet failed")
		os.Exit(1)
	}

	// Run "go fmt" on the current directory.
	fmtCmd := exec.Command("go", "fmt", "./...")
	fmtCmd.Stderr = os.Stderr
	if err := fmtCmd.Run(); err != nil {
		fmt.Println("go fmt failed")
		os.Exit(1)
	}

	// Run "go build" on the current directory.
	buildCmd := exec.Command("go", "build", "./...")
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		fmt.Println("go build failed")
		os.Exit(1)
	}

	fmt.Println("go-check succeeded")
	fmt.Println("Your code is clean")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

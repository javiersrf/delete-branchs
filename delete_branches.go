package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getLocalBranches() ([]string, error) {
	cmd := exec.Command("git", "branch")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	branches := strings.Split(out.String(), "\n")
	var cleanedBranches []string
	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if strings.HasPrefix(branch, "*") {
			branch = strings.TrimSpace(branch[1:])
		}
		if branch != "" {
			cleanedBranches = append(cleanedBranches, branch)
		}
	}
	return cleanedBranches, nil
}

func deleteLocalBranch(branch string) error {
	cmd := exec.Command("git", "branch", "-D", branch)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to delete branch %s: %v", branch, err)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide branches to ignore.")
		os.Exit(1)
	}

	ignoreBranches := os.Args[1:]

	localBranches, err := getLocalBranches()
	if err != nil {
		fmt.Printf("Error getting local branches: %v\n", err)
		os.Exit(1)
	}

	for _, branch := range localBranches {
		if contains(ignoreBranches, branch) {
			fmt.Printf("Skipping branch: %s\n", branch)
		} else {
			fmt.Printf("Deleting branch: %s\n", branch)
			if err := deleteLocalBranch(branch); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/joho/godotenv"
)

func openRepo(repoPath string) (*git.Repository, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repoPath := "."

	repo, err := openRepo(repoPath)
	if err != nil {
		fmt.Println("Error opening repository:", err)
		os.Exit(1)
	}

	// Now you have the 'repo' object and can work with the Git repository
	// For example, you can get the HEAD reference like this:
	head, err := repo.Head()
	if err != nil {
		fmt.Println("Error getting HEAD reference:", err)
		os.Exit(1)
	}

	fmt.Println("HEAD reference:", head.Name())
}

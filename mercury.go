package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func openRepo(repoPath string) (*git.Repository, error) {
	return git.PlainOpen(repoPath)
}

func getDiffBetweenCommits(repo *git.Repository, commit1, commit2 *object.Commit) (object.Changes, error) {
	commit1Tree, err := commit1.Tree()
	if err != nil {
		return nil, err
	}

	commit2Tree, err := commit2.Tree()
	if err != nil {
		return nil, err
	}

	return commit1Tree.Diff(commit2Tree)
}

func preLine(input string, delim string) string {
	// Split the input string into lines
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	// Add "+" at the start of each line
	for i := range lines {
		lines[i] = delim + lines[i]
	}

	// Join the lines back into a single string
	result := strings.Join(lines, "\n")

	return result
}

func gitDiff(repo *git.Repository, from, to *object.Commit) (string, error) {
	diff, err := getDiffBetweenCommits(repo, from, to)
	if err != nil {
		return "", err
	}

	for _, change := range diff {
		fmt.Println(change.From.Name, change.To.Name)
		patch, err := change.Patch()
		if err != nil {
			fmt.Println("Error getting patch:", err)
			continue
		}

		for _, patch := range patch.FilePatches() {
			// from, to := patch.Files()
			// fmt.Println(from.Path, to.Path)
			// var previous diff.Chunk
			prev := patch.Chunks()[0]
			// cur := patch.Chunks()[1]
			next := patch.Chunks()[1]
			for _, chunk := range patch.Chunks() {
				switch chunk.Type() {
				case 0: // Equal
					fmt.Println("Equal")
				case 1: // Add
					if prev.Type() == 0 {
						fmt.Println(preLine(prev.Content(), "  "))
					}
					fmt.Println(preLine(chunk.Content(), "+ "))
					if next.Type() == 0 {
						fmt.Println(preLine(next.Content(), "  "))
					}
				case 2: // Delete
					fmt.Println(preLine(chunk.Content(), "- "))
				}
				prev = chunk
			}
			fmt.Println()
		}
	}

	return "hello", nil
}

func chatCompletion(prompt string) (openai.ChatCompletionResponse, error) {
	client := openai.NewClient(os.Getenv("OPENAI_TOKEN"))
	return client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repoPath := "."

	repo, err := openRepo(repoPath)
	if err != nil {
		fmt.Println("Error opening the repository:", err)
		os.Exit(1)
	}

	head, err := repo.Head()
	if err != nil {
		fmt.Println("Error getting head:", err)
		os.Exit(1)
	}

	child, err := repo.CommitObject(head.Hash())
	if err != nil {
		fmt.Println("Error getting commit objects:", err)
		os.Exit(1)
	}

	parent, err := child.Parent(0)
	if err != nil {
		fmt.Println("Error getting commit objects:", err)
		os.Exit(1)
	}

	fmt.Println(child.Hash)
	fmt.Println(parent.Hash)

	gitDiff(repo, parent, child)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	resp, err := chatCompletion("Hello!")
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Message.Content)
}

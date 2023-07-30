package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
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

func main() {
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

	from, err := repo.CommitObject(head.Hash())
	if err != nil {
		fmt.Println("Error getting commit objects:", err)
		os.Exit(1)
	}

	parent, err := from.Parent(0)
	if err != nil {
		fmt.Println("Error getting commit objects:", err)
		os.Exit(1)
	}

	fmt.Println(from.Hash)
	fmt.Println(parent.Hash)

	diff, err := getDiffBetweenCommits(repo, from, parent)
	if err != nil {
		fmt.Println("Error getting diff between commits:", err)
		os.Exit(1)
	}

	// Print the diff data
	for _, change := range diff {
		fmt.Println(change)
		patch, err := change.Patch()
		if err != nil {
			fmt.Println("Error getting patch:", err)
			continue
		}

		fmt.Println(patch.FilePatches())

		for _, patch := range patch.FilePatches() {
			// from, to := patch.Files()
			// fmt.Println(from.Path, to.Path)
			for index, chunk := range patch.Chunks() {
				fmt.Println(index, chunk.Content())
			}
			fmt.Println()
		}
	}
}

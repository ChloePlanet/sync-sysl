package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {
	org := os.Getenv("ORG")

	since := time.Date(2020, 8, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Since %s:\n", since)

	repos := []string{"conf-demo", "decimal", "decimal-reference", "go-course", "go-samplerest", "go-slides", "gosysl", "homebrew-sysl", "mermaid-go", "protoc-gen-sysl", "sysl", "sysl-catalog", "sysl-examples", "sysl-go", "sysl-go-demo", "sysl-playground", "sysl-template", "sysl-website", "syslgen-examples", "syslgen-proto", "test-rig-example", "vscode-sysl"}
	for _, repo := range repos {
		orgURL := fmt.Sprintf("https://github.com/%s/%s.git", org, repo)
		anzbankURL := fmt.Sprintf("https://github.com/anz-bank/%s.git", repo)
		CompareTwoRepo(repo, orgURL, anzbankURL, &since)
	}
	orgURL := fmt.Sprintf("https://github.com/%s/%s.git", org, "anz-pkg")
	anzbankURL := fmt.Sprintf("https://github.com/anz-bank/%s.git", "pkg")
	CompareTwoRepo("anz-pkg", orgURL, anzbankURL, &since)

	// arraiRepos := []string{"arrai", "wbnf"}
	// for _, repo := range repos {
	// 	orgURL := fmt.Sprintf("https://github.com/%s/%s.git", org, repo)
	// 	arraiURL := fmt.Sprintf("https://github.com/arr-ai/%s.git", repo)
	// 	CompareTwoRepo(repo, orgURL, arraiURL, &since)
	// }
}

func CompareTwoRepo(repo, aURL, bURL string, since *time.Time) {
	aLogs, err := GetLogs(aURL, since)
	if err != nil {
		fmt.Sprintf("💣 %-20s: get repo %s logs error: %s\n", repo, aURL, err.Error())
		return
	}
	bLogs, err := GetLogs(bURL, since)
	if err != nil {
		fmt.Sprintf("💣 %-20s: get repo %s logs error: %s\n", repo, bURL, err.Error())
		return
	}

	switch {
	case len(aLogs) == len(bLogs):
		if len(aLogs) > 0 && aLogs[0].ID() != bLogs[0].ID() {
			fmt.Printf("❓ %-20s: %s and %s are different\n", repo, aURL, bURL)
			return
		}
		fmt.Printf("✅ %-20s: %s and %s histories are the same\n", repo, aURL, bURL)
	case len(aLogs) > len(bLogs):
		if aLogs[len(aLogs)-len(bLogs)].ID() == bLogs[0].ID() {
			fmt.Printf("💛 %-20s: %s is ahead of %s by %d commits\n", repo, aURL, bURL, len(aLogs)-len(bLogs))
			return
		}
		fmt.Printf("%-20s: ❓ %s and %s are different\n", repo, aURL, bURL)
	case len(aLogs) < len(bLogs):
		if bLogs[len(bLogs)-len(aLogs)].ID() == aLogs[0].ID() {
			fmt.Printf("💛 %-20s: %s is behind %s by %d commits\n", repo, aURL, bURL, len(bLogs)-len(aLogs))
			return
		}
		fmt.Printf("❓ %-20s: %s and %s are different\n", repo, aURL, bURL)
	}
}

func GetLogs(url string, since *time.Time) ([]*object.Commit, error) {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Username: "bot",
			Password: os.Getenv("SYSL_GITHUB_TOKEN"),
		},
		SingleBranch: true,
	})
	if err != nil {
		return nil, err
	}

	ref, err := r.Head()
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: since})
	if err != nil {
		return nil, err
	}

	logs := make([]*object.Commit, 0)
	err = cIter.ForEach(func(c *object.Commit) error {
		logs = append(logs, c)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return logs, nil
}

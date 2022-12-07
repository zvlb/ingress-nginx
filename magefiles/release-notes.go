//go:build mage

package main

import (
	"context"
	"fmt"
	gogithub "github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
	"strings"
)

// LatestCommitLogs Retrieves the commit log between the latest two controller versions.
func (Release) LatestCommitLogs() {
	commitsBetweenTags()
}

func commitsBetweenTags() {
	tags := getAllControllerTags()
	Info("Getting Commits between %v and %v", tags[0], tags[1])
	commitLog, err := git("log", fmt.Sprintf("%v..%v", tags[1], tags[0]))
	if commitLog == "" {
		Warning("All Controller Tags is empty")
	}
	CheckIfError(err, "Retrieving Commit log")
	temp := strings.Split(commitLog, "\n")
	Info("There are %v commits", len(commitLog))
	Info("Commits between %v..%v", tags[1], tags[0])
	for i, s := range temp {
		Info("#%v Version %v", i, s)
	}
}

// Generate Release Notes
func (Release) ReleaseNotes() {
	makeReleaseNotes()
}

type GitHubClient struct {
	client            *gogithub.Client
	owner, repository string
}

func (Release) GithubReleaseNotes() {
	gh, err := githubClient()
	CheckIfError(err, "Get Latest Release Client error")
	release, _, err := gh.client.Repositories.GetLatestRelease(context.Background(), gh.owner, gh.repository)
	CheckIfError(err, "Get Latest Release")
	Info("Latest Release %v", release.String())
}

func githubClient() (*GitHubClient, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GITHUB_TOKEN},
	)
	oauthClient := oauth2.NewClient(context.Background(), ts)
	return &GitHubClient{
		client:     gogithub.NewClient(oauthClient),
		owner:      ORG,
		repository: REPO,
	}, nil
}
func makeReleaseNotes() {
}

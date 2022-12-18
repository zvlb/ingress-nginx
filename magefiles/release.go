//go:build mage

package main

import (
	"fmt"
	gogithub "github.com/google/go-github"
	"github.com/magefile/mage/mg"
	"os"
	"strings"
)

type Release mg.Namespace

var ORG = "strongjz"        // the owner so we can test from forks
var REPO = "ingress-nginx"  // the repo to pull from
var RELEASE_BRANCH = "main" //we only release from main
var GITHUB_TOKEN string     // the Google/gogithub lib needs an PAT to access the GitHub API

func init() {
	GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")

}

// Release Create a new release of ingress nginx controller
func (Release) Release(version string) {
	//update git controller tag TODO
	mg.Deps(mg.F(Tag.ControllerTag, version))
	//update ingress-nginx version
	mg.Deps(mg.F(Tag.BumpNginx, version))
	//update helm chart app version
	mg.Deps(mg.F(Helm.UpdateVersion))
	//make release notes TODO
	//update helm chart release notes TODO
	//update documentation with ingres-nginx version TODO
	//git commit TODO
	//make Pull Request TODO
	//make release
	//mg.Deps(mg.F(Release.CreateRelease, version))
}

//// CreateRelease Creates a new GitHub Release
//func (Release) CreateRelease(name string) {
//	releaser, err := gh_release.NewReleaser(ORG, REPO, GITHUB_TOKEN)
//	CheckIfError(err, "GitHub Release Client error")
//	newRelease, err := releaser.Create(fmt.Sprintf("controller-%s", name))
//	CheckIfError(err, "Create release error")
//	Info("New Release: Tag %v, ID: %v", newRelease.TagName, newRelease.ID)
//}

type ControllerImage struct {
	Digest string
}

type ReleaseNote struct {
	CurrentVersion       string
	PreviousVersion      string
	ControllerImages     []ControllerImage
	ImportantUpdates     []string
	AllOtherUpdates      []string
	NewContributors      []string
	PreviousReleaseNotes *ReleaseNote
}

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

func (Release) Latest() (*RepositoryRelease, error) {
	gh, err := githubClient()
	if err != nil {
		ErrorF("Error creating github client %s", err)
		return nil, error
	}

	release, _, err := gh.client.Repositories.GetLatestRelease(context.Background(), gh.owner, gh.repository)
	if err != nil {
		ErrorF("retrieving latest release %s", err)
		return nil, error
	}

	Info("Latest Release %v", release.String())
	return release, nil
}

func (Release) GithubReleaseNotes() {

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
	//previousReleaseNotes := ReleaseNote{}
	//newReleaseNotes := ReleaseNote{}

	//populate previous release note

	//current version
	//previous version
	//new contributors list
	//dependency_updates
	//all_updates
	//controller_image_digests
	//important_updates
}

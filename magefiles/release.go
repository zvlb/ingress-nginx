//go:build mage

package main

import (
	"context"
	"fmt"

	"github.com/magefile/mage/mg"
	gh_release "github.com/mysteriumnetwork/go-ci/github"
	"os"
)

type Release mg.Namespace

var ORG = "strongjz"        // the owner so we can test from forks
var REPO = "ingress-nginx"  // the repo to pull from
var RELEASE_BRANCH = "main" //we only release from main
var GITHUB_TOKEN string     // the google/gogithub lib needs an PAT to access the GitHub API

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
	mg.Deps(mg.F(Release.CreateRelease, version))
}

// CreateRelease Creates a new GitHub Release
func (Release) CreateRelease(name string) {
	releaser, err := gh_release.NewReleaser(ORG, REPO, GITHUB_TOKEN)
	CheckIfError(err, "Github Release Client error")
	newRelease, err := releaser.Create(fmt.Sprintf("controller-%s", name))
	CheckIfError(err, "Create release error")
	Info("New Release: Tag %v, ID: %v", newRelease.TagName, newRelease.ID)
}

// CurrentRelease retrieves the current Ingress Nginx Controller Release
func (Release) CurrentRelease() {
	gh, err := githubClient()
	CheckIfError(err, "Get Latest Release Client error")
	release, _, err := gh.client.Repositories.GetLatestRelease(context.Background(), gh.owner, gh.repository)
	CheckIfError(err, "Get Latest Release")
	Info("Latest Release %v", release.String())
}

//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
)

type Release mg.Namespace

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
	//make Pull Release TODO
}

// Generate Release Notes
func (Release) ReleaseNotes() {
	makeReleaseNotes()
}

func makeReleaseNotes() {

}
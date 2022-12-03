//go:build mage

package main

import (
	"fmt"
	semver "github.com/blang/semver/v4"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"strings"
)

type Tag mg.Namespace

func (Tag) Nginx() {
	tag, err := getIngressNGINXVersion()
	CheckIfError(err, "")
	fmt.Printf("%v", tag)
}

func getIngressNGINXVersion() (string, error) {
	dat, err := os.ReadFile("TAG")
	if err != nil {
		return "", err
	}
	datString := string(dat)
	//remove newline
	datString = strings.Replace(datString, "\n", "", -1)
	return datString, nil
}

func checkSemVer(currentTag, newTag string) bool {
	cTag, err := semver.Make(currentTag[1:])
	if err != nil {
		ErrorF("Error Current Tag %v Making Semver : %v", currentTag[1:], err)
		return false
	}
	nTag, err := semver.Make(newTag[1:])
	if err != nil {
		ErrorF("%v Error Making Semver %v \n", newTag, err)
		return false
	}

	err = nTag.Validate()
	if err != nil {
		ErrorF("%v not a valid Semver %v \n", newTag, err)
		return false
	}

	//The result will be
	//0 if newTag == currentTag
	//-1 if newTag< currentTag
	//+1 if newTag > currentTag.
	comp := nTag.Compare(cTag)
	if comp <= 0 {
		Warning("SemVer:%v is not an update\n", newTag)
		return false
	}
	return true
}

func bump(currentTag, newTag string) {
	//check if semver is valid
	if !checkSemVer(currentTag, newTag) {
		ErrorF("ERROR: Semver is not valid %v \n", newTag)
		os.Exit(1)
	}

	Debug("Updating Tag %v to %v \n", currentTag, newTag)
	err := os.WriteFile("TAG", []byte(newTag), 0666)
	CheckIfError(err, "Error Writing New Tag File")
}

func (Tag) BumpNginx(newTag string) {
	currentTag, err := getIngressNGINXVersion()
	CheckIfError(err, "Getting Ingress-nginx Version")

	bump(currentTag, newTag)
}

func (Tag) Git() {
	getGitTag()
}

func getGitTag() {
	git := sh.OutCmd("git")
	tag, err := git("describe", "--tags", "--abbrev=0")
	CheckIfError(err, "Reading Git data")
	Debug("Git Tag: %s", tag)
}

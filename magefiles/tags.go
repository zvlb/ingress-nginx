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

func getTag() (string, error) {
	var tag string
	dat, err := os.ReadFile("TAG")
	if err != nil {
		return "", err
	}
	//remove newline
	tag = strings.Replace(string(dat), "\n", "", -1)
	return tag, nil
}

// Generate Release Notes
func ReleaseNotes() {

}

type Tag mg.Namespace

func (Tag) Nginx() {
	tag, err := getTag()
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Printf("%v", tag)
}

func checkSemVer(currentTag, newTag string) bool {
	cTag, err := semver.Make(currentTag[1:])
	if err != nil {
		fmt.Printf("%v Not a valid Semver\n", newTag)
		return false
	}
	nTag, err := semver.Make(newTag[1:])
	if err != nil {
		fmt.Printf("%v Not a valid Semver\n", newTag)
		return false
	}

	err = nTag.Validate()
	if err != nil {
		fmt.Printf("%v Not a valid Semver\n", newTag)
		return false
	}

	//The result will be
	//0 if newTag == currentTag
	//-1 if newTag< currentTag
	//+1 if newTag > currentTag.
	comp := nTag.Compare(cTag)
	if comp <= 0 {
		fmt.Printf("SemVer:%v is not an update\n", newTag)
		return false
	}
	return true
}

func bump(currentTag, newTag string) {
	//check if semver is valid
	if !checkSemVer(currentTag, newTag) {
		fmt.Printf("ERROR: Semver is not valid %v\n", newTag)
		os.Exit(1)
	}

	fmt.Printf("Updating Tag %v to %v\n", currentTag, newTag)
	err := os.WriteFile("TAG", []byte(newTag), 0666)
	if err != nil {
		fmt.Printf("Error %v", err)
		os.Exit(1)
	}
}

func (Tag) Bump(newTag string) {
	currentTag, err := os.ReadFile("TAG")
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	bump(string(currentTag), newTag)
}
func (Tag) Git() {
	getGitTag()
}

func getGitTag() {
	git := sh.OutCmd("git")
	tag, err := git("describe", "--tags", "--abbrev=0")
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Printf("%s", tag)
}

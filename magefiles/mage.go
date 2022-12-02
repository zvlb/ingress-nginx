//go:build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg"

	"github.com/magefile/mage/sh"
	"os"
)

func getTag() (string, error) {
	dat, err := os.ReadFile("TAG")
	if err != nil {
		return "", err
	}

	return string(dat), nil
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

func bump(tag string) {
	dat, err := os.ReadFile("TAG")
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Printf("Updating Tag %v to %v\n", string(dat), tag)
	os.WriteFile("TAG", []byte(tag), 0666)
	if err != nil {
		fmt.Printf("Error %v", err)
	}
}

func (Tag) Bump(tag string) {
	bump(tag)
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

//go:build mage

package main

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
)

func getLatestGitTag() {
	var r, err = git.PlainOpen(".")
	if err != nil {
		fmt.Printf("Error Opening Repo %v", err)
	}
	iter, err := r.Tags()
	if err != nil {
		// Handle error
		fmt.Printf("Error Retreiving Tags %v", err)
	}

	if err := iter.ForEach(func(ref *plumbing.Reference) (string, error) {
		obj, err := r.TagObject(ref.Hash())
		switch err {
		case nil:
			// Tag object present
			return "", errors.New("Tag not present")
		case plumbing.ErrObjectNotFound:
			// Not a tag object
			return "", errors.New("Not a tag object")
		default:
			// Some other error
			return "", err
		}

		return obj.Name, nil
	}); err != nil {
		// Handle outer iterator error
		fmt.Printf("error reteirving tag details %v\n", err)
	}

}

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

func NginxTag() {
	tag, err := getTag()
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Printf("%v", tag)
}

func GitTag() {
	getLatestGitTag()
}

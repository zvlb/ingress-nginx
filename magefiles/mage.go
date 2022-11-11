//go:build mage

package main

import (
	"fmt"
	"os"
)

func getLatestGitTag() {

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
	tag, err := getTag()
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Printf("Tag is %v", tag)
}

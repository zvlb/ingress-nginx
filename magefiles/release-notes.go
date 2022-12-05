//go:build mage

package main

import (
	"fmt"
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

func makeReleaseNotes() {

}

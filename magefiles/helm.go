//go:build mage

package main

import (
	"fmt"
	"github.com/helm/helm/pkg/chartutil"
	"github.com/magefile/mage/mg"

	semver "github.com/blang/semver/v4"
	"os"
)

const HelmChartPath = "charts/ingress-nginx/Chart.yaml"
const HelmChartValues = "charts/ingress-nginx/values.yaml"

type Helm mg.Namespace

func (Helm) Read(file string) {

}

func (Helm) UpdateAppVersion() {
	updateAppVersion()
}

func updateAppVersion() {

}

func (Helm) UpdateVersion() {
	updateVersion()
}

func updateVersion() {
	fmt.Printf("Reading File %v\n", HelmChartPath)

	chart, err := chartutil.LoadChartfile(HelmChartPath)
	CheckIfError(err, "Error Loading Chart")

	//Get the current tag
	appVersionV, err := getNginxVer()
	CheckIfError(err, "Get Nginx Version")

	//remove the v from TAG
	appVersion := appVersionV[1:]

	fmt.Printf("Ignress Nginx App Version: %s Chart AppVersion: %s\n", appVersion, chart.AppVersion)
	if appVersion == chart.AppVersion {
		Warning("Ingress NGINX Version didnt change")
		return
	}

	//Update the helm chart
	chart.AppVersion = appVersion
	cTag, err := semver.Make(chart.Version)
	if err != nil {
		fmt.Printf("ERROR Creating Chart Version: %v", err)
		os.Exit(1)
	}

	if err = cTag.IncrementPatch(); err != nil {
		Info("ERROR Incrementing Chart Version: %v", err)
		os.Exit(1)
	}
	chart.Version = cTag.String()
	fmt.Printf("DEBUG: Updated Chart Version: %v", chart.Version)

	if err := chartutil.SaveChartfile(HelmChartPath, chart); err != nil {
		fmt.Printf("ERROR Saving new Chart: %v", err)
		os.Exit(1)
	}
}

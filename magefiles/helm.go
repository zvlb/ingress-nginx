//go:build mage

package main

import (
	"github.com/helm/helm/pkg/chartutil"
	"github.com/magefile/mage/mg"

	semver "github.com/blang/semver/v4"
	"os"
)

const HelmChartPath = "charts/ingress-nginx/Chart.yaml"
const HelmChartValues = "charts/ingress-nginx/values.yaml"

type Helm mg.Namespace

// UpdateAppVersion Updates the Helm App Version of Ingress Nginx Controller
func (Helm) UpdateAppVersion() {
	updateAppVersion()
}

func updateAppVersion() {

}

// UpdateVersion Update Helm Version of the Chart
func (Helm) UpdateVersion() {
	updateVersion()
}

func updateVersion() {
	Info("Reading File %v\n", HelmChartPath)

	chart, err := chartutil.LoadChartfile(HelmChartPath)
	CheckIfError(err, "Could not Load Chart")

	//Get the current tag
	appVersionV, err := getIngressNGINXVersion()
	CheckIfError(err, "Issue Retrieving the Ingress Nginx Version")

	//remove the v from TAG
	appVersion := appVersionV[1:]

	Info("Ingress-Nginx App Version: %s Chart AppVersion: %s\n", appVersion, chart.AppVersion)
	if appVersion == chart.AppVersion {
		Warning("Ingress NGINX Version didnt change")
		return
	}

	//Update the helm chart
	chart.AppVersion = appVersion
	cTag, err := semver.Make(chart.Version)
	CheckIfError(err, "Creating Chart Version: %v", err)

	if err = cTag.IncrementPatch(); err != nil {
		ErrorF("Incrementing Chart Version: %v", err)
		os.Exit(1)
	}
	chart.Version = cTag.String()
	Debug("Updated Chart Version: %v", chart.Version)

	err = chartutil.SaveChartfile(HelmChartPath, chart)
	CheckIfError(err, "Saving new Chart")
}

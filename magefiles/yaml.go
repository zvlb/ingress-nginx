//go:build mage

package main

import (
	"fmt"
	"github.com/helm/helm/pkg/chartutil"
	"github.com/magefile/mage/mg"
	"os"
	"strings"
)

const HelmChartPath = "charts/ingress-nginx/Chart.yaml"
const HelmChartValues = "charts/ingress-nginx/values.yaml"

type Yaml mg.Namespace

func (Yaml) Read(file string) {

}

func (Yaml) UpdateHelm() {
	updateHelmChart()
}

func updateHelmChart() {
	fmt.Printf("Reading File %v\n", HelmChartPath)

	//Read in the helm chart
	chartFile, err := os.ReadFile(HelmChartPath)
	if err != nil {
		fmt.Printf("Error Reading Yaml file:%v", err)
	}

	//marshal chart
	chart, err := chartutil.UnmarshalChartfile(chartFile)
	if err != nil {
		fmt.Errorf("Error Unmarshalling Chart: %v", err)
	}

	//Get the new tag
	tagV, err := getTag()
	if err != nil {
		fmt.Errorf("ERROR reading tag: %v", err)
		os.Exit(1)
	}

	//remove newline
	tagV = strings.Replace(tagV, "\n", "", -1)
	//remove the v from TAG
	tag := tagV[1:]

	fmt.Printf("TAG: %s Chart AppVersion: %s\n", tag, chart.AppVersion)
	//compare chart versions
	if tag == chart.AppVersion {
		fmt.Printf("EXITING TAG: %v Match Chart Verison: %v \n", tag, chart.AppVersion)
		os.Exit(1)
	}

	//Update the helm chart
	chart.AppVersion = tag

	if err := chartutil.SaveChartfile(HelmChartPath, chart); err != nil {
		fmt.Printf("ERROR Saving new Chart: %v", err)
		os.Exit(1)
	}
}

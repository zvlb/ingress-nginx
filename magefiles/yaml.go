//go:build mage

package main

import (
	"fmt"
	"github.com/helm/helm/pkg/chartutil"
	"github.com/magefile/mage/mg"
	"os"
)

type Yaml mg.Namespace

func (Yaml) Read(file string) {
	fmt.Printf("Reading File %v\n", file)

	chartFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error Reading Yaml file:%v", err)
	}
	chart, err := chartutil.UnmarshalChartfile(chartFile)
	if err != nil {
		fmt.Errorf("Error Unmarshelling Chart: %v", err)
	}
	fmt.Printf("Value: %#v\n", chart.Name)
}

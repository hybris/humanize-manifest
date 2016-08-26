package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v1"
	"github.com/hybris/gogobosh/local"
	"github.com/hybris/gogobosh/models"
)

func fatalIf(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func currentBoshManifest() string {
	configPath, err := local.DefaultBoshConfigPath()
	fatalIf(err)
	config, err := local.LoadBoshConfig(configPath)
	fatalIf(err)
	return config.CurrentDeploymentManifest()
}

func main() {
	manifestPath := os.Args[1]
	if manifestPath == "" {
		manifestPath = currentBoshManifest()
	}

	contents, err := ioutil.ReadFile(manifestPath)
	fatalIf(err)
	manifest := &models.DeploymentManifest{}
	yaml.Unmarshal(contents, manifest)

	str, err := yaml.Marshal(*manifest)
	fatalIf(err)
	fmt.Println(string(str))
}

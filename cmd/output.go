package cmd

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

func printResources(res interface{}) {
	switch CLI.Output {
	case "yaml":
		printResourcesYaml(res)
		break
	case "json":
		printResourcesJson(res)
	}
}

func printResourcesYaml(res interface{}) {
	out, err := yaml.Marshal(res)
	if err != nil {
		fmt.Printf("Error marshalling json: %s", err.Error())
	}

	fmt.Println(string(out))
}

func printResourcesJson(res interface{}) {
	out, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error marshalling json: %s", err.Error())
	}

	fmt.Println(string(out))
}

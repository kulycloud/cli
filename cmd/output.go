package cmd

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)



func printResources(res interface{}) error {
	switch CLI.Output {
	case "yaml":
		return printResourcesYaml(res)
	case "json":
		return printResourcesJson(res)
	}
	return fmt.Errorf("unknown output format \"%s\"", CLI.Output)
}

func printResourcesYaml(res interface{}) error {
	out, err := yaml.Marshal(res)
	if err != nil {
		return fmt.Errorf("error marshalling yaml: %w", err)
	}

	fmt.Println(string(out))
	return nil
}

func printResourcesJson(res interface{}) error {
	out, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("error marshalling yaml: %w", err)
	}

	fmt.Println(string(out))
	return nil
}

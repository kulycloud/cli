package cmd

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/kulycloud/cli/client"
)

type GetCmd struct {
	Type string `arg:"" name:"type" help:"Type of resource" enum:"service,route,s,r"`
	Resource string `arg:"" optional:"" name:"resource" help:"Resource to get"`
}

func (get *GetCmd) Run(ctx *kong.Context) error {
	c := client.NewClient(CLI.Instance)
	if get.Type == "service" || get.Type == "s" {
		if get.Resource != "" {
			res, err := c.GetService(CLI.Namespace, get.Resource)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return nil
			}
			printResources(res)

		} else {
			res, err := c.GetServices(CLI.Namespace)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return nil
			}

			printResources(res)
		}
	} else {
		if get.Resource != "" {
			res, err := c.GetRoute(CLI.Namespace, get.Resource)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return nil
			}
			printResources(res)

		} else {
			res, err := c.GetRoutes(CLI.Namespace)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return nil
			}

			printResources(res)
		}
	}

	return nil
}

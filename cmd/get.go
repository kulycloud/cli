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

func checkErrorAndPrint(res interface{}, err error) error {
	if err != nil {
		return err
	}
	return printResources(res)
}

func (get *GetCmd) Run(_ *kong.Context) error {
	c := client.NewClient(CLI.Instance)

	if get.Type == "service" || get.Type == "s" {
		if get.Resource != "" {
			return checkErrorAndPrint(c.GetService(CLI.Namespace, get.Resource))
		}
		return checkErrorAndPrint(c.GetServices(CLI.Namespace))
	} else if get.Type == "route" || get.Type == "r" {
		if get.Resource != "" {
			return checkErrorAndPrint(c.GetRoute(CLI.Namespace, get.Resource))
		}
		return checkErrorAndPrint(c.GetRoutes(CLI.Namespace))
	}
	return fmt.Errorf("unknown type \"%s\"", get.Type)
}



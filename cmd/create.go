package cmd

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/kulycloud/cli/client"
)

type CreateCmd struct {
	Type string `arg:"" name:"type" help:"Type of resource"`
	Resource string `arg:"" name:"resource" help:"Resource to get"`
	InputFile string `name:"file" short:"f" help:"Input file, if not specified use STDIN" optional:""`
}

func (create *CreateCmd) Run(_ *kong.Context) error {
	c := client.NewClient(CLI.Instance)

	if create.Type == "service" || create.Type == "s" {
		input, err := GetAndConvertInput(create.InputFile)
		if err != nil {
			return err
		}
		return checkErrorAndPrint(c.CreateService(CLI.Namespace, create.Resource, input))
	} else if create.Type == "route" || create.Type == "r" {
		input, err := GetAndConvertInput(create.InputFile)
		if err != nil {
			return err
		}
		return checkErrorAndPrint(c.CreateRoute(CLI.Namespace, create.Resource, input))
	}
	return fmt.Errorf("unknown type \"%s\"", create.Type)
}

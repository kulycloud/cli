package cmd

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/kulycloud/cli/client"
)

type UpdateCmd struct {
	Type string `arg:"" name:"type" help:"Type of resource"`
	Resource string `arg:"" name:"resource" help:"Resource to get"`
	InputFile string `name:"file" short:"f" help:"Input file, if not specified use STDIN" optional:""`
}

func (update *UpdateCmd) Run(_ *kong.Context) error {
	c := client.NewClient(CLI.Instance)

	if update.Type == "service" || update.Type == "s" {
		input, err := GetAndConvertInput(update.InputFile)
		if err != nil {
			return err
		}
		return checkErrorAndPrint(c.UpdateService(CLI.Namespace, update.Resource, input))
	} else if update.Type == "route" || update.Type == "r" {
		input, err := GetAndConvertInput(update.InputFile)
		if err != nil {
			return err
		}
		return checkErrorAndPrint(c.UpdateRoute(CLI.Namespace, update.Resource, input))
	}
	return fmt.Errorf("unknown type \"%s\"", update.Type)
}

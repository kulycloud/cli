package cmd

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/kulycloud/cli/client"
)

type DeleteCmd struct {
	Type string `arg:"" name:"type" help:"Type of resource"`
	Resource string `arg:"" name:"resource" help:"Resource to get"`
}

func (delete *DeleteCmd) Run(_ *kong.Context) error {
	c := client.NewClient(CLI.Instance)

	if delete.Type == "service" || delete.Type == "s" {
		return checkErrorAndPrint(c.DeleteService(CLI.Namespace, delete.Resource))
	} else if delete.Type == "route" || delete.Type == "r" {
		return checkErrorAndPrint(c.DeleteRoute(CLI.Namespace, delete.Resource))
	}
	return fmt.Errorf("unknown type \"%s\"", delete.Type)
}

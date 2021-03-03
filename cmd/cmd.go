package cmd

import (
	"github.com/alecthomas/kong"
)



type CreateCmd struct {
	Type string `arg:"" name:"type" help:"Type of resource"`
	Resource string `arg:"" name:"resource" help:"Resource to get"`
}

type UpdateCmd struct {
	Type string `arg:"" name:"type" help:"Type of resource"`
	Resource string `arg:"" name:"resource" help:"Resource to get"`
}

type DeleteCmd struct {
	Type string `arg:"" name:"type" help:"Type of resource"`
	Resource string `arg:"" name:"resource" help:"Resource to get"`
}

var CLI struct {
	Instance string `help:"Kuly instance" env:"KULY_INSTANCE" default:"http://localhost:8080" short:"i"`
	Namespace string `help:"Kuly namespace" env:"KULY_NAMESPACE" short:"n" required:""`
	Output string `help:"Output format" env:"KULY_OUTPUT" short:"o" default:"yaml" enum:"yaml,json"`
	Get GetCmd `cmd:"" help:"Get Resources"`
	Create CreateCmd `cmd:"" help:"Create Resources"`
	Update UpdateCmd `cmd:"" help:"Update Resources"`
	Delete DeleteCmd`cmd:"" help:"Delete Resources"`
}

func Run() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}

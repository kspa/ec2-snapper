package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
)

func main() {

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	// CLI stuff
	c := cli.NewCLI("ec2-snapper", "0.5.3a")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"create": func() (cli.Command, error) {
			return &CreateCommand{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorNone,
					ErrorColor:  cli.UiColorRed,
					WarnColor:   cli.UiColorYellow,
					InfoColor:   cli.UiColorGreen,
				},
			}, nil
		},
		"delete": func() (cli.Command, error) {
			return &DeleteCommand{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorNone,
					ErrorColor:  cli.UiColorRed,
					WarnColor:   cli.UiColorYellow,
					InfoColor:   cli.UiColorGreen,
				},
			}, nil
		},
		"report": func() (cli.Command, error) {
			return &ReportCommand{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorNone,
					ErrorColor:  cli.UiColorRed,
					WarnColor:   cli.UiColorYellow,
					InfoColor:   cli.UiColorGreen,
				},
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &VersionCommand{
				cliRef: *c,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		//fmt.Println(os.Stderr, err.Error())
		fmt.Println("Error: ", err.Error())
	}

	os.Exit(exitStatus)

}

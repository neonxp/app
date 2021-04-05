package app_test

import (
	"flag"
	"fmt"
	"os"

	"github.com/neonxp/app"
)

func ExampleApp() {
	app := app.New(commands)
	if err := app.Start(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var commands = []app.Command{}

type exampleCommand struct {
	test string
}

func (c *exampleCommand) Name() string {
	return "example"
}

func (c *exampleCommand) Usage() string {
	return "just example command"
}

func (c *exampleCommand) Flags() *flag.FlagSet {
	fs := flag.NewFlagSet(c.Name(), flag.ExitOnError)
	fs.StringVar(&c.test, "test", "value", "test flag")
	return fs
}

func (c *exampleCommand) Run(args []string) error {
	fmt.Printf("Runned %s command with flag -test=%s and args: %v", c.Name(), c.test, args)
	return nil
}

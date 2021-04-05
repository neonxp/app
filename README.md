# app

[![Go Reference](https://pkg.go.dev/badge/github.com/neonxp/app.svg)](https://pkg.go.dev/github.com/neonxp/app)

Simple cli applications runner.

## Usage

Commands must implements `app.Command` interface:

```go
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
```

Main file example:

```go
package main

import "github.com/neonxp/app"

func main() {
    app := app.New([]app.Command{
        exampleCommand{}, // <- Register command
    })
	if err := app.Start(os.Args); err != nil { // <- Run!
		fmt.Println(err)
		os.Exit(1)
	}
}
```

That's all!

[Full example](/example/example.go)

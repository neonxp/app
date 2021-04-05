package app

/*
This file is part of the App project.
Copyright (C) 2021  Alexander Kiryukhin <a.kiryukhin@mail.ru>
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

@license GPL-3.0+ <http://spdx.org/licenses/GPL-3.0+>
*/

import (
	"fmt"
	"os"
)

type App struct {
	commands []Command
}

var ErrNoCommand = fmt.Errorf("No command provided")

// New returns new App.
func New(commands []Command) *App {
	return &App{commands: commands}
}

// Start cli application
func (a *App) Start(args []string) error {
	if len(args) < 2 {
		return a.showHelpAndExit()
	}

	cmd := args[1]
	for _, c := range a.commands {
		if cmd == c.Name() {
			fs := c.Flags()
			if err := fs.Parse(args[2:]); err != nil {
				fs.Usage()
				return err
			}
			if err := c.Run(fs.Args()); err != nil {
				return err
			}
			return nil
		}
	}

	return a.showHelpAndExit()
}

func (a *App) showHelpAndExit() error {
	fmt.Printf("Usage: %s [command]\nCommands:\n", os.Args[0])
	for _, c := range a.commands {
		fmt.Printf("\t%s - %s", c.Name(), c.Usage())
	}
	return ErrNoCommand
}

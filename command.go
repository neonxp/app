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

import "flag"

// Command represents top level command of cli application
type Command interface {
	// Name returns name of command
	Name() string
	// Usage returns short help to command
	Usage() string
	// Flags returns FlagSet with flags for current command
	Flags() *flag.FlagSet
	// Run command
	Run(args []string) error
}

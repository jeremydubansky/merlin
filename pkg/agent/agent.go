// +build !windows

//Merlin is a post-exploitation command and control framework.
//This file is part of Merlin.
//Copyright (C) 2017  Russel Van Tuyl

//Merlin is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//Merlin is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
//along with Merlin.  If not, see <http://www.gnu.org/licenses/>.
package agent

import (
	"os/exec"
	"fmt"
	"github.com/mattn/go-shellwords"
)

func ExecuteCommand(name string, arg string) (stdout string, stderr string) {
	var cmd *exec.Cmd

	argS, errS := shellwords.Parse(arg)
	if errS != nil {fmt.Println("There was an error parsing command line argments")}

	cmd = exec.Command(name, argS...)

	out, err := cmd.CombinedOutput();
	stdout = string(out);
	stderr = "";

	if err != nil {
		stderr = err.Error();
	}

	return stdout, stderr
}
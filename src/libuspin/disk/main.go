//
// Copyright © 2016 Ikey Doherty <ikey@solus-project.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package disk provides convenience functions for manipulating disks and I/O
// functions within libuspin.
package disk

import (
	"os"
	"os/exec"
	"strings"
)

// ExecStdout is a convenience function to execute a command to the stdout
// and return the error, if any
func ExecStdout(command string) error {
	splits := strings.Fields(command)
	var c *exec.Cmd
	cmdName := splits[0]
	var err error
	// Search the path if necessary
	if !strings.Contains(cmdName, "/") {
		cmdName, err = exec.LookPath(cmdName)
		if err != nil {
			return err
		}
	}
	// Ensure we pass arguments
	if len(splits) == 1 {
		c = exec.Command(cmdName)
	} else {
		c = exec.Command(cmdName, splits[1:]...)
	}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
//
// Copyright 2021 Bryan T. Meyers <root@datadrake.com>
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

package cli

import (
	"github.com/DataDrake/cli-ng/cmd"
	"github.com/DataDrake/todo/tasks"
	"os"
)

func init() {
	cmd.Register(&Return)
}

// Return moves a task from the TODO list to the backlog
var Return = cmd.Sub{
	Name:  "return",
	Alias: ".",
	Short: "Return a task to the Backlog from TODO",
	Args:  &ReturnArgs{},
	Run:   ReturnRun,
}

// ReturnArgs specifies the ID of the task to return
type ReturnArgs struct {
	ID uint64 `desc:"ID of Task to return"`
}

// ReturnRun carries out the "return" sub-command
func ReturnRun(r *cmd.Root, s *cmd.Sub) {
	args := s.Args.(*ReturnArgs)
	if ok := tasks.Return(int(args.ID)); !ok {
		os.Exit(1)
	}
}

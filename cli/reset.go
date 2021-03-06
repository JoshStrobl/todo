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
	cmd.Register(&ResetCompleted)
}

// ResetCompleted permanently removes all completed tasks
var ResetCompleted = cmd.Sub{
	Name:  "reset-completed",
	Short: "Remove all Completed tasks",
	Run:   ResetCompletedRun,
}

// ResetCompletedRun carries out the "reset-completed" sub-command
func ResetCompletedRun(r *cmd.Root, s *cmd.Sub) {
	if ok := tasks.ResetCompleted(); !ok {
		os.Exit(1)
	}
}

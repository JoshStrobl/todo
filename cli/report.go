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
	cmd.Register(&Report)
}

// Report prints a list of every task, regardless of status
var Report = cmd.Sub{
	Name:  "report",
	Short: "Generate TODO.md summary of all tasks",
	Run:   ReportRun,
}

// ReportRun carries out the "report" sub-command
func ReportRun(r *cmd.Root, s *cmd.Sub) {
	if ok := tasks.Report(); !ok {
		os.Exit(1)
	}
}

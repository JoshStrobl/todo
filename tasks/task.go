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

package tasks

import (
	"fmt"
	"github.com/DataDrake/todo/colors"
	"io"
	"os"
	"strings"
	"time"
)

// Task is a work item which is tracked by this program
type Task struct {
	ID       int
	Created  time.Time
	Finished time.Time
	Project  string
	Label    string
	Name     string
}

// TaskFmt is the Scanf and Printf format string used for reading and saving Tasks
const TaskFmt = "T:%d C:%q F:%q P:%q L:%q N:%q\n"

// Read parses a Task from a file
func Read(r io.Reader) (t Task, err error) {
	var created, finished string
	_, err = fmt.Fscanf(r, TaskFmt, &t.ID, &created, &finished, &t.Project, &t.Label, &t.Name)
	if err != nil {
		return
	}
	if t.Created, err = time.Parse(time.RFC3339, created); err != nil {
		return
	}
	if len(finished) != 0 {
		t.Finished, err = time.Parse(time.RFC3339, finished)
	}
	return
}

// Write encodes a Task into a file
func (t Task) Write(w io.Writer) (err error) {
	created := t.Created.Format(time.RFC3339)
	var finished string
	if !t.Finished.IsZero() {
		finished = t.Finished.Format(time.RFC3339)
	}
	_, err = fmt.Fprintf(w, TaskFmt, t.ID, created, finished, t.Project, t.Label, t.Name)
	return
}

// Print writes a task to console
func (t Task) Print(tw io.Writer) (err error) {
	created := formatTime(t.Created)
	var finished string
	if !t.Finished.IsZero() {
		finished = formatTime(t.Finished)
	}
	pColor := colors.Color("projects", t.Project)
	lColor := colors.Color("labels", t.Label)
	_, err = fmt.Fprintf(tw, "\033[0m%d\t%s\t%s\t\033[%03dm \033[49m\t%s\t\033[%03dm \033[49m\t%s\t%s\033[0m\n", t.ID, created, finished, pColor, t.Project, lColor, t.Label, t.Name)
	return
}

// Report prints out the task as Markdown
func (t Task) Report(w io.Writer) (err error) {
	if !t.Finished.IsZero() {
		_, err = fmt.Fprintf(w, " - [x] %s\n", t.Name)
	} else {
		_, err = fmt.Fprintf(w, " - [ ] %s\n", t.Name)
	}
	return
}

// parse reads and validates a task specification
func parse(args []string) (t Task, ok bool) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "No task specified. Example: @project :label \"Task Name\"")
		return
	}
	if len(args) > 3 {
		fmt.Fprintln(os.Stderr, "Invalid task specified. Example: @project :label \"Task Name\"")
		return
	}
	for _, piece := range args {
		switch {
		case strings.HasPrefix(piece, "@"):
			t.Project, ok = ParseProject(piece)
		case strings.HasPrefix(piece, ":"):
			t.Label, ok = ParseLabel(piece)
		default:
			t.Name = piece
		}
	}
	if len(t.Name) == 0 {
		fmt.Fprintln(os.Stderr, "Task must have a Name.")
		return
	}
	ok = true
	return
}

func parsePrefixedName(token, prefix string) (name string, ok bool) {
	if len(token) < 2 {
		return
	}
	if !strings.HasPrefix(token, prefix) {
		return
	}
	fmt.Sscanf(token, prefix+"%q", &name)
	if len(name) == 0 {
		name = strings.TrimPrefix(token, prefix)
	}
	ok = true
	return
}

// ParseLabel decodes the name of a label
func ParseLabel(token string) (string, bool) {
	return parsePrefixedName(token, ":")
}

// ParseProject decodes the name of a project
func ParseProject(token string) (string, bool) {
	return parsePrefixedName(token, "@")
}

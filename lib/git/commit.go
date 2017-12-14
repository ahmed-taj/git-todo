// Copyright © 2017 Ahmed T. Ali <ah.tajelsir@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package git

import (
	// Native
	"os"
	"os/exec"

	// Ours
	"github.com/ahmed-taj/git-todos/lib/log"
)

// Commit runs `git commit` and returns the output/error strings
func Commit(commit CommitMessage) {
	cmd := exec.Command("git", "commit", "-m", commit.Format())

	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("Failed to commit. Did you forget to stage your changes?")
		log.Warn(`If you installed git-todos using snap (https://snapcraft.io) then you may
  need to tell git who you are!

  Run:
    git config --local user.email "you@example.com"
    git config --local user.name "Your Name"
`,
		)
		os.Exit(1)
	}
	log.Info("Your work has been commited")
}

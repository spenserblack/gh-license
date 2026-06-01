// Package git provides utilities for interfacing with Git.
package git

import (
	"os/exec"

	"github.com/cli/safeexec"
)

// Git is an interface that should implement git utilities.
type Git interface {
	// GetConfig gets a config value for a given key.
	GetConfig(key string) (string, error)
}

// Default returns the default Git implementation.
func Default() (Git, error) {
	path, err := safeexec.LookPath("git")
	git := git{
		path: path,
	}
	return git, err
}

// git is the default Git implementation.
type git struct {
	// path is the path to the git executable.
	path string
}

// output executes a git command and returns the output.
func (git git) output(subcommand string, args ...string) (string, error) {
	cmdArgs := make([]string, 0, 1+len(args))
	cmdArgs = append(cmdArgs, subcommand)
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command(git.path, cmdArgs...)
	bytes, err := cmd.Output()
	if err != nil {
		return "", err
	}
	s := string(bytes)
	return s, err
}

// GetConfig implements Git.
func (git git) GetConfig(key string) (string, error) {
	return git.output("config", key)
}

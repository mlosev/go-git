package git

import (
	"os/exec"
)

type runner interface {
	Run() error
}

var (
	execCommand func(...string) runner = func(args ...string) runner { return exec.Command("git", args...) }
)

// Init initializes a repository in dir, using the specified template.
func Init(dir, template string) error {
	args := []string{"init"}
	if template != "" {
		args = append(args, "--template='"+template+"'")
	}
	if dir != "" {
		args = append(args, dir)
	}
	return execCommand(args...).Run()
}

// Add adds the specified files to the working tree. If no files are provided all files will be added.
func Add(files ...string) error {
	args := []string{"add"}
	if len(files) == 0 {
		args = append(args, ".")
	} else {
		args = append(args, files...)
	}
	return execCommand(args...).Run()
}

// Remove removes the specified file from the working tree. If no files are provided all files will be removed.
func Remove(files ...string) error {
	args := []string{"rm"}
	if len(files) == 0 {
		args = append(args, "-r", ".")
	} else {
		args = append(args, files...)
	}
	return execCommand(args...).Run()
}

// Commit commits all changes from the working tree to the index.
func Commit(msg string) error {
	args := []string{"commit"}
	if msg != "" {
		args = append(args, "--message='"+msg+"'")
	} else {
		args = append(args, []string{"--allow-empty-message", "--message=''"}...)
	}
	return execCommand(args...).Run()
}

// Branch creates a new branch.
func Branch(name string) error {
	args := []string{"branch", name}
	return execCommand(args...).Run()
}

// DeleteBranch deletes an existing branch.
func DeleteBranch(name string) error {
	args := []string{"branch", "-d", name}
	return execCommand(args...).Run()
}

// Checkout checks out a branch.
func Checkout(branch string) error {
	args := []string{"checkout", branch}
	return execCommand(args...).Run()
}

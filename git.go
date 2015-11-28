package git

import (
	"errors"
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

// Clone clones the specified repository into dir.
// If dir is not provided the specified repository is cloned into the present working directory.
func Clone(repo, dir string) error {
	if repo == "" {
		return errors.New("go-git: Clone() no repository specified")
	}
	args := []string{"clone", repo}
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
func Remove(recursive bool, files ...string) error {
	args := []string{"rm"}
	if len(files) == 0 && !recursive {
		return errors.New("go-git: Remove() called without specifying files or recursive")
	} else if len(files) == 0 {
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
	if name == "" {
		return errors.New("go-git: Branch() no branch name specified")
	}
	return execCommand("branch", name).Run()
}

// DeleteBranch deletes an existing branch.
func DeleteBranch(name string) error {
	if name == "" {
		return errors.New("go-git: DeleteBranch() no branch name specified")
	}
	return execCommand("branch", "-d", name).Run()
}

// Checkout checks out a branch.
func Checkout(branch string) error {
	if branch == "" {
		return errors.New("go-git: Checkout() no branch name specified")
	}
	return execCommand("checkout", branch).Run()
}

// Tag creates a new tag with the provided name and message
func Tag(name, msg string) error {
	if name == "" {
		return errors.New("go-git: Tag() no tag name specified")
	}
	args := []string{"tag"}
	if msg != "" {
		args = append(args, "-m='"+msg+"'")
	} else {
		args = append(args, "-a")
	}
	args = append(args, name)
	return execCommand(args...).Run()
}

// DeleteTag deletes the named tag.
func DeleteTag(name string) error {
	if name == "" {
		return errors.New("go-git: DeleteTag() no tag name specified")
	}
	return execCommand("tag", "-d", name).Run()
}

// Merge Merges branch with the current branch.
func Merge(branch, msg string, fastforward bool) error {
	if branch == "" {
		return errors.New("go-git: Merge() called without specifying a branch")
	}
	args := []string{"merge", "-m='" + msg + "'"}
	if !fastforward {
		args = append(args, "--no-ff")
	}
	args = append(args, branch)
	return execCommand(args...).Run()
}

func RemoteAdd(name, location string) error {
	if name == "" {
		return errors.New("go-git: RemoteAdd() no name specified")
	}
	if location == "" {
		return errors.New("go-git: RemoteAdd() no location specified")
	}
	return execCommand("remote", "add", name, location).Run()
}

func RemoteRemove(name string) error {
	if name == "" {
		return errors.New("go-git: RemoteRemove() no name specified")
	}
	return execCommand("remote", "rm", name).Run()
}

func RemoteSetURL(name, location string) error {
	if name == "" {
		return errors.New("go-git: RemoteSetURL() no name specified")
	}
	if location == "" {
		return errors.New("go-git: RemoteSetURL() no location specified")
	}
	return execCommand("remote", "set-url", name, location).Run()
}

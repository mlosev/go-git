package git

import (
	"reflect"
	"testing"
)

type mockRunner struct{}

func (m *mockRunner) Run() error {
	return nil
}

func TestExecCommand(t *testing.T) {
	execCommand()
}

func TestInit(t *testing.T) {
	cases := []struct {
		CaseName   string
		Dir        string
		Template   string
		ExpectArgs []string
	}{
		{
			CaseName:   "Dir specified",
			Dir:        "repo-dir",
			Template:   "",
			ExpectArgs: []string{"init", "repo-dir"},
		},
		{
			CaseName:   "Dir not specified",
			Dir:        "",
			Template:   "",
			ExpectArgs: []string{"init"},
		},
		{
			CaseName:   "Template specified",
			Dir:        "",
			Template:   "template-dir",
			ExpectArgs: []string{"init", "--template='template-dir'"},
		},
		{
			CaseName:   "Dir and Template specified",
			Dir:        "repo-dir",
			Template:   "template-dir",
			ExpectArgs: []string{"init", "--template='template-dir'", "repo-dir"},
		},
	}
	for _, c := range cases {
		var gotArgs []string
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		Init(c.Dir, c.Template)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) {
			t.Errorf("\nexpected %v\ngot      : %v", c.ExpectArgs, gotArgs)
		}
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		CaseName   string
		Files      []string
		ExpectArgs []string
	}{
		{
			CaseName:   "No files",
			Files:      []string{},
			ExpectArgs: []string{"add", "."},
		},
		{
			CaseName:   "With files",
			Files:      []string{"file-1", "file-2"},
			ExpectArgs: []string{"add", "file-1", "file-2"},
		},
	}
	for _, c := range cases {
		var gotArgs []string
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		Add(c.Files...)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) {
			t.Errorf("%s\nexpected : %v\ngot      : %v", c.CaseName, c.ExpectArgs, gotArgs)
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		CaseName   string
		Files      []string
		ExpectArgs []string
	}{
		{
			CaseName:   "No files",
			Files:      []string{},
			ExpectArgs: []string{"rm", "-r", "."},
		},
		{
			CaseName:   "With files",
			Files:      []string{"file-1", "file-2"},
			ExpectArgs: []string{"rm", "file-1", "file-2"},
		},
	}
	for _, c := range cases {
		var gotArgs []string
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		Remove(c.Files...)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) {
			t.Errorf("%s\nexpected : %v\ngot      : %v", c.CaseName, c.ExpectArgs, gotArgs)
		}
	}
}

func TestCommit(t *testing.T) {
	cases := []struct {
		CaseName   string
		Msg        string
		ExpectArgs []string
	}{
		{
			CaseName:   "Commit message provided",
			Msg:        "commit message",
			ExpectArgs: []string{"commit", "--message='commit message'"},
		},
		{
			CaseName:   "No commit message provided",
			Msg:        "",
			ExpectArgs: []string{"commit", "--allow-empty-message", "--message=''"},
		},
	}
	for _, c := range cases {
		var gotArgs []string
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		Commit(c.Msg)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) {
			t.Errorf("%s\nexpected : %v\ngot      : %v", c.CaseName, c.ExpectArgs, gotArgs)
		}
	}
}

func TestBranch(t *testing.T) {
	cases := []struct {
		CaseName   string
		Name       string
		ExpectArgs []string
	}{
		{
			CaseName:   "Create a new branch",
			Name:       "new-branch",
			ExpectArgs: []string{"branch", "new-branch"},
		},
	}
	for _, c := range cases {
		var gotArgs []string
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		Branch(c.Name)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) {
			t.Errorf("%s\nexpected : %v\ngot      : %v", c.CaseName, c.ExpectArgs, gotArgs)
		}
	}
}

func TestDeleteBranch(t *testing.T) {
	cases := []struct {
		CaseName   string
		Name       string
		ExpectArgs []string
	}{
		{
			CaseName:   "Delete a branch",
			Name:       "branch-to-be-deleted",
			ExpectArgs: []string{"branch", "-d", "branch-to-be-deleted"},
		},
	}
	for _, c := range cases {
		var gotArgs []string
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		DeleteBranch(c.Name)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) {
			t.Errorf("%s\nexpected : %v\ngot      : %v", c.CaseName, c.ExpectArgs, gotArgs)
		}
	}
}

func TestCheckout(t *testing.T) {
	cases := []struct {
		CaseName   string
		Branch     string
		ExpectArgs []string
	}{
		{
			CaseName:   "Checkout a branch",
			Branch:     "branch",
			ExpectArgs: []string{"checkout", "branch"},
		},
	}
	for _, c := range cases {
		var gotArgs []string
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		Checkout(c.Branch)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) {
			t.Errorf("%s\nexpected : %v\ngot      : %v", c.CaseName, c.ExpectArgs, gotArgs)
		}
	}
}

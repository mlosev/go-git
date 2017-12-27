package git

import (
	"errors"
	"reflect"
	"testing"
)

type mockRunner struct{}

func (m *mockRunner) Run() error {
	return nil
}
func equalErr(errA, errB error) bool {
	if errA != nil && errB != nil {
		return errA.Error() == errB.Error()
	} else if errA == nil && errB == nil {
		return true
	}
	return false
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
		gotArgs := []string{}
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

func TestClone(t *testing.T) {
	cases := []struct {
		CaseName   string
		Repo       string
		Dir        string
		ExpectArgs []string
		ExpectErr  error
	}{
		{
			CaseName:   "Attempt to clone a repository without specifying a repository",
			Repo:       "",
			Dir:        "",
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: Clone() no repository specified"),
		},
		{
			CaseName:   "Clone a repository",
			Repo:       "repo-name",
			Dir:        "",
			ExpectArgs: []string{"clone", "repo-name"},
			ExpectErr:  nil,
		},
		{
			CaseName:   "Clone a repository into a specified directory",
			Repo:       "repo-name",
			Dir:        "dir-name",
			ExpectArgs: []string{"clone", "repo-name", "dir-name"},
			ExpectErr:  nil,
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := Clone(c.Repo, c.Dir)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
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
		gotArgs := []string{}
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
		Recursive  bool
		Files      []string
		ExpectArgs []string
		ExpectErr  error
	}{
		{
			CaseName:   "No files with recursive",
			Recursive:  true,
			Files:      []string{},
			ExpectArgs: []string{"rm", "-r", "."},
			ExpectErr:  nil,
		},
		{
			CaseName:   "No files without recursive",
			Recursive:  false,
			Files:      []string{},
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: Remove() called without specifying files or recursive"),
		},
		{
			CaseName:   "With files",
			Recursive:  false,
			Files:      []string{"file-1", "file-2"},
			ExpectArgs: []string{"rm", "file-1", "file-2"},
			ExpectErr:  nil,
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := Remove(c.Recursive, c.Files...)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
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
		gotArgs := []string{}
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
		ExpectErr  error
	}{
		{
			CaseName:   "Create a new branch",
			Name:       "new-branch",
			ExpectArgs: []string{"branch", "new-branch"},
			ExpectErr:  nil,
		},
		{
			CaseName:   "Create a new branch without specifying a name",
			Name:       "",
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: Branch() no branch name specified"),
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := Branch(c.Name)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
		}
	}
}

func TestDeleteBranch(t *testing.T) {
	cases := []struct {
		CaseName   string
		Name       string
		ExpectArgs []string
		ExpectErr  error
	}{
		{
			CaseName:   "Delete an unspecified branch",
			Name:       "",
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: DeleteBranch() no branch name specified"),
		},
		{
			CaseName:   "Delete a branch",
			Name:       "branch-to-be-deleted",
			ExpectArgs: []string{"branch", "-d", "branch-to-be-deleted"},
			ExpectErr:  nil,
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := DeleteBranch(c.Name)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
		}
	}
}

func TestCheckout(t *testing.T) {
	cases := []struct {
		CaseName   string
		Branch     string
		ExpectArgs []string
		ExpectErr  error
	}{
		{
			CaseName:   "Checkout a branch",
			Branch:     "branch",
			ExpectArgs: []string{"checkout", "branch"},
			ExpectErr:  nil,
		},
		{
			CaseName:   "Checkout an unspecified branch",
			Branch:     "",
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: Checkout() no branch name specified"),
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := Checkout(c.Branch)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v \ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
		}
	}
}

func TestTag(t *testing.T) {
	cases := []struct {
		CaseName   string
		Name       string
		Msg        string
		ExpectArgs []string
		ExpectErr  error
	}{
		{
			CaseName:   "Create a tag without specifying a tag name",
			Name:       "",
			Msg:        "",
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: Tag() no tag name specified"),
		},
		{
			CaseName:   "Create a tag without a message",
			Name:       "tag-name",
			Msg:        "",
			ExpectArgs: []string{"tag", "-a", "tag-name"},
			ExpectErr:  nil,
		},
		{
			CaseName:   "Create a tag with a message",
			Name:       "tag-name",
			Msg:        "tag-msg",
			ExpectArgs: []string{"tag", "-m='tag-msg'", "tag-name"},
			ExpectErr:  nil,
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := Tag(c.Name, c.Msg)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
		}
	}
}

func TestDeleteTag(t *testing.T) {
	cases := []struct {
		CaseName   string
		Name       string
		ExpectArgs []string
		ExpectErr  error
	}{
		{
			CaseName:   "Delete a tag without specifying a tag name",
			Name:       "",
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: DeleteTag() no tag name specified"),
		},
		{
			CaseName:   "Delete a tag",
			Name:       "tag-name",
			ExpectArgs: []string{"tag", "-d", "tag-name"},
			ExpectErr:  nil,
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := DeleteTag(c.Name)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
		}
	}
}

func TestMerge(t *testing.T) {
	cases := []struct {
		CaseName    string
		Branch      string
		Msg         string
		FastForward bool
		ExpectArgs  []string
		ExpectErr   error
	}{
		{
			CaseName:    "Merge a branch without specifying a branch",
			Branch:      "",
			Msg:         "",
			FastForward: true,
			ExpectArgs:  []string{},
			ExpectErr:   errors.New("go-git: Merge() called without specifying a branch"),
		},
		{
			CaseName:    "Merge a branch",
			Branch:      "branch-name",
			Msg:         "merge-message",
			FastForward: true,
			ExpectArgs:  []string{"merge", "-m='merge-message'", "branch-name"},
			ExpectErr:   nil,
		},
		{
			CaseName:    "Merge a branch without fastforwarding",
			Branch:      "branch-name",
			Msg:         "merge-message",
			FastForward: false,
			ExpectArgs:  []string{"merge", "-m='merge-message'", "--no-ff", "branch-name"},
			ExpectErr:   nil,
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := Merge(c.Branch, c.Msg, c.FastForward)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
		}
	}
}

func TestRemoteAdd(t *testing.T) {
	cases := []struct {
		CaseName   string
		Name       string
		Location   string
		ExpectArgs []string
		ExpectErr  error
	}{
		{
			CaseName:   "No name specified",
			Name:       "",
			Location:   "",
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: RemoteAdd() no name specified"),
		},
		{
			CaseName:   "No location specified",
			Name:       "remote-name",
			Location:   "",
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: RemoteAdd() no location specified"),
		},
		{
			CaseName:   "Add remote",
			Name:       "remote-name",
			Location:   "remote-location",
			ExpectArgs: []string{"remote", "add", "remote-name", "remote-location"},
			ExpectErr:  nil,
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := RemoteAdd(c.Name, c.Location)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
		}
	}
}

func TestRemoteRemove(t *testing.T) {
	cases := []struct {
		CaseName   string
		Name       string
		ExpectArgs []string
		ExpectErr  error
	}{
		{
			CaseName:   "No name specified",
			Name:       "",
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: RemoteRemove() no name specified"),
		},
		{
			CaseName:   "Remove remote",
			Name:       "remote-name",
			ExpectArgs: []string{"remote", "rm", "remote-name"},
			ExpectErr:  nil,
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := RemoteRemove(c.Name)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
		}
	}
}

func TestRemoteSetURL(t *testing.T) {
	cases := []struct {
		CaseName   string
		Name       string
		Location   string
		ExpectArgs []string
		ExpectErr  error
	}{
		{
			CaseName:   "No name specified",
			Name:       "",
			Location:   "",
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: RemoteSetURL() no name specified"),
		},
		{
			CaseName:   "No location specified",
			Name:       "remote-name",
			Location:   "",
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: RemoteSetURL() no location specified"),
		},
		{
			CaseName:   "Set remote url",
			Name:       "remote-name",
			Location:   "remote-location",
			ExpectArgs: []string{"remote", "set-url", "remote-name", "remote-location"},
			ExpectErr:  nil,
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := RemoteSetURL(c.Name, c.Location)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
		}
	}
}

func TestFetch(t *testing.T) {
	cases := []struct {
		CaseName   string
		Remote     string
		Branches   []string
		ExpectArgs []string
		ExpectErr  error
	}{
		{
			CaseName:   "No remote specified",
			Remote:     "",
			Branches:   []string{},
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: Fetch() no remote specified"),
		},
		{
			CaseName:   "No branches specified",
			Remote:     "remote-location",
			Branches:   []string{},
			ExpectArgs: []string{"fetch", "remote-location", "--all"},
			ExpectErr:  nil,
		},
		{
			CaseName:   "Fetch specific branches",
			Remote:     "remote-name",
			Branches:   []string{"branch-1", "branch-2", "branch-3"},
			ExpectArgs: []string{"fetch", "remote-name", "branch-1", "branch-2", "branch-3"},
			ExpectErr:  nil,
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := Fetch(c.Remote, c.Branches...)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
		}
	}
}

func TestPull(t *testing.T) {
	cases := []struct {
		CaseName   string
		Remote     string
		Branches   []string
		ExpectArgs []string
		ExpectErr  error
	}{
		{
			CaseName:   "No remote specified",
			Remote:     "",
			Branches:   []string{},
			ExpectArgs: []string{},
			ExpectErr:  errors.New("go-git: Pull() no remote specified"),
		},
		{
			CaseName:   "No branches specified",
			Remote:     "remote-location",
			Branches:   []string{},
			ExpectArgs: []string{"pull", "remote-location"},
			ExpectErr:  nil,
		},
		{
			CaseName:   "Pull specific branches",
			Remote:     "remote-name",
			Branches:   []string{"branch-1", "branch-2", "branch-3"},
			ExpectArgs: []string{"pull", "remote-name", "branch-1", "branch-2", "branch-3"},
			ExpectErr:  nil,
		},
	}
	for _, c := range cases {
		gotArgs := []string{}
		execCommand = func(args ...string) runner {
			gotArgs = args
			return &mockRunner{}
		}
		gotErr := Pull(c.Remote, c.Branches...)
		if !reflect.DeepEqual(c.ExpectArgs, gotArgs) || !equalErr(c.ExpectErr, gotErr) {
			t.Errorf("%s\nexpected : %v, %v\ngot      : %v, %v",
				c.CaseName,
				c.ExpectArgs, c.ExpectErr,
				gotArgs, gotErr,
			)
		}
	}
}

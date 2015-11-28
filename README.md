go-git
================================================================================

go-git is a simple go wrapper around git.

Installation
--------------------------------------------------------------------------------

Run 'go get' to install,
```bash
go get github.com/sgen/go-git
```

or clone the repository directly.
```bash
git clone http://github.com/sgen/go-git
```

Documentation
--------------------------------------------------------------------------------

Detailed documentation can be found at [https://godoc.org/github.com/sgen/go-git](https://godoc.org/github.com/sgen/go-git).


Usage
--------------------------------------------------------------------------------

Initialize a new repository.
```go
git.Init("repo-dir", "template-dir")
```

Add specific files to the working tree.
```go
git.Add("file1", "file2")
```

Remove specific files from the working tree.
```go
git.Remove("file1", "file2")
```

Commit changes to the index.
```go
git.Commit("commit msg")
```

ToDo
--------------------------------------------------------------------------------

Commands slated for addition.

Clone
Remote Add
Status
Push
Fetch
Pull
Merge
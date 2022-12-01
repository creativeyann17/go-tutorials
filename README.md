# Formation
https://www.youtube.com/watch?v=YS4e4q9oBaU&t=274s

# Setup workspace

Best way to *go* is to have a single workspace for all the go projects

In `./zshrc` add the following:

```
export GOROOT=~/Programs/go-1.19.3
export GOPATH=~/Workspaces/go
export GO111MODULE=auto
```

Init the workspace and add projects:

```
go work init
go work use go-tutorials project1 project2 ...
```

# Copy vs Reference

*Copy*
- primitives
- arrays
- structs

*Reference*
- slices
- maps

# Interfaces

Best practices:
- Use many, small interfaces
- `interface{} <=> any`
- design functions/methods to receive interfaces

# Modules
This is the name of the module, no folders will be created:
```
go mod init github.com/creativeyann17/tutorials/modules
```
Download deps:
```
go get gonum.org/v1/gonum/stat 
```
Update/Reset modules, based on the current imports:
```
go mod tidy
```

# Tests
Verbos + Cover (noth optional)
```
go test -v --cover
```
Benchmark
```
go test -bench=.
```

# Gorountines

`go run -race main.go` will detect when variables in race condition

# Env Vars.
```
export GOROOT=/home/foo/Programs/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/home/foo/Workspaces/go
```
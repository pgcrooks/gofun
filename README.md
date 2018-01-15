# Go Fun

This repository is to allow me to mess around with Go and learn.

## Getting Started

Grab the Go compiler:

```bash
sudo apt-get install golang
```

This repo is intended to be part of a larger workspace. As such, we assume the workspace looks like this:

```bash
bin/
pkg/
src/
  github.com/
    pgcrooks/
      hello/
```

This will require the following BASH environmental variables to be set (they are different from the defaults):

```bash
# Go paths
export GOPATH=$HOME/Documents/projects/go
export GOBIN=$GOPATH/bin
PATH=$PATH:${GOBIN}
```

The project can be built and installed into ```bin``` as follows:

```bash
go install github.com/pgcrooks/hello
```

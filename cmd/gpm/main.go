package main

import (
	"github.com/dhiguero/grpc-proto-manager/cmd/gpm/commands"
)

// Version of the command
var Version string

// Commit from which the command was built
var Commit string

func main() {
	commands.Execute(Version, Commit)
}

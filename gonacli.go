package main

import (
	"github.com/wenlng/gonacli/cli"
)

const (
	Name    = "goncali"
	Version = "v1.0.10"
)

func main() {
	c := cli.CLI{}
	c.Run(Name, Version)
}

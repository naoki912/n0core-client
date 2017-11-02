package main

import (
	"github.com/naoki912/n0core-client/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}

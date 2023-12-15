package main

import "github.com/reikanv/calc-cli/internal/cli"

func main() {
	xf, err := cli.ReadFlags()
	if err != nil {
		panic(err)
	}
	cli.Out(xf)
}

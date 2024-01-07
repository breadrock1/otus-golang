package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		printFatal("go-envdir: usage: go-envdir dir child")
	}
	env, err := ReadDir(args[0])
	if err != nil {
		printFatal(fmt.Errorf("go-envdir: fatal: %w", err))
	}
	code := RunCmd(args[1:], env)
	if code == 101 {
		printFatal(fmt.Errorf("go-envdir: unable to run %s", args[1]))
	}
	os.Exit(code)
}

func printFatal(a interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, a)
	os.Exit(101)
}

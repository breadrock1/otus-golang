package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := execCommand(cmd[0], cmd[1:])
	command.Env = extractEnvironment(env)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
	}

	return 0
}

func execCommand(bin string, args []string) *exec.Cmd {
	return exec.Command(bin, args...)
}

func extractEnvironment(env Environment) []string {
	environVars := make([]string, 0)
	for key, envVar := range env {
		if envVar.NeedRemove {
			continue
		}

		envValue := fmt.Sprintf("%s=%s", key, envVar.Value)
		environVars = append(environVars, envValue)
	}
	return environVars
}

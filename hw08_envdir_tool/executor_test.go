package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("test.sh", func(t *testing.T) {
		_ = os.Setenv("ADDED", "from original env")
		env, _ := ReadDir("./testdata/env")
		command := []string{"/bin/bash", "./testdata/echo.sh", "arg1=1", "arg2=2"}
		cmp := "HELLO is (\"hello\")\nBAR is (bar)\nFOO is (   foo\nwith new line)\n"
		cmp += "UNSET is ()\nADDED is (from original env)\nEMPTY is ()\narguments are arg1=1 arg2=2\n"

		var returnCode int
		result := capturer.CaptureStdout(func() {
			returnCode = RunCmd(command, env)
		})
		require.Equal(t, 0, returnCode)
		require.Equal(t, cmp, result)
	})

	t.Run("Success", func(t *testing.T) {
		// подготовка тестовых данных
		dir := os.TempDir()
		defer func() {
			_ = os.RemoveAll(dir)
		}()

		// папка с переменными окружения
		err := os.Mkdir(filepath.Join(dir, "vars"), 0o777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(dir, "vars", "BAR"), []byte("bar"), 0o666)
		require.NoError(t, err)
		// баш-скрипт (распечатывает свой первый аргумент и переменную окружения BAR)
		err = os.WriteFile(filepath.Join(dir, "t.sh"), []byte("#!/usr/bin/env bash\necho $1\necho $BAR\n"), 0o666)
		require.NoError(t, err)
		err = os.Chmod(filepath.Join(dir, "t.sh"), 0o777)
		require.NoError(t, err)
		// конец подготовки тестовых данных

		env, err := ReadDir(filepath.Join(dir, "vars"))
		require.NoError(t, err)

		var returnCode int
		result := capturer.CaptureStdout(func() {
			returnCode = RunCmd([]string{filepath.Join(dir, "t.sh"), "something"}, env)
		})
		require.Equal(t, 0, returnCode)
		require.Equal(t, "something\nbar\n", result)
	})

	t.Run("Rewrite HOME", func(t *testing.T) {
		// подготовка тестовых данных
		dir := os.TempDir()
		defer func() {
			_ = os.RemoveAll(dir)
		}()

		// папка с переменными окружения
		err := os.Mkdir(filepath.Join(dir, "vars"), 0o777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(dir, "vars", "HOME"), []byte("42"), 0o666)
		require.NoError(t, err)
		// баш-скрипт (распечатывает переменную окружения HOME)
		err = os.WriteFile(filepath.Join(dir, "t.sh"), []byte("#!/usr/bin/env bash\necho $HOME\n"), 0o666)
		require.NoError(t, err)
		err = os.Chmod(filepath.Join(dir, "t.sh"), 0o777)
		require.NoError(t, err)
		// конец подготовки тестовых данных

		env, err := ReadDir(filepath.Join(dir, "vars"))
		require.NoError(t, err)

		var returnCode int
		result := capturer.CaptureStdout(func() {
			returnCode = RunCmd([]string{filepath.Join(dir, "t.sh")}, env)
		})
		require.Equal(t, 0, returnCode)
		require.Equal(t, "42\n", result)
	})

	t.Run("Fail with wrong command", func(t *testing.T) {
		env, err := ReadDir("testdata/env")
		require.NoError(t, err)

		var returnCode int
		result := capturer.CaptureStderr(func() {
			returnCode = RunCmd([]string{"cat", "."}, env)
		})
		require.Equal(t, 1, returnCode)
		require.Equal(t, "cat: .: Is a directory\n", result)
	})
}

package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("Success with testdata", func(t *testing.T) {
		expectEnv := Environment{
			"BAR":   EnvValue{Value: "bar"},
			"FOO":   EnvValue{Value: "   foo\nwith new line"},
			"HELLO": EnvValue{Value: `"hello"`},
			"UNSET": EnvValue{NeedRemove: true},
			"EMPTY": EnvValue{NeedRemove: true},
		}
		env, err := ReadDir("testdata/env")
		require.NoError(t, err)
		require.Equal(t, env, expectEnv)
	})

	t.Run("Success with empty dir", func(t *testing.T) {
		dir, err := ioutil.TempDir("", "test")
		require.NoError(t, err)
		defer os.RemoveAll(dir)

		env, err := ReadDir(dir)
		require.NoError(t, err)
		require.Len(t, env, 0)
	})

	t.Run("Fail with dir not exists", func(t *testing.T) {
		_, err := ReadDir("some name")
		require.Error(t, err)
	})
}

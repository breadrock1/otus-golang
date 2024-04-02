package main

import (
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

		resultEnv := make(Environment)
		for key := range expectEnv {
			val, exists := env[key]
			if exists {
				resultEnv[key] = val
			}
		}

		require.NoError(t, err)
		require.Equal(t, resultEnv, expectEnv)
	})

	t.Run("Fail with dir not exists", func(t *testing.T) {
		_, err := ReadDir("some name")
		require.Error(t, err)
	})
}

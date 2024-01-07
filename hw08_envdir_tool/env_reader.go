package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	sysEnvs := loadSysEnvs()
	for _, entry := range dirEntries {
		entryName := entry.Name()
		_, exists := sysEnvs[entryName]
		if exists {
			delete(sysEnvs, entryName)
		}

		filePath := fmt.Sprintf("%s/%s", dir, entryName)
		fileData, err := readFileValue(filePath)
		var entryVar EnvValue
		if err != nil {
			entryVar = EnvValue{Value: fileData, NeedRemove: true}
			sysEnvs[entryName] = entryVar
		} else {
			fileData = strings.ReplaceAll(fileData, "\x00", "\n")
			entryVar = EnvValue{Value: fileData, NeedRemove: false}
			if len(fileData) < 1 {
				entryVar.NeedRemove = true
			}
			sysEnvs[entryName] = entryVar
		}
	}

	return sysEnvs, nil
}

func readFileValue(filePath string) (string, error) {
	handle, err := os.Open(filePath)
	defer handle.Close()
	if err != nil {
		return "", err
	}

	fileData, err := io.ReadAll(handle)
	if err != nil {
		return "", nil
	}

	strFileData := string(fileData)
	splitData := strings.Split(strFileData, "\n")
	strFileData = splitData[0]
	strFileData = strings.ReplaceAll(strFileData, "\x00", "\n")
	if strFileData == "" || strFileData == " " {
		return "", errors.New("empty file data")
	}

	return strFileData, nil
}

func loadSysEnvs() Environment {
	sysEnvs := make(Environment)
	for _, env := range os.Environ() {
		envArr := strings.Split(env, "=")
		value := strings.Join(envArr[1:], "=")
		envValue := EnvValue{Value: value, NeedRemove: false}
		sysEnvs[envArr[0]] = envValue
	}
	return sysEnvs
}

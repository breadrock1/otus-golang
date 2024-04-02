package main

import (
	"errors"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func openSourceFile(fromPath string, offset int64) (*os.File, error) {
	fromFileInfo, err := os.Stat(fromPath)
	if err != nil {
		return nil, os.ErrNotExist
	}

	if !fromFileInfo.Mode().IsRegular() {
		return nil, ErrUnsupportedFile
	}

	if fromFileInfo.Size() <= offset {
		return nil, ErrOffsetExceedsFileSize
	}

	fromFileHandle, err := os.Open(fromPath)
	if err != nil {
		return nil, os.ErrNotExist
	}

	return fromFileHandle, nil
}

func openTargetFile(toPath string) (*os.File, error) {
	toFileInfo, err := os.Stat(toPath)
	if err == nil {
		if !toFileInfo.Mode().IsRegular() {
			return nil, ErrUnsupportedFile
		}
		_ = os.Remove(toPath)
	}

	toFileHandle, err := os.Create(toPath)
	if err != nil {
		return nil, err
	}

	return toFileHandle, nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	if limit < 0 || offset < 0 {
		return errors.New("invalid seek parameters value")
	}

	// Create and get target file handle.
	toFileHandle, err := openTargetFile(toPath)
	if err != nil {
		return err
	}

	// Open and get source file handle.
	fromFileHandle, err := openSourceFile(fromPath, offset)
	if err != nil {
		return err
	}

	defer toFileHandle.Close()
	defer fromFileHandle.Close()

	// Seek by offset and check limit value.
	fromFileInfo, _ := os.Stat(fromPath)
	_, _ = fromFileHandle.Seek(offset, io.SeekStart)
	if limit == 0 {
		limit = fromFileInfo.Size()
	}

	// Create progressbar + defers.
	pBar := progressbar.New64(fromFileInfo.Size())
	defer pBar.Close()
	defer pBar.Finish()

	// Copy file data between files.
	_ = io.MultiWriter(toFileHandle, pBar)
	_, err = io.CopyN(toFileHandle, fromFileHandle, limit)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	return nil
}

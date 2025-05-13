package zdutil

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

// FolderExists checks if a given path corresponds to an existing directory.
// It returns true if the path is a directory, and false if it is not a directory
// or if there is an error accessing the path.
func FolderExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil || os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

// FileExists checks if a given path corresponds to an existing file.
// It returns true if the path is a regular file, and false if it is not a regular file
// or if there is an error accessing the path.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil || os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// RenameFile renames a file from `src` to `dst`. If the renaming fails during any part of the process
// (opening the source file, creating the destination file, copying the contents of the source file to
// the destination, or removing the original file), an error is returned with the first error encountered.
// Note that the destination file is closed after the copy is complete, but the source file is only closed
// if the copy is successful. If the copy fails, the source file is left open to allow the caller to
// inspect the contents of the source file.
func RenameFile(src, dst string) error {
	inputFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %s", err)
	}
	outputFile, err := os.Create(dst)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("failed to open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("failed to write to output file: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(src)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}
	return nil
}

// RenameDir renames a directory from `src` to `dst`, recursively renaming all of the files and
// subdirectories inside of `src`. If any of the files or directories inside of `src` have the same
// name as one of the strings in `ignoreExists`, they are skipped and not renamed.
func RenameDir(src, dst string, ignoreExists ...string) error {
	os.Mkdir(dst, 0644)
	err := filepath.WalkDir(src, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		newDst := dst + "/" + path[len(src):]
		if info.IsDir() {
			err = os.Mkdir(newDst, 0644)

			if err != nil && !errors.Is(err, fs.ErrExist) {
				return err
			}
			return nil
		}

		if FileExists(newDst) { // check if file should be ignored
			for _, ignore := range ignoreExists {
				if ignore == path[len(src):] {
					return nil
				}
			}
		}

		err = RenameFile(path, newDst)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// Rename attempts to rename a file or directory from `src` to `dst`.
// If `src` is a directory, it calls RenameDir to handle recursive renaming.
// If `src` is a file and the destination already exists with a name found
// in `ignoreExists`, the operation is skipped.
// Returns an error if the operation fails at any point.
func Rename(src, dst string, ignoreExists ...string) error {
	err := os.Rename(src, dst)
	if err == nil {
		return nil
	}

	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return RenameDir(src, dst, ignoreExists...)
	}

	if FileExists(dst) { // check if file should be ignored
		for _, ignore := range ignoreExists {
			if ignore == path.Base(src) {
				return nil
			}
		}
	}

	return RenameFile(src, dst)
}

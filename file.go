package zdgoutil

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

// FolderExists checks if a folder exists
func FolderExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil || os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

// FileExists checks if a file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil || os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// RenameFile renames a file. Mainly used when a docker image
// doesn't contains an os rename command by default
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

// RenameDir ranames a folder along with its subfolders and subfiles. Mainly used when
// a docker image doesn't contains an os rename command by default
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

// Rename raname file or folder, along with its subfolders and subfiles. Mainly used when
// a docker image doesn't contains an os rename command by default
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

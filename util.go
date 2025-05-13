package zdutil

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// GetExecPath returns the path of the current executable.
// It resolves symlinks and returns the final path.
// If the path is messed up, it will return an error.
func GetExecPath() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return path, err
	}

	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		return path, err
	}

	index := strings.LastIndex(path, "/")
	if index == -1 {
		index = strings.LastIndex(path, "\\")
		if index == -1 {
			return path, fmt.Errorf("exec path is messed up [path=%s]", path)
		}
	}
	path = path[:index]

	return path, err
}

// GetCallerDir returns the directory of the caller's file.
// It uses runtime.Caller to retrieve the file path of the caller,
// and path.Dir to extract the directory from the file path.
func GetCallerDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// Recover recovers from a panic and calls the given function with the recovered
// value and the given arguments.
func Recover(fn func(interface{}, ...interface{}), args ...interface{}) {
	if r := recover(); r != nil {
		fn(r, args...)
	}
}

// OnExit calls the given function with the given arguments and the received signal when the
// process receives any of the following signals: syscall.SIGQUIT, syscall.SIGINT,
// syscall.SIGILL, syscall.SIGTERM.
func OnExit(fn func(os.Signal, ...interface{}), args ...interface{}) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGILL, syscall.SIGTERM)

	sig := <-sigChan
	fmt.Println(" - cleanup request started with", sig, "signal")
	fn(sig, args...)
}

// OnExitWithContext calls the given function with the given arguments and the received signal when the
// process receives any of the following signals: syscall.SIGQUIT, syscall.SIGINT,
// syscall.SIGILL, syscall.SIGTERM, or when the given context is canceled.
func OnExitWithContext(ctx context.Context, fn func(os.Signal, ...interface{}), args ...interface{}) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGILL, syscall.SIGTERM)

	var sig os.Signal
	select {
	case <-ctx.Done():
		return
	case sig = <-sigChan:
	}

	fmt.Println(" - cleanup request started with", sig, "signal")
	fn(sig, args...)
}

// InTestingMode checks if the current program is running in test mode.
// It checks if the name of the current executable ends with ".test" or ".test.exe".
func InTestingMode() bool {
	return strings.HasSuffix(os.Args[0], ".test") || strings.HasSuffix(os.Args[0], ".test.exe")
}

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

// GetExecPath returns the source directory of executable
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

// GetCallerDir returns base directory of where it was called
func GetCallerDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// Recover quality of life function that call a function after panic
func Recover(fn func(interface{}, ...interface{}), args ...interface{}) {
	if r := recover(); r != nil {
		fn(r, args...)
	}
}

// OnExit calls function before exiting program
func OnExit(fn func(os.Signal, ...interface{}), args ...interface{}) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGILL, syscall.SIGTERM)

	sig := <-sigChan
	fmt.Println(" - cleanup request started with", sig, "signal")
	fn(sig, args...)
}

// OnExitWithContext calls function before exiting program and if context cancel, will stop listening for on exit events
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

// InTestingMode returns true if the parent caller is a unit test function
func InTestingMode() bool {
	return strings.HasSuffix(os.Args[0], ".test") || strings.HasSuffix(os.Args[0], ".test.exe")
}

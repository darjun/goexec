package goexec

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Option func(*exec.Cmd)

func WithStdin(stdin io.Reader) Option {
	return func(c *exec.Cmd) {
		c.Stdin = stdin
	}
}

func WithStdout(stdout io.Writer) Option {
	return func(c *exec.Cmd) {
		c.Stdout = stdout
	}
}

func WithStderr(stderr io.Writer) Option {
	return func(c *exec.Cmd) {
		c.Stderr = stderr
	}
}

func WithOutWriter(out io.Writer) Option {
	return func(c *exec.Cmd) {
		c.Stdout = out
		c.Stderr = out
	}
}

func WithEnv(key, value string) Option {
	return func(c *exec.Cmd) {
		c.Env = append(os.Environ(), fmt.Sprintf("%s=%s", key, value))
	}
}

func applyOptions(cmd *exec.Cmd, opts []Option) {
	for _, opt := range opts {
		opt(cmd)
	}
}

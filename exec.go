package goexec

import (
	"bytes"
	"os/exec"
)

func RunCommand(cmd string, arg []string, opts ...Option) error {
	c := exec.Command(cmd, arg...)
	applyOptions(c, opts)
	return c.Run()
}

func CombinedOutput(cmd string, arg []string, opts ...Option) ([]byte, error) {
	c := exec.Command(cmd, arg...)
	applyOptions(c, opts)
	return c.CombinedOutput()
}

func CombinedOutputString(cmd string, arg []string, opts ...Option) (string, error) {
	output, err := CombinedOutput(cmd, arg, opts...)
	return string(output), err
}

func Output(cmd string, arg []string, opts ...Option) ([]byte, error) {
	c := exec.Command(cmd, arg...)
	applyOptions(c, opts)
	return c.Output()
}

func OutputString(cmd string, arg []string, opts ...Option) (string, error) {
	output, err := Output(cmd, arg, opts...)
	return string(output), err
}

func SeparateOutput(cmd string, arg []string, opts ...Option) ([]byte, []byte, error) {
	var stdout, stderr bytes.Buffer
	c := exec.Command(cmd, arg...)
	applyOptions(c, opts)
	c.Stdout = &stdout
	c.Stderr = &stderr
	err := c.Run()
	return stdout.Bytes(), stderr.Bytes(), err
}

func SeparateOutputString(cmd string, arg []string, opts ...Option) (string, string, error) {
	stdout, stderr, err := SeparateOutput(cmd, arg, opts...)
	return string(stdout), string(stderr), err
}

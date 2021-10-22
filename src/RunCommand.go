package main

import (
	"io"
	"os/exec"
)

func RunCommand(name string, arg ...string) (io.ReadCloser, io.ReadCloser, error) {
	cmd := exec.Command(name, arg...)

	StdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	StdErrPipe, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, nil, err
	}

	return StdoutPipe, StdErrPipe, nil
}

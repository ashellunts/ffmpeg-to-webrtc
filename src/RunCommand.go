package main

import (
	"io"
	"os/exec"
)

func RunCommand(name string, arg ...string) (io.ReadCloser, error) {
	cmd := exec.Command(name, arg...)

	dataPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return dataPipe, nil
}

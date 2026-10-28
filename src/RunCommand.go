package main

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"sync"
)

// commandReadCloser waits for the child after stdout is drained. StdoutPipe
// must be read before Wait is called.
type commandReadCloser struct {
	stdout   io.ReadCloser
	cmd      *exec.Cmd
	waitOnce sync.Once
	waitErr  error
}

func (r *commandReadCloser) wait() error {
	r.waitOnce.Do(func() {
		r.waitErr = r.cmd.Wait()
	})
	return r.waitErr
}

func (r *commandReadCloser) Read(p []byte) (int, error) {
	n, err := r.stdout.Read(p)
	if err == io.EOF {
		if waitErr := r.wait(); waitErr != nil {
			return n, waitErr
		}
	}
	return n, err
}

func (r *commandReadCloser) Close() error {
	closeErr := r.stdout.Close()
	if errors.Is(closeErr, os.ErrClosed) {
		closeErr = nil
	}
	waitErr := r.wait()
	if closeErr != nil {
		return closeErr
	}
	return waitErr
}

func RunCommand(name string, arg ...string) (io.ReadCloser, error) {
	cmd := exec.Command(name, arg...)

	dataPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		_ = dataPipe.Close()
		return nil, err
	}

	return &commandReadCloser{stdout: dataPipe, cmd: cmd}, nil
}

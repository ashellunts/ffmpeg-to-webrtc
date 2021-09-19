package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCommand(t *testing.T) {
	require := require.New(t)

	dataPipe, err := RunCommand("test_program.exe")
	require.NoError(err)

	program_output := make([]byte, 0)
	for {
		tmp := make([]byte, 4)
		n, err := dataPipe.Read(tmp)
		if err != nil {
			break
		}
		program_output = append(program_output, tmp[0:n]...)
	}

	expected := []byte{72, 101, 108, 108, 111, 10}
	require.Equal(expected, program_output)
}

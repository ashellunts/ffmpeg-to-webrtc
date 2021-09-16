package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCommand(t *testing.T) {
	require := require.New(t)

	dataPipe, err := RunCommand("test_program.exe")
	require.NoError(err)

	program_output := make([]byte, 100)
	n, err := dataPipe.Read(program_output)
	require.NoError(err)
	require.Equal(6, n)

	expected := []byte{72, 101, 108, 108, 111, 10}
	require.Equal(expected, program_output[0:n])
}

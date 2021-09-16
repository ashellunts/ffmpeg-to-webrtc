package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	assert := assert.New(t)

	dataPipe, err := RunCommand("test_program.exe")
	assert.Nil(err)

	program_output := make([]byte, 100)
	n, err := dataPipe.Read(program_output)
	assert.Nil(err)
	assert.Equal(6, n)

	expected := []byte{72, 101, 108, 108, 111, 10}
	assert.Equal(expected, program_output[0:n])
}

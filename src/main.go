package main

import (
	"fmt"
)

func main() {
	dataPipe, err := RunCommand("test_program.exe")

	if err != nil {
		return
	}

	program_output := make([]byte, 100)
	n, err := dataPipe.Read(program_output)
	if err != nil {
		return
	}

	fmt.Print(program_output[0:n])
}

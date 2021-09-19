package main

import (
	"fmt"
)

func main() {
	dataPipe, err := RunCommand("test_program.exe")

	if err != nil {
		return
	}

	program_output := make([]byte, 0)
	for {
		tmp := make([]byte, 4)
		n, err := dataPipe.Read(tmp)
		if err != nil {
			break
		}
		program_output = append(program_output, tmp[0:n]...)
	}

	fmt.Print(program_output)
}

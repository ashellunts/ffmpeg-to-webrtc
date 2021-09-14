package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func execCmd(cmd *exec.Cmd) error {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0775)
	if err != nil {
		return errors.New("could not create log. " + err.Error())
	}

	cmd.Stderr = logFile

	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}

func main() {
	cmd := exec.Command("git", "status")
	fmt.Println("RUN - ", cmd.Args)

	dataPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("could not create named pipe ", err)
		return
	}

	if err := execCmd(cmd); err != nil {
		fmt.Println("could not run ", err)
		return
	}

	for {
		framebytes := make([]byte, 600000)
		n, err := dataPipe.Read(framebytes)
		if err != nil {
			fmt.Println("could not read pipe. ", err)
			return
		} else {
			fmt.Println("read bytes", framebytes[0:n])
		}
	}
}

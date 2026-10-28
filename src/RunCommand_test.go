package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

const runCommandHelperEnv = "FFMPEG_TO_WEBRTC_RUN_COMMAND_HELPER"

func TestRunCommand(t *testing.T) {
	t.Setenv(runCommandHelperEnv, "success")

	dataPipe, err := RunCommand(os.Args[0], "-test.run=^TestRunCommandHelperProcess$")
	if err != nil {
		t.Fatalf("RunCommand() error = %v", err)
	}

	output, readErr := io.ReadAll(dataPipe)
	closeErr := dataPipe.Close()
	if readErr != nil {
		t.Fatalf("reading successful helper output: %v", readErr)
	}
	if closeErr != nil {
		t.Fatalf("closing successful helper output: %v", closeErr)
	}
	if got, want := string(output), "Hello\n"; got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
}

func TestRunCommandReturnsProcessFailure(t *testing.T) {
	t.Setenv(runCommandHelperEnv, "failure")

	dataPipe, err := RunCommand(os.Args[0], "-test.run=^TestRunCommandHelperProcess$")
	if err != nil {
		t.Fatalf("RunCommand() error = %v", err)
	}

	output, readErr := io.ReadAll(dataPipe)
	closeErr := dataPipe.Close()
	if got, want := string(output), "synthetic output\n"; got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
	if readErr == nil || !strings.Contains(readErr.Error(), "exit status 23") {
		t.Fatalf("reading failed helper: error = %v, want exit status 23", readErr)
	}
	if closeErr == nil || !strings.Contains(closeErr.Error(), "exit status 23") {
		t.Fatalf("closing failed helper: error = %v, want exit status 23", closeErr)
	}
}

// TestRunCommandHelperProcess is invoked only as a child process by the tests
// above. It stands in for ffmpeg without reading user media or accessing a
// real capture device, peer, or network service.
func TestRunCommandHelperProcess(t *testing.T) {
	switch os.Getenv(runCommandHelperEnv) {
	case "success":
		_, _ = io.WriteString(os.Stdout, "Hello\n")
		os.Exit(0)
	case "failure":
		_, _ = io.WriteString(os.Stdout, "synthetic output\n")
		os.Exit(23)
	}
}

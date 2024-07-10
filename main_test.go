package main

import (
	"bytes"
	"os"
	"testing"
)

// TestMainFunction tests the main function's behavior.
// This is a basic example that assumes cmd.Execute() prints something to stdout.
// Adjust the test according to the actual behavior you expect from cmd.Execute().
func TestMainFunction(t *testing.T) {
	var (
		err error
	)

	// Capture stdout
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Execute main function
	main()

	// Close write-end of the pipe and restore stdout
	if err = w.Close(); err != nil {
		return
	}
	os.Stdout = originalStdout

	// Read output
	var buf bytes.Buffer

	if _, err = buf.ReadFrom(r); err != nil {
		return
	}

	output := buf.String()

	// Assert output is as expected
	expectedOutput := "Hello World!\n"
	if output != expectedOutput {
		t.Errorf("Unexpected output from main. Got '%s', want '%s'", output, expectedOutput)
	}
}

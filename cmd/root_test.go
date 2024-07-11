package cmd

import (
	"bytes"
	"os"
	"testing"
)

// TestNewRootCmd tests the NewRootCmd function for expected command properties.
func TestNewRootCmd(t *testing.T) {
	cmd := NewRootCmd()

	// Assert the command is not nil
	if cmd == nil {
		t.Fatal("NewRootCmd() returned nil")
	}

	// Assert the command properties
	if cmd.Use != "mono" {
		t.Errorf("expected Use to be 'mono', got '%s'", cmd.Use)
	}

	if cmd.Short != "A CLI tool to help with tasks common to dealing with monorepos." {
		format := "expected Short to be 'A CLI tool to help with tasks common to dealing with monorepos.', got '%s'"
		t.Errorf(format, cmd.Short)
	}
}

// TestExecute tests the Execute function for expected behavior.
func TestExecute(t *testing.T) {
	// Redirect stdout to capture output
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Execute the function
	Execute()

	// Close write-end of the pipe and restore stdout
	w.Close()
	os.Stdout = originalStdout

	// Read output
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		return
	}
	output := buf.String()

	// Assert no error occurred and possibly check output
	// Since Execute exits on error, reaching this point means no error occurred.
	// If your Execute function is expected to produce specific output, you can assert it here.
	expectedOutput := "Hello World!\n" // Adjust according to the actual expected output
	if output != expectedOutput {
		t.Errorf("Unexpected output from Execute. Got '%s', want '%s'", output, expectedOutput)
	}
}

package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	target "github.com/ocrosby/mono/cmd"
	"github.com/spf13/cobra"
)

// Setup function that optionally returns a teardown function.
func setup() (*cobra.Command, *bytes.Buffer, func()) {
	// Create a new instance of VersionCmd
	cmd := target.NewVersionCmd()

	// Perform setup tasks here (if any)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// Return a function to perform teardown tasks
	return cmd, b, func() {
		// Teardown tasks here
	}
}

func TestVersionCommand(t *testing.T) {
	// Arrange
	cmd, b, teardown := setup()
	t.Cleanup(teardown)

	// Act
	if err := cmd.Execute(); err != nil {
		t.Fatalf("versionCmd.Execute() failed: %v", err)
	}

	// Assert
	output := b.String()
	expectedVersionPrefix := "Mono version"
	if !strings.HasPrefix(output, expectedVersionPrefix) {
		t.Fatalf("expected \"%s\" to start with \"%s\"", output, expectedVersionPrefix)
	}

	if versionString, err := target.ReadVersionFromRoot(); err != nil {
		t.Fatalf("ReadVersionFromRoot() failed: %v", err)
	} else if !strings.Contains(output, versionString) {
		t.Fatalf("expected \"%s\" to contain \"%s\"", output, versionString)
	}
}

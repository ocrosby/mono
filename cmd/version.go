package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// findProjectRoot searches for the project root directory by looking for a unique marker.
// In this example, we're using the presence of a ".git" directory as the marker.
func findProjectRoot() (string, error) {
	var (
		dir string
		err error
	)

	if dir, err = os.Getwd(); err != nil {
		return "", fmt.Errorf("error getting current directory: %w", err)
	}

	for {
		// Check if the ".git" directory exists in the current directory
		if _, err = os.Stat(filepath.Join(dir, ".git")); !os.IsNotExist(err) {
			return dir, nil // Found the project root
		}

		// Move up one directory
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// If the parent directory is the same as the current directory, we've reached the filesystem root
			return "", errors.New("project root not found")
		}
		dir = parentDir
	}
}

// ReadVersionFromRoot reads the VERSION file from the project root directory.
func ReadVersionFromRoot() (string, error) {
	var (
		content          []byte
		err              error
		workingDirectory string
	)
	rootDir, err := findProjectRoot()
	if err != nil {
		return "", fmt.Errorf("could not find project root: %w", err)
	}

	// Assuming the VERSION file is in the root directory of your project
	// and your project structure allows for this relative path access.
	versionFilePath := filepath.Join(rootDir, "VERSION")

	if content, err = os.ReadFile(versionFilePath); err != nil {
		workingDirectory, err = os.Getwd()
		if err != nil {
			err = fmt.Errorf("could not read version file: %w", err)
		} else {
			err = fmt.Errorf("could not read version file: %w. Working directory: %s", err, workingDirectory)
		}

		return "", err
	}

	return string(content), nil
}

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Mono",
		Long:  `All software has versions. This is Mono's`,
		Run: func(cmd *cobra.Command, _ []string) {
			var (
				versionString string
				err           error
			)

			if versionString, err = ReadVersionFromRoot(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error reading versionString: %v\n", err)
				os.Exit(1)
			}

			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Mono versionString %s\n", versionString)
		},
	}
}

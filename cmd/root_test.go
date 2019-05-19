package cmd_test

import (
	"testing"

	"github.com/benmatselby/hagen/cmd"
)

func TestNewRootCommand(t *testing.T) {
	cmd := cmd.NewRootCommand()

	use := "hagen"
	short := "CLI application for retrieving data from GitHub"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

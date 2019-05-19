package cmd_test

import (
	"testing"

	"github.com/benmatselby/hagen/cmd"
	"github.com/benmatselby/hagen/pkg"
	"github.com/golang/mock/gomock"
)

func TestNewIssuesCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := pkg.NewMockProvider(ctrl)

	cmd := cmd.NewIssuesCommand(client)

	use := "issues"
	short := "List issues given the search criteria"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected short: %s; got %s", short, cmd.Short)
	}

	if cmd.Flag("count").DefValue != "-1" {
		t.Fatalf("expected count default to be: -1; got %s", cmd.Flag("count").DefValue)
	}
}

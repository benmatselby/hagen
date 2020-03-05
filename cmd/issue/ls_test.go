package issue_test

import (
	"testing"

	"github.com/benmatselby/hagen/cmd/issue"
	"github.com/benmatselby/hagen/pkg"
	"github.com/golang/mock/gomock"
)

func TestNewLsCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := pkg.NewMockProvider(ctrl)

	cmd := issue.NewLsCommand(client)

	use := "ls"
	short := "List issues given the search criteria. Default query is to list issues where the author is ${GITHUB_OWNER}"

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

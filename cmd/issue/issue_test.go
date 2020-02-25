package issue_test

import (
	"testing"

	"github.com/benmatselby/hagen/cmd/issue"
	"github.com/benmatselby/hagen/pkg"
	"github.com/golang/mock/gomock"
)

func TestNewIssueCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := pkg.NewMockProvider(ctrl)
	cmd := issue.NewIssueCommand(client)

	use := "issue"
	short := "Issue related commands"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

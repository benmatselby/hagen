package repo_test

import (
	"testing"

	"github.com/benmatselby/hagen/cmd/repo"
	"github.com/benmatselby/hagen/pkg"
	"github.com/golang/mock/gomock"
)

func TestNewRepoCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := pkg.NewMockProvider(ctrl)
	cmd := repo.NewRepoCommand(client)

	use := "repo"
	short := "Repository related commands"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

package repo_test

import (
	"testing"

	"github.com/benmatselby/hagen/cmd/repo"
	"github.com/benmatselby/hagen/pkg"
	"github.com/golang/mock/gomock"
)

func TestNewRepoLsCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := pkg.NewMockProvider(ctrl)
	cmd := repo.NewRepoLsCommand(client)

	use := "ls"
	short := "List the repositories"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

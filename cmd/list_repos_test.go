package cmd_test

import (
	"testing"

	"github.com/benmatselby/hagen/cmd"
	"github.com/benmatselby/hagen/pkg"
	"github.com/golang/mock/gomock"
)

func TestNewRepoLsCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := pkg.NewMockProvider(ctrl)
	cmd := cmd.NewListReposCommand(client)

	use := "repos"
	short := "List the repositories based on a query. Default query is to list repos by ${GITHUB_OWNER}"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

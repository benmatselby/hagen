package cmd_test

import (
	"bufio"
	"bytes"
	"errors"
	"testing"

	"github.com/benmatselby/hagen/cmd"
	"github.com/benmatselby/hagen/pkg"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/github"
)

func TestNewBurndownCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := pkg.NewMockProvider(ctrl)

	cmd := cmd.NewBurndownCommand(client)

	use := "burndown"
	short := "Provide a burndown table for a project"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

func TestDisplayBurndown(t *testing.T) {
	tt := []struct {
		name        string
		projectName string
		output      string
		projectErr  error
		issueErr    error
		columnErr   error
	}{
		{name: "return error if no project is supplied", projectName: "", output: "", projectErr: errors.New("a project needs to be specified --project"), columnErr: nil, issueErr: nil},
		{name: "return error if issues return an error", projectName: "ava-maria", output: "", projectErr: nil, columnErr: nil, issueErr: errors.New("boom")},
		{name: "return error if columns return an error", projectName: "ava-maria", output: "", projectErr: nil, columnErr: errors.New("boom"), issueErr: nil},
		{name: "can return tabular data for project", projectName: "ava-maria", output: `Column      Cards Story Points
------      ----- ------------
To do       1     21
In progress 2     18
-----       ----- ------------
Total       3     39
-----       ----- ------------
`, projectErr: nil, columnErr: nil, issueErr: nil},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := pkg.NewMockProvider(ctrl)

			columnOneName := "To do"
			columnOneID := int64(1)
			columnOne := github.ProjectColumn{
				ID:   &columnOneID,
				Name: &columnOneName,
			}

			columnTwoName := "In progress"
			columnTwoID := int64(2)
			columnTwo := github.ProjectColumn{
				ID:   &columnTwoID,
				Name: &columnTwoName,
			}

			columns := []*github.ProjectColumn{
				&columnOne, &columnTwo,
			}

			issueOneName := "(21) Massive"
			issueOne := &github.Issue{
				Title: &issueOneName,
			}

			issuesOne := []*github.Issue{
				issueOne,
			}

			issueTwoName := "(13) Attach"
			issueTwo := &github.Issue{
				Title: &issueTwoName,
			}

			issueThreeName := "(5) Teardrop"
			issueThree := &github.Issue{
				Title: &issueThreeName,
			}

			issuesTwo := []*github.Issue{
				issueTwo, issueThree,
			}

			client.
				EXPECT().
				ListColumnsForProject(gomock.Eq(tc.projectName), gomock.Eq("org-nick-cave"), gomock.Eq("repo-into-my-arms")).
				Return(columns, tc.columnErr).
				AnyTimes()

			client.
				EXPECT().
				ListIssuesForProjectColumn(gomock.Eq(columnOneID)).
				Return(issuesOne, tc.issueErr).
				AnyTimes()

			client.
				EXPECT().
				ListIssuesForProjectColumn(gomock.Eq(columnTwoID)).
				Return(issuesTwo, tc.issueErr).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			args := cmd.BurndownOptions{
				Project: tc.projectName,
				Org:     "org-nick-cave",
				Repo:    "repo-into-my-arms",
			}

			err := cmd.DisplayBurndown(client, args, writer)
			writer.Flush()

			if tc.projectErr != nil && err != nil {
				if err.Error() != tc.projectErr.Error() {
					t.Fatalf("expected error '%s'; got '%s'", tc.projectErr, err)
				}
			}

			if b.String() != tc.output {
				t.Fatalf("expected '%s'; got '%s'", tc.output, b.String())
			}
		})
	}
}

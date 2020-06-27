package cmd_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/benmatselby/hagen/cmd"
	"github.com/benmatselby/hagen/pkg"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/github"
	"github.com/spf13/viper"
)

func TestListProjectsCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := pkg.NewMockProvider(ctrl)
	cmd := cmd.NewListProjectsCommand(client)

	use := "projects"
	short := "List the projects. Default query is to list project for ${GITHUB_OWNER}"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

func TestListProjects(t *testing.T) {
	tt := []struct {
		name           string
		expectedOutput string
		expectedError  error
		listOpts       github.ProjectListOptions
		userArgs       cmd.ListProjectsOptions
	}{
		{
			name:           "projects for org",
			expectedOutput: "Project Al\nProject Betty\n",
			expectedError:  nil,
			listOpts: github.ProjectListOptions{
				State: "closed",
			},
			userArgs: cmd.ListProjectsOptions{
				Org:   "buena-vista-social-club",
				State: "closed",
			},
		},
		{
			name:           "projects for repo",
			expectedOutput: "Project Al\nProject Betty\n",
			expectedError:  nil,
			listOpts: github.ProjectListOptions{
				State: "open",
			},
			userArgs: cmd.ListProjectsOptions{
				Repo:  "hagen",
				State: "open",
			},
		},
		{
			name:           "state defaults to open if entered incorrectly",
			expectedOutput: "Project Al\nProject Betty\n",
			expectedError:  nil,
			listOpts: github.ProjectListOptions{
				State: "open",
			},
			userArgs: cmd.ListProjectsOptions{
				Repo:  "hagen",
				State: "micky-bubbles",
			},
		},
		{
			name:           "defaults to projects for org if no org or repo specified",
			expectedOutput: "Project Al\nProject Betty\n",
			expectedError:  nil,
			listOpts: github.ProjectListOptions{
				State: "open",
			},
			userArgs: cmd.ListProjectsOptions{
				State: "open",
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			viper.Set("GITHUB_OWNER", "buena-vista-social-club")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			client := pkg.NewMockProvider(ctrl)

			projectAlName := "Project Al"
			projectAl := github.Project{
				Name: &projectAlName,
			}
			projectBettyName := "Project Betty"
			projectBetty := github.Project{
				Name: &projectBettyName,
			}
			projects := []*github.Project{
				&projectAl, &projectBetty,
			}

			client.
				EXPECT().
				ListProjectsForOrg("buena-vista-social-club", gomock.Eq(tc.listOpts)).
				Return(projects, nil, nil).
				AnyTimes()

			client.
				EXPECT().
				ListProjectsForRepo("hagen", gomock.Eq(tc.listOpts)).
				Return(projects, nil, nil).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			err := cmd.ListProjects(client, tc.userArgs, writer)
			writer.Flush()

			if b.String() != tc.expectedOutput {
				t.Fatalf("expected '%s'; got '%s'", tc.expectedOutput, b.String())
			}

			if err != tc.expectedError {
				t.Fatalf("expected error '%v'; got '%v'", tc.expectedError, err)
			}
		})
	}
}

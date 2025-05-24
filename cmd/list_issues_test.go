package cmd_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/benmatselby/hagen/cmd"
	"github.com/benmatselby/hagen/pkg"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v72/github"
)

func TestNewListCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := pkg.NewMockProvider(ctrl)

	cmd := cmd.NewListIssuesCommand(client)

	use := "issues"
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

func TestDefaultIssueDisplayStrategy(t *testing.T) {
	issues := []*github.Issue{
		{
			URL:    github.Ptr("https://api.github.com/repos/foo/bar/issues/1"),
			Number: github.Ptr(1),
			Title:  github.Ptr("Test Issue 1"),
			Labels: []*github.Label{{Name: github.Ptr("bug")}},
		},
		{
			URL:    github.Ptr("https://api.github.com/repos/foo/bar/issues/2"),
			Number: github.Ptr(2),
			Title:  github.Ptr("Test Issue 2"),
			Labels: []*github.Label{},
		},
	}
	result := &github.IssuesSearchResult{Issues: issues}
	opts := cmd.ListIssuesOptions{DisplayLabels: true}
	var buf bytes.Buffer
	strategy := cmd.DefaultIssueDisplayStrategy{}
	err := strategy.Display(result, opts, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := buf.String()
	if !strings.Contains(output, "Test Issue 1") || !strings.Contains(output, "bug") {
		t.Errorf("expected output to contain issue title and label, got: %s", output)
	}
}

func TestTableIssueDisplayStrategy(t *testing.T) {
	created := github.Timestamp{Time: time.Date(2024, 5, 24, 10, 30, 0, 0, time.UTC)}
	closed := github.Timestamp{Time: time.Date(2024, 5, 25, 11, 0, 0, 0, time.UTC)}
	issues := []*github.Issue{
		{
			URL:       github.Ptr("https://api.github.com/repos/foo/bar/issues/1"),
			Number:    github.Ptr(1),
			Title:     github.Ptr("Test Issue 1"),
			Labels:    []*github.Label{{Name: github.Ptr("bug")}},
			CreatedAt: &created,
			State:     github.Ptr("open"),
			ClosedAt:  &closed,
		},
		{
			URL:              github.Ptr("https://api.github.com/repos/foo/bar/issues/2"),
			Number:           github.Ptr(2),
			Title:            github.Ptr("Test PR 2"),
			Labels:           []*github.Label{},
			CreatedAt:        &created,
			State:            github.Ptr("closed"),
			ClosedAt:         &closed,
			PullRequestLinks: &github.PullRequestLinks{}, // Mark as PR
		},
	}
	result := &github.IssuesSearchResult{Issues: issues}
	opts := cmd.ListIssuesOptions{DisplayLabels: true}
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	strategy := cmd.TableIssueDisplayStrategy{}
	err := strategy.Display(result, opts, writer)
	writer.Flush()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `┌──────────────┬────────────┬────────┬──────────────┬────────┬────────┬─────────────────────┬─────────────────────┐
│     TYPE     │ REPOSITORY │ NUMBER │    TITLE     │ LABELS │ STATUS │     CREATED AT      │      CLOSED AT      │
├──────────────┼────────────┼────────┼──────────────┼────────┼────────┼─────────────────────┼─────────────────────┤
│ Issue        │ foo/bar    │ 1      │ Test Issue 1 │ bug    │ open   │ 2024-05-24 10:30:00 │ 2024-05-25 11:00:00 │
│ Pull Request │ foo/bar    │ 2      │ Test PR 2    │        │ closed │ 2024-05-24 10:30:00 │ 2024-05-25 11:00:00 │
└──────────────┴────────────┴────────┴──────────────┴────────┴────────┴─────────────────────┴─────────────────────┘
`
	actual := b.String()
	if actual != expected {
		t.Fatalf("expected\n'%s'\ngot\n'%s'\n", expected, actual)
	}
}

func TestTableIssueDisplayStrategy_HumanDates(t *testing.T) {
	created := github.Timestamp{Time: time.Date(2025, 3, 9, 12, 33, 0, 0, time.UTC)}
	closed := github.Timestamp{Time: time.Date(2025, 3, 10, 14, 0, 0, 0, time.UTC)}
	issues := []*github.Issue{
		{
			URL:       github.Ptr("https://api.github.com/repos/foo/bar/issues/1"),
			Number:    github.Ptr(1),
			Title:     github.Ptr("Test Issue 1"),
			Labels:    []*github.Label{{Name: github.Ptr("bug")}},
			CreatedAt: &created,
			State:     github.Ptr("open"),
			ClosedAt:  &closed,
		},
		{
			URL:              github.Ptr("https://api.github.com/repos/foo/bar/issues/2"),
			Number:           github.Ptr(2),
			Title:            github.Ptr("Test PR 2"),
			Labels:           []*github.Label{},
			CreatedAt:        &created,
			State:            github.Ptr("closed"),
			ClosedAt:         &closed,
			PullRequestLinks: &github.PullRequestLinks{}, // Mark as PR
		},
	}
	result := &github.IssuesSearchResult{Issues: issues}
	opts := cmd.ListIssuesOptions{DisplayLabels: true, HumanDates: true}
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	strategy := cmd.TableIssueDisplayStrategy{}
	err := strategy.Display(result, opts, writer)
	writer.Flush()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `┌──────────────┬────────────┬────────┬──────────────┬────────┬────────┬────────────────────────────────┬────────────────────────────────┐
│     TYPE     │ REPOSITORY │ NUMBER │    TITLE     │ LABELS │ STATUS │           CREATED AT           │           CLOSED AT            │
├──────────────┼────────────┼────────┼──────────────┼────────┼────────┼────────────────────────────────┼────────────────────────────────┤
│ Issue        │ foo/bar    │ 1      │ Test Issue 1 │ bug    │ open   │ Sunday 09 March, 2025 at 12:33 │ Monday 10 March, 2025 at 14:00 │
│ Pull Request │ foo/bar    │ 2      │ Test PR 2    │        │ closed │ Sunday 09 March, 2025 at 12:33 │ Monday 10 March, 2025 at 14:00 │
└──────────────┴────────────┴────────┴──────────────┴────────┴────────┴────────────────────────────────┴────────────────────────────────┘
`
	actual := b.String()
	if actual != expected {
		t.Fatalf("expected\n'%s'\ngot\n'%s'\n", expected, actual)
	}
}

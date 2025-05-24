package cmd_test

import (
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
	issues := []*github.Issue{
		{
			URL:       github.Ptr("https://api.github.com/repos/foo/bar/issues/1"),
			Number:    github.Ptr(1),
			Title:     github.Ptr("Test Issue 1"),
			Labels:    []*github.Label{{Name: github.Ptr("bug")}},
			CreatedAt: &created,
		},
		{
			URL:              github.Ptr("https://api.github.com/repos/foo/bar/issues/2"),
			Number:           github.Ptr(2),
			Title:            github.Ptr("Test PR 2"),
			Labels:           []*github.Label{},
			CreatedAt:        &created,
			PullRequestLinks: &github.PullRequestLinks{}, // Mark as PR
		},
	}
	result := &github.IssuesSearchResult{Issues: issues}
	opts := cmd.ListIssuesOptions{DisplayLabels: true}
	var buf bytes.Buffer
	strategy := cmd.TableIssueDisplayStrategy{}
	err := strategy.Display(result, opts, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := buf.String()
	if !strings.Contains(strings.ToLower(output), "repository") ||
		!strings.Contains(output, "Test Issue 1") ||
		!strings.Contains(output, "bug") ||
		!strings.Contains(output, "Issue") ||
		!strings.Contains(output, "Pull Request") ||
		!strings.Contains(output, "2024-05-24 10:30:00") {
		t.Errorf("expected table output to contain headers, type, created at, issue title, and label, got: %s", output)
	}
}

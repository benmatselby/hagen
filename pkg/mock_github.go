// Code generated by MockGen. DO NOT EDIT.
// Source: provider.go

// Package pkg is a generated GoMock package.
package pkg

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/github"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// GetProjectByName mocks base method.
func (m *MockProvider) GetProjectByName(name, org, repo string) *github.Project {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectByName", name, org, repo)
	ret0, _ := ret[0].(*github.Project)
	return ret0
}

// GetProjectByName indicates an expected call of GetProjectByName.
func (mr *MockProviderMockRecorder) GetProjectByName(name, org, repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectByName", reflect.TypeOf((*MockProvider)(nil).GetProjectByName), name, org, repo)
}

// ListColumnsForProject mocks base method.
func (m *MockProvider) ListColumnsForProject(projectName, org, repo string) ([]*github.ProjectColumn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListColumnsForProject", projectName, org, repo)
	ret0, _ := ret[0].([]*github.ProjectColumn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListColumnsForProject indicates an expected call of ListColumnsForProject.
func (mr *MockProviderMockRecorder) ListColumnsForProject(projectName, org, repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListColumnsForProject", reflect.TypeOf((*MockProvider)(nil).ListColumnsForProject), projectName, org, repo)
}

// ListIssues mocks base method.
func (m *MockProvider) ListIssues(query string, opts github.SearchOptions) (*github.IssuesSearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIssues", query, opts)
	ret0, _ := ret[0].(*github.IssuesSearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIssues indicates an expected call of ListIssues.
func (mr *MockProviderMockRecorder) ListIssues(query, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssues", reflect.TypeOf((*MockProvider)(nil).ListIssues), query, opts)
}

// ListIssuesForProjectColumn mocks base method.
func (m *MockProvider) ListIssuesForProjectColumn(columnID int64) ([]*github.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIssuesForProjectColumn", columnID)
	ret0, _ := ret[0].([]*github.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIssuesForProjectColumn indicates an expected call of ListIssuesForProjectColumn.
func (mr *MockProviderMockRecorder) ListIssuesForProjectColumn(columnID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssuesForProjectColumn", reflect.TypeOf((*MockProvider)(nil).ListIssuesForProjectColumn), columnID)
}

// ListProjectsForOrg mocks base method.
func (m *MockProvider) ListProjectsForOrg(orgName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectsForOrg", orgName, opts)
	ret0, _ := ret[0].([]*github.Project)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectsForOrg indicates an expected call of ListProjectsForOrg.
func (mr *MockProviderMockRecorder) ListProjectsForOrg(orgName, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsForOrg", reflect.TypeOf((*MockProvider)(nil).ListProjectsForOrg), orgName, opts)
}

// ListProjectsForRepo mocks base method.
func (m *MockProvider) ListProjectsForRepo(repoName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectsForRepo", repoName, opts)
	ret0, _ := ret[0].([]*github.Project)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectsForRepo indicates an expected call of ListProjectsForRepo.
func (mr *MockProviderMockRecorder) ListProjectsForRepo(repoName, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsForRepo", reflect.TypeOf((*MockProvider)(nil).ListProjectsForRepo), repoName, opts)
}

// ListRepos mocks base method.
func (m *MockProvider) ListRepos(query string, opts github.SearchOptions) (*github.RepositoriesSearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRepos", query, opts)
	ret0, _ := ret[0].(*github.RepositoriesSearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRepos indicates an expected call of ListRepos.
func (mr *MockProviderMockRecorder) ListRepos(query, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRepos", reflect.TypeOf((*MockProvider)(nil).ListRepos), query, opts)
}

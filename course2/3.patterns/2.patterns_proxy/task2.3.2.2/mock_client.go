// Code generated by MockGen. DO NOT EDIT.
// Source: main.go
//
// Generated by this command:
//
//	mockgen -source=main.go -destination=mock_client.go -package=main
//

// Package main is a generated GoMock package.
package main

import (
	context "context"
	reflect "reflect"

	github "github.com/google/go-github/v53/github"
	gomock "go.uber.org/mock/gomock"
)

// MockRepoLister is a mock of RepoLister interface.
type MockRepoLister struct {
	ctrl     *gomock.Controller
	recorder *MockRepoListerMockRecorder
}

// MockRepoListerMockRecorder is the mock recorder for MockRepoLister.
type MockRepoListerMockRecorder struct {
	mock *MockRepoLister
}

// NewMockRepoLister creates a new mock instance.
func NewMockRepoLister(ctrl *gomock.Controller) *MockRepoLister {
	mock := &MockRepoLister{ctrl: ctrl}
	mock.recorder = &MockRepoListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepoLister) EXPECT() *MockRepoListerMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockRepoLister) List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, username, opt)
	ret0, _ := ret[0].([]*github.Repository)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockRepoListerMockRecorder) List(ctx, username, opt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepoLister)(nil).List), ctx, username, opt)
}

// MockGistLister is a mock of GistLister interface.
type MockGistLister struct {
	ctrl     *gomock.Controller
	recorder *MockGistListerMockRecorder
}

// MockGistListerMockRecorder is the mock recorder for MockGistLister.
type MockGistListerMockRecorder struct {
	mock *MockGistLister
}

// NewMockGistLister creates a new mock instance.
func NewMockGistLister(ctrl *gomock.Controller) *MockGistLister {
	mock := &MockGistLister{ctrl: ctrl}
	mock.recorder = &MockGistListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGistLister) EXPECT() *MockGistListerMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockGistLister) List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, username, opt)
	ret0, _ := ret[0].([]*github.Gist)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockGistListerMockRecorder) List(ctx, username, opt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGistLister)(nil).List), ctx, username, opt)
}

// MockGithuber is a mock of Githuber interface.
type MockGithuber struct {
	ctrl     *gomock.Controller
	recorder *MockGithuberMockRecorder
}

// MockGithuberMockRecorder is the mock recorder for MockGithuber.
type MockGithuberMockRecorder struct {
	mock *MockGithuber
}

// NewMockGithuber creates a new mock instance.
func NewMockGithuber(ctrl *gomock.Controller) *MockGithuber {
	mock := &MockGithuber{ctrl: ctrl}
	mock.recorder = &MockGithuberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGithuber) EXPECT() *MockGithuberMockRecorder {
	return m.recorder
}

// GetGists mocks base method.
func (m *MockGithuber) GetGists(ctx context.Context, username string) ([]Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGists", ctx, username)
	ret0, _ := ret[0].([]Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGists indicates an expected call of GetGists.
func (mr *MockGithuberMockRecorder) GetGists(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGists", reflect.TypeOf((*MockGithuber)(nil).GetGists), ctx, username)
}

// GetRepos mocks base method.
func (m *MockGithuber) GetRepos(ctx context.Context, username string) ([]Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepos", ctx, username)
	ret0, _ := ret[0].([]Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepos indicates an expected call of GetRepos.
func (mr *MockGithuberMockRecorder) GetRepos(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepos", reflect.TypeOf((*MockGithuber)(nil).GetRepos), ctx, username)
}

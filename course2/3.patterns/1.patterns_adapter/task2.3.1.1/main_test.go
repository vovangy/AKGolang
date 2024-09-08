package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/v53/github"
)

type MockGistLister struct {
	GistsFunc func(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

func (m *MockGistLister) List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
	if m.GistsFunc != nil {
		return m.GistsFunc(ctx, username, opt)
	}
	return nil, nil, nil
}

type MockRepoLister struct {
	RepoFunc func(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

func (m *MockRepoLister) List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	if m.RepoFunc != nil {
		return m.RepoFunc(ctx, username, opt)
	}
	return nil, nil, nil
}

func TestGetRepos_Success(t *testing.T) {
	mockRepoService := &MockRepoLister{
		RepoFunc: func(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
			return []*github.Repository{
				{
					Name:        github.String("Repo1"),
					Description: github.String("Test Repo 1"),
					CloneURL:    github.String("http://repo1"),
				},
				{
					Name:        github.String("Repo2"),
					Description: github.String("Test Repo 2"),
					CloneURL:    github.String("http://repo2"),
				},
			}, nil, nil
		},
	}
	adapter := GithubAdapter{RepoList: mockRepoService}
	ctx := context.Background()
	repos, err := adapter.GetRepos(ctx, "test")

	if err != nil {
		t.Fatalf("Unexpected error while retrieving repositories: %v", err)
	}

	if len(repos) != 2 {
		t.Fatalf("Expected 2 repositories, but got %d", len(repos))
	}
}

func TestGetGists_Success(t *testing.T) {
	mockGistService := &MockGistLister{
		GistsFunc: func(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
			return []*github.Gist{
				{
					ID:          github.String("1"),
					Description: github.String("Test Gist 1"),
					GitPullURL:  github.String("http://gist1"),
				},
				{
					ID:          github.String("2"),
					Description: github.String("Test Gist 2"),
					GitPullURL:  github.String("http://gist2"),
				},
			}, nil, nil
		},
	}
	adapter := GithubAdapter{GistList: mockGistService}
	ctx := context.Background()
	gists, err := adapter.GetGists(ctx, "test")

	if err != nil {
		t.Fatalf("Unexpected error while retrieving gists: %v", err)
	}

	if len(gists) != 2 {
		t.Fatalf("Expected 2 gists, but got %d", len(gists))
	}

}

func TestGetRepos_Error(t *testing.T) {
	mockRepoService := &MockRepoLister{
		RepoFunc: func(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
			return nil, nil, fmt.Errorf("an error occurred while fetching repositories")
		},
	}
	adapter := GithubAdapter{RepoList: mockRepoService}
	ctx := context.Background()
	_, err := adapter.GetRepos(ctx, "test")

	if err == nil {
		t.Fatalf("Expected an error when fetching repositories, but got none")
	}
}

func TestGetGists_Error(t *testing.T) {
	mockGistService := &MockGistLister{
		GistsFunc: func(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
			return nil, nil, fmt.Errorf("an error occurred while fetching gists")
		},
	}
	adapter := GithubAdapter{GistList: mockGistService}
	ctx := context.Background()
	_, err := adapter.GetGists(ctx, "test")

	if err == nil {
		t.Fatalf("Expected an error when fetching gists, but got none")
	}
}

package main

import (
	"context"
	"testing"

	"github.com/google/go-github/v53/github"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGithubAdapter_GetGists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGistLister := NewMockGistLister(ctrl)
	adapter := &GithubAdapter{
		GistList: mockGistLister,
	}

	gistID := "123"
	gistDesc := "Sample Gist"
	gistURL := "https://gist.github.com/123"

	// Setting expectations
	mockGistLister.EXPECT().
		List(gomock.Any(), "testuser", nil).
		Return([]*github.Gist{
			{
				ID:          &gistID,
				Description: &gistDesc,
				GitPullURL:  &gistURL,
			},
		}, nil, nil)

	ctx := context.Background()
	items, err := adapter.GetGists(ctx, "testuser")

	assert.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, gistID, items[0].Title)
	assert.Equal(t, gistDesc, items[0].Description)
	assert.Equal(t, gistURL, items[0].Link)
}

func TestGithubAdapter_GetRepos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepoLister := NewMockRepoLister(ctrl)
	adapter := &GithubAdapter{
		RepoList: mockRepoLister,
	}

	repoName := "sample-repo"
	repoDesc := "Sample Description"
	repoURL := "https://github.com/sample-repo"

	mockRepoLister.EXPECT().
		List(gomock.Any(), "testuser", nil).
		Return([]*github.Repository{
			{
				Name:        &repoName,
				Description: &repoDesc,
				CloneURL:    &repoURL,
			},
		}, nil, nil)

	ctx := context.Background()
	items, err := adapter.GetRepos(ctx, "testuser")

	assert.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, repoName, items[0].Title)
	assert.Equal(t, repoDesc, items[0].Description)
	assert.Equal(t, repoURL, items[0].Link)
}

func TestGithubProxy_GetRepos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGithuber := NewMockGithuber(ctrl)
	proxy := GithubProxy{
		github: mockGithuber,
		cache:  make(map[string][]Item),
	}

	repoName := "sample-repo"
	repoDesc := "Sample Description"
	repoURL := "https://github.com/sample-repo"

	mockGithuber.EXPECT().
		GetRepos(gomock.Any(), "testuser").
		Return([]Item{
			{Title: repoName, Description: repoDesc, Link: repoURL},
		}, nil).Times(1)

	ctx := context.Background()

	items, err := proxy.GetRepos(ctx, "testuser")
	assert.NoError(t, err)
	assert.Len(t, items, 1)

	cachedItems, err := proxy.GetRepos(ctx, "testuser")
	assert.NoError(t, err)
	assert.Len(t, cachedItems, 1)

	assert.Equal(t, items, cachedItems)
}

func TestGithubProxy_GetGists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGithuber := NewMockGithuber(ctrl)
	proxy := GithubProxy{
		github: mockGithuber,
		cache:  make(map[string][]Item),
	}

	gistTitle := "123"
	gistDesc := "Sample Gist"
	gistURL := "https://gist.github.com/123"

	mockGithuber.EXPECT().
		GetGists(gomock.Any(), "testuser").
		Return([]Item{
			{Title: gistTitle, Description: gistDesc, Link: gistURL},
		}, nil).Times(1)

	ctx := context.Background()

	items, err := proxy.GetGists(ctx, "testuser")
	assert.NoError(t, err)
	assert.Len(t, items, 1)

	cachedItems, err := proxy.GetGists(ctx, "testuser")
	assert.NoError(t, err)
	assert.Len(t, cachedItems, 1)

	assert.Equal(t, items, cachedItems)
}

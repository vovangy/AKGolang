package main

import (
	"context"
	"testing"

	"github.com/google/go-github/v53/github"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGithubRepo_GetItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepoLister := NewMockRepoLister(ctrl)
	repo := &GithubRepo{RepoList: mockRepoLister}

	ctx := context.Background()
	username := "testuser"

	mockRepos := []*github.Repository{
		{Name: github.String("repo1"), Description: github.String("Description 1"), CloneURL: github.String("https://github.com/repo1")},
		{Name: github.String("repo2"), Description: github.String("Description 2"), CloneURL: github.String("https://github.com/repo2")},
	}

	mockRepoLister.EXPECT().List(ctx, username, nil).Return(mockRepos, nil, nil)

	items, err := repo.GetItems(ctx, username)

	assert.NoError(t, err)
	assert.Len(t, items, 2)
	assert.Equal(t, "repo1", items[0].Title)
	assert.Equal(t, "Description 1", items[0].Description)
	assert.Equal(t, "https://github.com/repo1", items[0].Link)
}

func TestGithubGist_GetItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGistLister := NewMockGistLister(ctrl)
	gist := &GithubGist{GistList: mockGistLister}

	ctx := context.Background()
	username := "testuser"

	mockGists := []*github.Gist{
		{ID: github.String("gist1"), Description: github.String("Gist 1"), GitPullURL: github.String("https://gist.github.com/gist1")},
		{ID: github.String("gist2"), Description: github.String("Gist 2"), GitPullURL: github.String("https://gist.github.com/gist2")},
	}

	mockGistLister.EXPECT().List(ctx, username, nil).Return(mockGists, nil, nil)

	items, err := gist.GetItems(ctx, username)

	assert.NoError(t, err)
	assert.Len(t, items, 2)
	assert.Equal(t, "gist1", items[0].Title)
	assert.Equal(t, "Gist 1", items[0].Description)
	assert.Equal(t, "https://gist.github.com/gist1", items[0].Link)
}

func TestGeneralGithub_GetItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGithubLister := NewMockGithubLister(ctrl)
	generalGithub := &GeneralGithub{}

	ctx := context.Background()
	username := "testuser"

	mockItems := []Item{
		{Title: "item1", Description: "Item 1", Link: "https://github.com/item1"},
		{Title: "item2", Description: "Item 2", Link: "https://github.com/item2"},
	}

	mockGithubLister.EXPECT().GetItems(ctx, username).Return(mockItems, nil)

	items, err := generalGithub.GetItems(ctx, username, mockGithubLister)

	assert.NoError(t, err)
	assert.Len(t, items, 2)
	assert.Equal(t, "item1", items[0].Title)
	assert.Equal(t, "Item 1", items[0].Description)
	assert.Equal(t, "https://github.com/item1", items[0].Link)
}

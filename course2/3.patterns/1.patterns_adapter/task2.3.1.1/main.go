package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
)

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error)
}

type GithubAdapter struct {
	RepoList RepoLister
	GistList GistLister
}

type Item struct {
	Title       string
	Description string
	Link        string
}

func (g *GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	if g.RepoList == nil {
		return nil, fmt.Errorf("RepoList is not initialized")
	}

	repos, _, err := g.RepoList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	items := make([]Item, len(repos))
	for i, repo := range repos {
		items[i] = Item{
			Title:       repo.GetName(),
			Description: repo.GetDescription(),
			Link:        repo.GetHTMLURL(),
		}
	}
	return items, nil
}

func (g *GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	if g.GistList == nil {
		return nil, fmt.Errorf("GistList is not initialized")
	}

	gists, _, err := g.GistList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	items := make([]Item, len(gists))
	for i, gist := range gists {
		items[i] = Item{
			Title:       gist.GetID(),
			Description: gist.GetDescription(),
			Link:        gist.GetGitPullURL(),
		}
	}
	return items, nil
}

func NewGithub(client *github.Client) *GithubAdapter {
	return &GithubAdapter{
		RepoList: client.Repositories,
		GistList: client.Gists,
	}
}

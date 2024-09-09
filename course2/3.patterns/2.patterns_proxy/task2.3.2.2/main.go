package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
)

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) (
		[]*github.Repository, *github.Response, error)
}

type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist,
		*github.Response, error)
}

type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error)
}

type GithubAdapter struct {
	RepoList RepoLister
	GistList GistLister
}

func NewGithubAdapter(githubClient *github.Client) *GithubAdapter {
	g := &GithubAdapter{
		RepoList: githubClient.Repositories,
		GistList: githubClient.Gists,
	}
	return g
}

func (adapter *GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := adapter.GistList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	gsts := []Item{}
	for _, gist := range gists {
		gsts = append(gsts, Item{Title: *gist.ID, Description: *gist.Description, Link: *gist.GitPullURL})
	}

	return gsts, err
}

func (adapter *GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	repos, _, err := adapter.RepoList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	reps := []Item{}
	for _, repo := range repos {
		reps = append(reps, Item{Title: repo.GetName(), Description: repo.GetDescription(), Link: repo.GetCloneURL()})
	}

	return reps, err
}

type GithubProxy struct {
	github Githuber
	cache  map[string][]Item
}

func NewProxy(adapter *GithubAdapter) *GithubProxy {
	cache := make(map[string][]Item)
	proxy := GithubProxy{github: adapter, cache: cache}

	return &proxy
}

func (proxy *GithubProxy) GetRepos(ctx context.Context, username string) ([]Item, error) {
	reposKey := fmt.Sprintf("repos_of_%s", username)

	if value, ok := proxy.cache[reposKey]; ok {
		return value, nil
	}

	info, err := proxy.github.GetRepos(ctx, username)

	if err != nil {
		return nil, err
	}

	proxy.cache[reposKey] = info
	return info, nil
}

func (proxy *GithubProxy) GetGists(ctx context.Context, username string) ([]Item, error) {
	gistsKey := fmt.Sprintf("gists_of_%s", username)

	if value, ok := proxy.cache[gistsKey]; ok {
		return value, nil
	}

	info, err := proxy.github.GetGists(ctx, username)

	if err != nil {
		return nil, err
	}

	proxy.cache[gistsKey] = info
	return info, nil
}

type Item struct {
	Title       string
	Description string
	Link        string
}

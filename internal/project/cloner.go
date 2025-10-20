package project

import (
	"spinup/internal/config"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func TryRemote(path, url string, remote config.TemplatesRemote) (*git.Repository, error) {
	if remote.Secret == "none" || len(remote.Secret) == 0 {
		repo, err := git.PlainClone(path, false, &git.CloneOptions{
			URL: url,
		})
		return repo, err
	}

	repo, err := git.PlainClone(path, false, &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Username: remote.Name,
			Password: remote.Secret,
		},
	})
	return repo, err
}

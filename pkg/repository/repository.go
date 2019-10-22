// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repository

import (
	utils "github.com/onosproject/onos-docs/pkg/utils"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Repository data structure to represent a repository information
type Repository struct {
	path         string
	cloneOptions git.CloneOptions
	gitRepo      *git.Repository
	tagName      string
}

// Builder repository builder interface
type Builder interface {
	SetPath(string) Builder
	SetCloneOptions(git.CloneOptions) Builder
	SetGitRepo(*git.Repository) Builder
	SetTagName(string) Builder
	Build() Repository
}

// New Creates an instance of repository builder
func New() Builder {
	return &Repository{
		path:         "/tmp/test",
		cloneOptions: git.CloneOptions{},
		gitRepo:      &git.Repository{},
		tagName:      "master",
	}
}

// Build build a Repository instance
func (repo *Repository) Build() Repository {
	return Repository{
		path:         repo.path,
		cloneOptions: repo.cloneOptions,
		gitRepo:      repo.gitRepo,
		tagName:      repo.tagName,
	}

}

// SetPath set repo path
func (repo *Repository) SetPath(path string) Builder {
	repo.path = path
	return repo
}

// SetCloneOptions set git clone options
func (repo *Repository) SetCloneOptions(cloneOptions git.CloneOptions) Builder {
	repo.cloneOptions = cloneOptions
	return repo
}

// SetGitRepo set git repo
func (repo *Repository) SetGitRepo(gitRepo *git.Repository) Builder {
	repo.gitRepo = gitRepo
	return repo
}

// SetTagName set the git repo tag name
func (repo *Repository) SetTagName(tagName string) Builder {
	repo.tagName = tagName
	return repo
}

//Clone clones a repo based on a given url and a path
func (repo *Repository) Clone() error {
	utils.Info("git clone %s", repo.cloneOptions.URL)
	r, err := git.PlainClone(repo.path, false, &repo.cloneOptions)
	repo.SetGitRepo(r)
	return err
}

// GetTag get a repo tag reference based on a given tag name
func (repo *Repository) GetTag() (*plumbing.Reference, error) {
	_, err := repo.gitRepo.Worktree()
	if err != nil {
		return nil, err
	}

	tagRepo, err := repo.gitRepo.Tag(repo.tagName)
	if err != nil {
		return nil, err
	}

	return tagRepo, err
}

// CheckOutTag check out a repo based on the given tag name
func (repo *Repository) CheckOutTag() error {
	w, err := repo.gitRepo.Worktree()
	if err != nil {
		return err
	}

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName("refs/tags/" + repo.tagName),
	})

	return err

}

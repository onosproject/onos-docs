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

package build

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/src-d/go-git.v4/plumbing"

	"github.com/onosproject/onos-docs/pkg/common"

	"github.com/onosproject/onos-docs/pkg/repository"
	utils "github.com/onosproject/onos-docs/pkg/utils"
	"gopkg.in/src-d/go-git.v4"
)

func cloneRepo(path string, indexVersion int, indexRepo int, docsConfig utils.DocsYamlConfig) {
	var cloneOptions git.CloneOptions

	repo := docsConfig.Versions[indexVersion].Repos[indexRepo]
	if repo.TagName != "master" {
		cloneOptions = git.CloneOptions{
			URL:           repo.URL,
			Depth:         1,
			ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/tags/%s", repo.TagName)),
			SingleBranch:  true,
			Tags:          git.NoTags,
		}
	} else {
		cloneOptions = git.CloneOptions{
			URL:           repo.URL,
			ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", repo.TagName)),
			SingleBranch:  true,
		}

	}
	gitRepo := repository.New().
		SetCloneOptions(cloneOptions).
		SetTagName(repo.TagName).
		SetPath(path).
		Build()

	err := gitRepo.Clone()
	utils.CheckIfError(err)
	err = utils.RemoveCode(path)
	utils.CheckIfError(err)
}

func clone(wgClone *sync.WaitGroup, repoIndex int, versionIndex int, docsDir string, config utils.DocsYamlConfig) {
	defer wgClone.Done()
	repo := config.Versions[versionIndex].Repos[repoIndex]
	path := docsDir + repo.Name
	utils.CreateDir(path)
	cloneRepo(path, versionIndex, repoIndex, config)
}

// VersionHandler handle different versions of docs
func (db *DocsBuilderConfig) VersionHandler(config *utils.DocsConfig) {
	versions := config.GetDocsYamlConfig().Versions
	db.latestVersion = config.GetDocsYamlConfig().LatestVersion
	utils.CreateDir(common.SiteDirName)
	versionsArray := make([]string, len(versions))

	for index, val := range versions {
		versionsArray[index] = val.Ver
	}
	db.versions = versionsArray
	for versionIndex, val := range versions {
		db.tagName = val.Ver
		repos := val.Repos
		var wgClone sync.WaitGroup
		switch val.Ver {
		case db.latestVersion:
			for repoIndex := range repos {
				wgClone.Add(1)
				go clone(&wgClone, repoIndex, versionIndex, os.Args[2], config.GetDocsYamlConfig())
			}
			wgClone.Wait()
			db.build()

		default:
			docsDir := os.Args[2] + val.Ver + "/"
			for repoIndex := range repos {
				wgClone.Add(1)
				go clone(&wgClone, repoIndex, versionIndex, docsDir, config.GetDocsYamlConfig())
			}
			wgClone.Wait()
			db.build()

		}
	}

}

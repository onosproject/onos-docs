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
	"os"

	"github.com/onosproject/onos-docs/pkg/common"

	"github.com/onosproject/onos-docs/pkg/repository"
	utils "github.com/onosproject/onos-docs/pkg/utils"
	"gopkg.in/src-d/go-git.v4"
)

// VersionHandler handle different versions of docs
func (db *DocsBuilderConfig) VersionHandler(config *utils.DocsConfig) {
	versions := config.GetDocsYamlConfig().Versions
	db.LatestVersion = config.GetDocsYamlConfig().LatestVersion
	err := os.MkdirAll(common.SiteDirName, common.PermissionMode)
	utils.CheckIfError(err)
	versionsArray := make([]string, len(versions))

	for index, val := range versions {
		versionsArray[index] = val.Ver
	}
	db.versions = versionsArray

	for _, val := range versions {
		db.tagName = val.Ver
		switch val.Ver {
		case db.LatestVersion:
			repos := val.Repos
			for _, repo := range repos {
				path := os.Args[2] + repo.Name
				err := os.MkdirAll(path, common.PermissionMode)
				utils.CheckIfError(err)
				cloneOptions := git.CloneOptions{
					URL:   repo.URL,
					Depth: 1,
					Tags:  git.NoTags,
				}
				gitRepo := repository.New().
					SetCloneOptions(cloneOptions).
					SetTagName(repo.TagName).
					SetPath(path).
					Build()

				err = gitRepo.Clone()
				utils.CheckIfError(err)
				err = utils.RemoveContents(path)
				utils.CheckIfError(err)
			}
			db.build()
		default:
			repos := val.Repos
			docsDir := os.Args[2] + val.Ver + "/"
			for _, repo := range repos {
				path := docsDir + repo.Name
				if _, err := os.Stat(path); os.IsNotExist(err) {
					err = os.MkdirAll(path, common.PermissionMode)
					utils.CheckIfError(err)
				}

				cloneOptions := git.CloneOptions{
					URL:  repo.URL,
					Tags: git.AllTags,
				}
				gitRepo := repository.New().
					SetCloneOptions(cloneOptions).
					SetTagName(repo.TagName).
					SetPath(path).
					Build()

				err := gitRepo.Clone()
				utils.CheckIfError(err)
				err = gitRepo.CheckOutTag()
				utils.CheckIfError(err)
				err = utils.RemoveContents(path)
				utils.CheckIfError(err)
			}
			db.build()
		}
	}
}

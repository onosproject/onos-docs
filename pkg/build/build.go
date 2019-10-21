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
	"github.com/onosproject/onos-docs/pkg/manifest"
	"github.com/onosproject/onos-docs/pkg/menu"
	"github.com/onosproject/onos-docs/pkg/types"
	utils "github.com/onosproject/onos-docs/pkg/utils"
)

// build build docs website according to the given list of versions
func build(versions []string, tagName string) {
	manif, _ := manifest.Read(os.Args[3])
	var docsDir string

	manifestPath := common.MkdocsConfig
	var siteDir string
	if tagName == Latest.String() {
		docsDir = os.Args[2]
		siteDir = common.SiteDirName
	} else {
		docsDir = os.Args[2] + tagName + "/"
		siteDir = common.SiteDirName + tagName + ""
		utils.RunCommand("cp", "./content/README.md", docsDir)
		utils.RunCommand("cp", "-r", "./content/images", docsDir)
		utils.RunCommand("cp", "-r", "./content/developers", docsDir)
	}

	manif["docs_dir"] = docsDir
	err := manifest.Write(manifestPath, manif)
	utils.CheckIfError(err)
	menuConfig := types.MenuFiles{
		JsFile: os.Args[4],
	}
	menuContent := menu.GetTemplateContent(&menuConfig)

	versionsInfo := types.VersionsInformation{
		Current:     tagName,
		Latest:      Latest.String(),
		CurrentPath: docsDir,
	}

	err = menu.Build(versionsInfo, versions, menuContent)
	utils.CheckIfError(err)

	utils.RunCommand("mkdocs", "build", "--site-dir", siteDir, "-q")
	manif["docs_dir"] = "content"
	err = manifest.Write(manifestPath, manif)
	utils.CheckIfError(err)

}

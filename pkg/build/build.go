// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package build

import (
	"os"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/onosproject/onos-docs/pkg/common"
	"github.com/onosproject/onos-docs/pkg/manifest"
	"github.com/onosproject/onos-docs/pkg/menu"
	"github.com/onosproject/onos-docs/pkg/types"
	utils "github.com/onosproject/onos-docs/pkg/utils"
)

var log = logging.GetLogger("build")

// DocsBuilderConfig docs builder configuration information
type DocsBuilderConfig struct {
	latestVersion string
	versions      []string
	tagName       string
}

// build build docs website according to the given list of versions
func (db *DocsBuilderConfig) build() {
	manif, err := manifest.Read(os.Args[3])
	if err != nil {
		log.Info(err)
	}
	nav, _ := manifest.Read("./configs/nav/nav_" + db.tagName + ".yml")
	var docsDir string
	manifestPath := common.MkdocsConfigPath
	var siteDir string
	if db.tagName == db.latestVersion {
		docsDir = os.Args[2]
		siteDir = common.SiteDirName
	} else {
		docsDir = os.Args[2] + db.tagName + "/"
		siteDir = common.SiteDirName + db.tagName + ""
		utils.RunCommand("cp", "./content/README.md", docsDir)
		utils.RunCommand("cp", "-r", "./content/images", docsDir)
		utils.RunCommand("cp", "-r", "./content/styles", docsDir)
		utils.RunCommand("cp", "./content/.nojekyll", docsDir)
	}

	manif["docs_dir"] = docsDir
	manif["nav"] = nav["nav"]
	err = manifest.Write(manifestPath, manif)
	utils.CheckIfError(err)
	menuConfig := types.MenuFiles{}
	if os.Args[4] != "" && os.Args[5] != "" {
		menuConfig = types.MenuFiles{
			JsFile:  os.Args[4],
			CSSFile: os.Args[5],
		}
	}
	menuContent := menu.GetTemplateContent(&menuConfig)

	versionsInfo := types.VersionsInformation{
		Current:     db.tagName,
		Latest:      db.latestVersion,
		CurrentPath: docsDir,
	}

	err = menu.Build(versionsInfo, db.versions, menuContent)
	utils.CheckIfError(err)

	utils.RunCommand("mkdocs", "build", "--site-dir", siteDir, "-q")
	manif["docs_dir"] = "content"
	manif["nav"] = nil
	err = manifest.Write(manifestPath, manif)
	utils.CheckIfError(err)

}

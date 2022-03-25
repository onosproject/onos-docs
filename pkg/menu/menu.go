// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package menu

import (
	"io/ioutil"

	"github.com/onosproject/onos-docs/pkg/common"
	"github.com/onosproject/onos-docs/pkg/manifest"
	"github.com/onosproject/onos-docs/pkg/types"
	utils "github.com/onosproject/onos-docs/pkg/utils"
	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/pkg/errors"
)

var log = logging.GetLogger("menu")

// Content Content of menu files.
type Content struct {
	Js  []byte
	CSS []byte
}

// GetTemplateContent Gets menu template content.
func GetTemplateContent(menu *types.MenuFiles) Content {
	var content Content

	if menu.HasJsFile() {
		jsContent, err := getMenuFileContent(menu.JsFile, menu.JsURL)
		if err != nil {
			return Content{}
		}
		content.Js = jsContent
	}

	if menu.HasCSSFile() {
		cssContent, err := getMenuFileContent(menu.CSSFile, menu.CSSURL)
		if err != nil {
			return Content{}
		}
		content.CSS = cssContent
	}

	return content
}

func getMenuFileContent(f string, u string) ([]byte, error) {
	if len(f) > 0 {
		content, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get template menu file content")
		}
		return content, nil
	}

	content, err := utils.Download(u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to download menu template")
	}
	return content, nil
}

// Build the menu.
func Build(versionsInfo types.VersionsInformation, branches []string, menuContent Content) error {
	manifestFile := common.MkdocsConfigPath
	manif, err := manifest.Read(manifest.FileName)
	if err != nil {
		return err
	}

	manifestDocsDir := manifest.GetDocsDir(manif, manifestFile)

	log.Info("Using docs_dir from manifest: ", manifestDocsDir)

	manifestJsFilePath, err := writeJsFile(manifestDocsDir, menuContent, versionsInfo, branches)
	if err != nil {
		return err
	}

	manifestCSSFilePath, err := writeCSSFile(manifestDocsDir, menuContent)
	if err != nil {
		return err
	}

	editManifest(manif, manifestJsFilePath, manifestCSSFilePath)

	err = manifest.Write(manifestFile, manif)
	if err != nil {
		return errors.Wrap(err, "error when edit MkDocs manifest")
	}

	return nil
}

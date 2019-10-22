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

package menu

import (
	"io/ioutil"
	"log"

	"github.com/onosproject/onos-docs/pkg/common"

	"github.com/containous/structor/file"
	"github.com/onosproject/onos-docs/pkg/manifest"
	"github.com/onosproject/onos-docs/pkg/types"

	"github.com/pkg/errors"
)

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

	content, err := file.Download(u)
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

	log.Printf("Using docs_dir from manifest: %s", manifestDocsDir)

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

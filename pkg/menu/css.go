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
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const menuCSSFileName = "onos-menu.css"

func writeCSSFile(manifestDocsDir string, menuContent Content) (string, error) {
	if len(menuContent.CSS) == 0 {
		return "", nil
	}

	cssDir := filepath.Join(manifestDocsDir, "theme", "css")
	if _, errStat := os.Stat(cssDir); os.IsNotExist(errStat) {
		errDir := os.MkdirAll(cssDir, os.ModePerm)
		if errDir != nil {
			return "", errors.Wrap(errDir, "error when create CSS folder")
		}
	}

	err := ioutil.WriteFile(filepath.Join(cssDir, menuCSSFileName), menuContent.CSS, os.ModePerm)
	if err != nil {
		return "", errors.Wrap(err, "error when trying ro write CSS file")
	}

	return filepath.Join("theme", "css", menuCSSFileName), nil
}

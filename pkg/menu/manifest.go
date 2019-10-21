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
	"github.com/onosproject/onos-docs/pkg/manifest"
)

func editManifest(manif map[string]interface{}, versionJsFile string, versionCSSFile string) {
	// Append menu JS file
	manifest.AppendExtraJs(manif, versionJsFile)

	// Append menu CSS file
	manifest.AppendExtraCSS(manif, versionCSSFile)

	// reset site URL
	manif["site_url"] = ""
}

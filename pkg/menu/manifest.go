// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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

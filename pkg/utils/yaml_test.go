// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package uitls

import (
	"testing"

	"gotest.tools/assert"
)

func TestYamlParser(t *testing.T) {

	config := NewDocsConfig("../../configs/versions.yml")

	err := config.Parse()
	assert.Equal(t, err, nil)
}

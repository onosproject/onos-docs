// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"github.com/onosproject/onos-docs/pkg/build"
	"github.com/onosproject/onos-lib-go/pkg/logging"

	utils "github.com/onosproject/onos-docs/pkg/utils"
)

var log = logging.GetLogger("main")

func main() {
	config := utils.NewDocsConfig(os.Args[1])
	err := config.Parse()
	log.Info("main error", err)
	utils.CheckIfError(err)

	var db build.DocsBuilderConfig

	log.Info(os.Args)
	db.VersionHandler(config)

}

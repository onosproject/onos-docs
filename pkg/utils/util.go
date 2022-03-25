// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package uitls

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/onosproject/onos-docs/pkg/common"

	"github.com/pkg/errors"
)

var log = logging.GetLogger("util")

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an example.
func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		log.Warn("Usage:", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}
	log.Error(err)
	os.Exit(1)
}

// RemoveCode remove all folders of a given dir except docs folder
func RemoveCode(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		if name != "docs" {
			err = os.RemoveAll(filepath.Join(dir, name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// RunCommand run a linux command
func RunCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	CheckIfError(err)
}

// Download Downloads a file.
func Download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return nil, errors.Errorf("failed to download %q: %s", url, resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

// CreateDir creates a dir based on a given path
func CreateDir(path string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		for {
			err := os.MkdirAll(path, common.PermissionMode)
			if err == nil {
				break
			}
		}

	}
}

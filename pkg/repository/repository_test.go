// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package repository

import (
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/src-d/go-git.v4/plumbing"

	"gopkg.in/src-d/go-git.v4"

	"gotest.tools/assert"
)

func TestClone(t *testing.T) {
	dir, err := ioutil.TempDir("/tmp", "test")
	if err != nil {
		panic(err)
	}
	rep := Repository{

		path: dir,
		cloneOptions: git.CloneOptions{
			URL:      "https://github.com/onosproject/onos-config",
			Tags:     git.AllTags,
			Progress: os.Stdout,
		},
	}
	err = rep.Clone()
	assert.Equal(t, err, nil)
	err = os.RemoveAll(dir)
	assert.Equal(t, err, nil)

}

func TestGetTags(t *testing.T) {
	dir, err := ioutil.TempDir("/tmp", "test")
	if err != nil {
		panic(err)
	}
	rep := Repository{
		path: dir,
		cloneOptions: git.CloneOptions{
			URL:      "https://github.com/onosproject/onos-config",
			Tags:     git.AllTags,
			Progress: os.Stdout,
		},
		tagName: "v0.1-onfconnect",
	}
	err = rep.Clone()
	assert.Equal(t, err, nil)
	_, err = rep.GetTag()
	assert.Equal(t, err, nil)
	err = os.RemoveAll(dir)
	assert.Equal(t, err, nil)
}

func TestCheckOut(t *testing.T) {
	dir, err := ioutil.TempDir("/tmp", "test")
	if err != nil {
		panic(err)
	}
	rep := Repository{
		path: dir,
		cloneOptions: git.CloneOptions{
			URL:      "https://github.com/onosproject/onos-config",
			Tags:     git.AllTags,
			Progress: os.Stdout,
		},
		tagName: "v0.1-onfconnect",
	}
	err = rep.Clone()
	assert.Equal(t, err, nil)
	err = rep.CheckOutTag()
	assert.Equal(t, err, nil)
	err = os.RemoveAll(dir)
	assert.Equal(t, err, nil)
}

// TestCloneInMemory test clone a repo in a memory storage
func TestCloneInMemory(t *testing.T) {
	rep := Repository{
		cloneOptions: git.CloneOptions{
			URL:           "https://github.com/onosproject/onos-docs",
			ReferenceName: plumbing.ReferenceName("refs/heads/gh-pages"),
			SingleBranch:  true,
		},
	}
	err := rep.CloneInMemory()
	assert.Equal(t, err, nil)
}

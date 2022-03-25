// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package menu

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/onosproject/onos-docs/pkg/types"

	"github.com/Masterminds/sprig"
	"github.com/hashicorp/go-version"
	"github.com/pkg/errors"
)

const menuJsFileName = "onos-menu.js"

const (
	stateLatest          = "LATEST"
	stateExperimental    = "EXPERIMENTAL"
	statePreFinalRelease = "PRE_FINAL_RELEASE"
)

type optionVersion struct {
	Path     string
	Text     string
	Name     string
	State    string
	Selected bool
}

func writeJsFile(manifestDocsDir string, menuContent Content, versionsInfo types.VersionsInformation, branches []string) (string, error) {
	if len(menuContent.Js) == 0 {
		return "", nil
	}

	jsDir := filepath.Join(manifestDocsDir, "theme", "js")
	if _, errStat := os.Stat(jsDir); os.IsNotExist(errStat) {
		errDir := os.MkdirAll(jsDir, os.ModePerm)
		if errDir != nil {
			return "", errors.Wrap(errDir, "error when create JS folder")
		}
	}

	menuFilePath := filepath.Join(jsDir, menuJsFileName)
	errBuild := buildJSFile(menuFilePath, versionsInfo, branches, string(menuContent.Js))
	if errBuild != nil {
		return "", errBuild
	}

	return filepath.Join("theme", "js", menuJsFileName), nil
}

func buildJSFile(filePath string, versionsInfo types.VersionsInformation, branches []string, menuTemplate string) error {
	temp := template.New("menu-js").Funcs(sprig.TxtFuncMap())

	_, err := temp.Parse(menuTemplate)
	if err != nil {
		return errors.Wrap(err, "error during parsing template")
	}

	versions, err := buildVersions(versionsInfo.Current, branches, versionsInfo.Latest, versionsInfo.Experimental)
	if err != nil {
		return errors.Wrap(err, "error when build versions")
	}

	model := struct {
		Latest   string
		Current  string
		Versions []optionVersion
	}{
		Latest:   versionsInfo.Latest,
		Current:  versionsInfo.Current,
		Versions: versions,
	}

	f, err := os.Create(filePath)
	if err != nil {
		return errors.Wrap(err, "error when create menu file")
	}

	return temp.Execute(f, model)
}

func buildVersions(currentVersion string, branches []string, latestTagName string, experimentalBranchName string) ([]optionVersion, error) {
	latestVersion, err := version.NewVersion(latestTagName)
	if err != nil {
		return nil, err
	}

	var versions []optionVersion
	for _, versionName := range branches {
		selected := currentVersion == versionName

		switch versionName {
		case latestTagName:
			// skip, because we must the branch instead of the tag
			versions = append(versions, optionVersion{
				Path:     "",
				Text:     "latest",
				Name:     latestTagName,
				State:    stateLatest,
				Selected: selected,
			})
		case experimentalBranchName:
			versions = append(versions, optionVersion{
				Path:     experimentalBranchName,
				Text:     "experimental",
				Name:     experimentalBranchName,
				State:    stateExperimental,
				Selected: selected,
			})
		default:
			simpleVersion, err := version.NewVersion(versionName)
			if err != nil {
				return nil, err
			}

			v := optionVersion{
				Name:     versionName,
				Selected: selected,
			}

			switch {
			case simpleVersion.GreaterThan(latestVersion):
				v.Path = versionName
				v.Text = versionName + " RC"
				v.State = statePreFinalRelease
			case sameMinor(simpleVersion, latestVersion):
				// latest version
				v.Text = versionName + " Latest"
				v.State = stateLatest
			default:
				v.Path = versionName
				v.Text = versionName
			}
			versions = append(versions, v)
		}
	}

	return versions, nil
}

func sameMinor(v1 *version.Version, v2 *version.Version) bool {
	v1Parts := v1.Segments()
	v2Parts := v2.Segments()

	return v1Parts[0] == v2Parts[0] && v1Parts[1] == v2Parts[1]
}

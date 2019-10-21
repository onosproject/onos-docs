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

package uitls

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// DocsConfig config data structure for config yaml file parser
type DocsConfig struct {
	docsYamlConfig DocsYamlConfig
	fileName       string
}

// NewDocsConfig returns an instance of DocsConfig data structure
func NewDocsConfig(fileName string) *DocsConfig {
	return &DocsConfig{
		fileName: fileName,
	}
}

// DocsYamlConfig corresponding golang structure for a versions yaml file
type DocsYamlConfig struct {
	Versions []struct {
		Ver   string `yaml:"ver"`
		Repos []struct {
			Name    string `yaml:"name"`
			URL     string `yaml:"url"`
			TagName string `yaml:"tagName"`
		} `yaml:"repos"`
	} `yaml:"versions"`
}

// GetDocsYamlConfig return the docs yaml config data structure
func (c *DocsConfig) GetDocsYamlConfig() DocsYamlConfig {
	return c.docsYamlConfig
}

// Parse parse a yaml config file
func (c *DocsConfig) Parse() error {
	data, err := ioutil.ReadFile(c.fileName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &c.docsYamlConfig)
	return err
}

// Copyright 2017 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://studygolang.com
// Author: polaris	polaris@studygolang.com

package cfg

import (
	"io/ioutil"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// YamlConfig represents a yaml configure.
// The root must be map.
type YamlConfig struct {
	file string
	data map[string]interface{}
}

// ParseYaml parses yaml file into YamlConfig, return error (not nil) if failure.
func ParseYaml(file string) (*YamlConfig, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "parse yaml, ReadFile error")
	}

	yamlConfig := &YamlConfig{
		file: file,
		data: make(map[string]interface{}),
	}
	if err = yaml.Unmarshal(buf, yamlConfig.data); err != nil {
		return nil, errors.Wrap(err, "parse yaml, unmarshal error")
	}

	return yamlConfig, nil
}

// Value returns the value of the key.
// It returns an error and empty string value if the key does not exist.
func (this *YamlConfig) Value(key string) (string, error) {
	if interVal, ok := this.data[key]; ok {
		if val, ok := interVal.(string); ok {
			return val, nil
		}

		return "", errors.Errorf("The value of %q is not string", key)
	}

	return "", errors.Errorf("The key: %q is not exists!", key)
}

// Int returns int type value.
// It returns an error and 0 if the key does not exist.
func (this *YamlConfig) Int(key string) (int, error) {
	if interVal, ok := this.data[key]; ok {
		if val, ok := interVal.(int); ok {
			return val, nil
		}

		return 0, errors.Errorf("The value of %q is not int", key)
	}

	return 0, errors.Errorf("The key: %q is not exists!", key)
}

// Float64 returns float64 type value.
// It returns an error and 0.0 if the key does not exist.
func (this *YamlConfig) Float64(key string) (float64, error) {
	if interVal, ok := this.data[key]; ok {
		if val, ok := interVal.(float64); ok {
			return val, nil
		}

		return 0.0, errors.Errorf("The value of %q is not float64", key)
	}

	return 0.0, errors.Errorf("The key: %q is not exists!", key)
}

// Bool returns bool type value.
// It returns an error and false if the key does not exist.
func (this *YamlConfig) Bool(key string) (bool, error) {
	if interVal, ok := this.data[key]; ok {
		if val, ok := interVal.(bool); ok {
			return val, nil
		}

		return false, errors.Errorf("The value of %q is not bool", key)
	}

	return false, errors.Errorf("The key: %q is not exists!", key)
}

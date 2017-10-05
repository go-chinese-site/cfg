// Copyright 2017 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://studygolang.com
// Author: polaris	polaris@studygolang.com

package cfg

import (
	"io/ioutil"
	"strings"

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
	if !this.Exists(key) {
		return "", errors.Errorf("The key: %q is not exists!", key)
	}

	if val, ok := this.data[key].(string); ok {
		return val, nil
	}

	return "", errors.Errorf("The value of %q is not string", key)
}

// Int returns int type value.
// It returns an error and 0 if the key does not exist.
func (this *YamlConfig) Int(key string) (int, error) {
	if !this.Exists(key) {
		return 0, errors.Errorf("The key: %q is not exists!", key)
	}

	if val, ok := this.data[key].(int); ok {
		return val, nil
	}

	return 0, errors.Errorf("The value of %q is not int", key)
}

// Float64 returns float64 type value.
// It returns an error and 0.0 if the key does not exist.
func (this *YamlConfig) Float64(key string) (float64, error) {
	if !this.Exists(key) {
		return 0.0, errors.Errorf("The key: %q is not exists!", key)
	}

	if val, ok := this.data[key].(float64); ok {
		return val, nil
	}

	return 0.0, errors.Errorf("The value of %q is not float64", key)

}

// Bool returns bool type value.
// It returns an error and false if the key does not exist.
func (this *YamlConfig) Bool(key string) (bool, error) {
	if !this.Exists(key) {
		return false, errors.Errorf("The key: %q is not exists!", key)
	}

	if val, ok := this.data[key].(bool); ok {
		return val, nil
	}

	return false, errors.Errorf("The value of %q is not bool", key)
}

// Map returns bool type value.
// It returns an error and false if the key does not exist.
func (this *YamlConfig) Map(key string) (map[interface{}]interface{}, error) {
	if !this.Exists(key) {
		return nil, errors.Errorf("The key: %q is not exists!", key)
	}

	if val, ok := this.data[key].(map[interface{}]interface{}); ok {
		return val, nil
	}

	return nil, errors.Errorf("The value of %q is not map[interface{}]interface{}", key)
}

// MustValue always returns value without error. It returns empty string if error occurs, or the default value if given.
func (this *YamlConfig) MustValue(key string, defaultVal ...string) string {
	val, err := this.Value(key)
	if len(defaultVal) > 0 && err != nil {
		return defaultVal[0]
	}
	return val
}

// Get returns bool type value.
// It returns an error and false if the key does not exist.
func (this *YamlConfig) Get(path string) *Result {
	pathes := strings.Split(path, ".")

	var result = &Result{}
	for _, key := range pathes {
		switch result.Type {
		case NotExists:
			if !this.Exists(key) {
				return result
			}

			result.Type = Interface
			result.Raw = this.data[key]
		case Interface:
			if isInt(key) {
				inters, ok := result.Raw.([]interface{})
				if !ok {
					return &Result{Type: NotExists}
				}

				index := mustInt(key)
				if index >= len(inters) {
					return &Result{Type: NotExists}
				}

				result.Raw = inters[index]
			} else {
				interMap, ok := result.Raw.(map[interface{}]interface{})
				if !ok {
					return &Result{Type: NotExists}
				}

				val, ok := interMap[key]
				if !ok {
					return &Result{Type: NotExists}
				}

				result.Raw = val
			}
		}
	}

	return result
}

// Exists returns true if key exists.
func (this *YamlConfig) Exists(key string) bool {
	if _, ok := this.data[key]; ok {
		return true
	}

	return false
}

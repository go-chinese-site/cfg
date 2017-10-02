// Copyright 2017 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://studygolang.com
// Author: polaris	polaris@studygolang.com

package cfg_test

import (
	"testing"

	"github.com/go-chinese-site/cfg"
)

func setUp(t *testing.T) *cfg.YamlConfig {
	yamlConfig, err := cfg.ParseYaml("testdata/cfg.yml")
	if err != nil {
		t.Error(err)
	}

	return yamlConfig
}

func TestYamlValue(t *testing.T) {
	yamlConfig := setUp(t)

	var intTests = []struct {
		key      string
		expected string
	}{
		{"host", "127.0.0.1"},
	}

	for _, tt := range intTests {
		actual, err := yamlConfig.Value(tt.key)
		if err != nil {
			t.Errorf("YamlConfig.Value(%s) had error: %v; expected %s", tt.key, err, tt.expected)
		}
		if actual != tt.expected {
			t.Errorf("YamlConfig.Value(%s) = %s; expected %s", tt.key, actual, tt.expected)
		}
	}
}

func TestYamlInt(t *testing.T) {
	yamlConfig := setUp(t)

	var intTests = []struct {
		key      string
		expected int
	}{
		{"port", 6061},
	}

	for _, tt := range intTests {
		actual, err := yamlConfig.Int(tt.key)
		if err != nil {
			t.Errorf("YamlConfig.Int(%s) had error: %v; expected %d", tt.key, err, tt.expected)
		}
		if actual != tt.expected {
			t.Errorf("YamlConfig.Int(%s) = %d; expected %d", tt.key, actual, tt.expected)
		}
	}
}

func TestYamlFloat64(t *testing.T) {
	yamlConfig := setUp(t)

	var intTests = []struct {
		key      string
		expected float64
	}{
		{"price", 12.3},
	}

	for _, tt := range intTests {
		actual, err := yamlConfig.Float64(tt.key)
		if err != nil {
			t.Errorf("YamlConfig.Float64(%s) had error: %v; expected %f", tt.key, err, tt.expected)
		}
		if actual != tt.expected {
			t.Errorf("YamlConfig.Float64(%s) = %f; expected %f", tt.key, actual, tt.expected)
		}
	}
}

func TestYamlBool(t *testing.T) {
	yamlConfig := setUp(t)

	var intTests = []struct {
		key      string
		expected bool
	}{
		{"isSet", true},
	}

	for _, tt := range intTests {
		actual, err := yamlConfig.Bool(tt.key)
		if err != nil {
			t.Errorf("YamlConfig.Bool(%s) had error: %v; expected %t", tt.key, err, tt.expected)
		}
		if actual != tt.expected {
			t.Errorf("YamlConfig.Bool(%s) = %t; expected %t", tt.key, actual, tt.expected)
		}
	}
}

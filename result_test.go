// Copyright 2017 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://studygolang.com
// Author: polaris  polaris@studygolang.com

package cfg_test

import (
	"testing"

	"github.com/go-chinese-site/cfg"
)

func TestResultString(t *testing.T) {
	tests := []struct {
		result   cfg.Result
		expected string
	}{
		{
			cfg.Result{Type: cfg.NotExists},
			"",
		},
		{
			cfg.Result{Type: cfg.Null},
			"",
		},
		{
			cfg.Result{Type: cfg.Number, Num: 10},
			"10",
		},
		{
			cfg.Result{Type: cfg.Number, Num: 10.2},
			"10.2",
		},
		{
			cfg.Result{Type: cfg.String, Str: "studygolang.com"},
			"studygolang.com",
		},
		{
			cfg.Result{Type: cfg.Interface, Raw: 10},
			"10",
		},
		{
			cfg.Result{Type: cfg.Interface, Raw: 10.2},
			"10.2",
		},
	}

	for _, tt := range tests {
		actual := tt.result.String()
		if actual != tt.expected {
			t.Errorf("The String of the Result: %#v is %s, expected: %s", tt.result, actual, tt.expected)
		}
	}
}

func TestResultInt(t *testing.T) {
	tests := []struct {
		result   cfg.Result
		expected int
	}{
		{
			cfg.Result{Type: cfg.NotExists},
			0,
		},
		{
			cfg.Result{Type: cfg.Null},
			0,
		},
		{
			cfg.Result{Type: cfg.Number, Num: 10},
			10,
		},
		{
			cfg.Result{Type: cfg.Interface, Raw: 10},
			10,
		},
	}

	for _, tt := range tests {
		actual := tt.result.Int()
		if actual != tt.expected {
			t.Errorf("The Int of the Result: %#v is %d, expected: %d", tt.result, actual, tt.expected)
		}
	}
}

// Copyright 2017 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://studygolang.com
// Author: polaris  polaris@studygolang.com

package cfg

import "strconv"

// Result represents a value that is returned from Get().
type Result struct {
	// Type is the config type
	Type Type
	// Str is the config string
	Str string
	// Num is the config number
	Num float64
	// Raw is unknow type value
	Raw interface{}
}

// Exists returns true if value exists.
func (this *Result) Exists() bool {
	return this.Type != NotExists
}

// String returns a string representation of the value. It returns empty string if error occurs, or the default value if given.
func (this *Result) String(defaultVal ...string) string {
	switch this.Type {
	default:
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return ""
	case False:
		return "false"
	case True:
		return "true"
	case Number:
		return strconv.FormatFloat(this.Num, 'f', -1, 64)
	case String:
		return this.Str
	case Interface:
		return mustString(this.Raw)
	}
}

// Int returns a int representation of the value. It returns 0 if error occurs, or the default value if given.
func (this *Result) Int(defaultVal ...int) int {
	switch this.Type {
	default:
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	case False:
		return 0
	case True:
		return 1
	case Number:
		// try to directly convert the float64 to int64
		n, ok := floatToInt(this.Num)
		if !ok {
			return 0
		}
		return int(n)
	case String:
		return mustInt(this.Str)
	case Interface:
		return mustInt(mustString(this.Raw))
	}
}

// Value returns one of these types:
/*
   bool, for config booleans
   float64, for config numbers
   Number, for config numbers
   string, for config string literals
   nil, for config null
*/
func (this *Result) Value() interface{} {
	switch this.Type {
	default:
		return ""
	case Null:
		return nil
	case False:
		return false
	case True:
		return true
	case Number:
		return this.Num
	case String:
		return this.Str
	case Interface:
		return this.Raw
	}
}

const (
	// NotExists is not exists in config
	NotExists Type = iota
	// Null is a null config value
	Null
	// False is a config false boolean
	False
	// True is a config true boolean
	True
	// Number is config number
	Number
	// String is a config string
	String
	// Interface is the unknow value
	Interface
)

// Type is Result type
type Type int

// String returns a string representation of the type.
func (t Type) String() string {
	switch t {
	default:
		return ""
	case NotExists:
		return "Not Exists"
	case Null:
		return "Null"
	case False:
		return "False"
	case True:
		return "True"
	case Number:
		return "Number"
	case String:
		return "String"
	case Interface:
		return "Interface"
	}
}

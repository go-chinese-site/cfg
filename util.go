// Copyright 2017 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://studygolang.com
// Author: polaris	polaris@studygolang.com

package cfg

import (
	"fmt"
	"strconv"
)

const minInt53 = -2251799813685248
const maxInt53 = 2251799813685247

func floatToInt(f float64) (n int64, ok bool) {
	n = int64(f)
	if float64(n) == f && n >= minInt53 && n <= maxInt53 {
		return n, true
	}
	return 0, false
}

func isInt(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}

	return true
}

func mustInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func mustString(inter interface{}) string {
	val, ok := inter.(string)
	if ok {
		return val
	}

	return fmt.Sprintf("%v", inter)
}

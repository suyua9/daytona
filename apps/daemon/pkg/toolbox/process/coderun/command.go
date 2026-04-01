// Copyright 2025 Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package coderun

import (
	"strconv"
	"strings"
)

func shellQuote(value string) string {
	return strconv.Quote(value)
}

func formatArgv(argv []string) string {
	if len(argv) == 0 {
		return ""
	}

	parts := make([]string, 0, len(argv))
	for _, arg := range argv {
		parts = append(parts, shellQuote(arg))
	}

	return strings.Join(parts, " ")
}

// SPDX-FileCopyrightText: Copyright 2021 The avo-sandbox Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/klauspost/asmfmt/cmd/asmfmt"
	_ "mvdan.cc/gofumpt"
)

//go:build tools
// +build tools

package tools

import (
	// Include Air Tool
	_ "github.com/air-verse/air"
	// Include Gum Tool
	_ "github.com/charmbracelet/gum"
	// Include Task Runner Tool
	_ "github.com/go-task/task/v3/cmd/task"
	// Include Golang Lint Tool
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	// Include PProf Tool
	_ "github.com/google/pprof"
	// Include Revive Tool
	_ "github.com/mgechev/revive"
	// Include Coverage View Tool
	_ "github.com/orlangure/gocovsh"
	// Include Markdown Documentation Tool
	_ "github.com/princjef/gomarkdoc/cmd/gomarkdoc"
	// Include Lines Formatter Tool
	_ "github.com/segmentio/golines"
	// Include SQLC Tool
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
	// Include Test Sum Tool
	_ "gotest.tools/gotestsum"
	// Include Static Check Tool
	_ "honnef.co/go/tools/cmd/staticcheck"
)

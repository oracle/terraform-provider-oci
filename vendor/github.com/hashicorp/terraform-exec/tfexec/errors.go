package tfexec

import (
	"fmt"
	"regexp"
)

func parseError(stderr string) error {
	switch {
	// case ErrTerraformNotFound.regexp.MatchString(stderr):
	// return ErrTerraformNotFound
	case regexp.MustCompile(usageRegexp).MatchString(stderr):
		return &ErrCLIUsage{stderr: stderr}
	case regexp.MustCompile(`Error: Could not satisfy plugin requirements`).MatchString(stderr):
		return &ErrNoInit{stderr: stderr}
	case regexp.MustCompile(`Error: No configuration files`).MatchString(stderr):
		return &ErrNoConfig{stderr: stderr}
	default:
		return &Err{stderr: stderr}
	}
}

type ErrNoSuitableBinary struct {
	err error
}

func (e *ErrNoSuitableBinary) Error() string {
	return fmt.Sprintf("no suitable terraform binary could be found: %s", e.err.Error())
}

// Not yet implemented.
// Intended for use when the detected Terraform version is not compatible with the command or flags being used in this invocation.
type ErrVersionMismatch struct{}

type ErrNoInit struct {
	stderr string
}

func (e *ErrNoInit) Error() string {
	return e.stderr
}

type ErrNoConfig struct {
	stderr string
}

func (e *ErrNoConfig) Error() string {
	return e.stderr
}

// Terraform CLI indicates usage errors in three different ways: either
// 1. Exit 1, with a custom error message on stderr.
// 2. Exit 1, with command usage logged to stderr.
// 3. Exit 127, with command usage logged to stdout.
// Currently cases 1 and 2 are handled.
// TODO KEM: Handle exit 127 case. How does this work on non-Unix platforms?
type ErrCLIUsage struct {
	stderr string
}

var usageRegexp = `Too many command line arguments|^Usage: .*Options:.*|Error: Invalid -\d+ option`

func (e *ErrCLIUsage) Error() string {
	return e.stderr
}

// catchall error
type Err struct {
	stderr string
}

func (e *Err) Error() string {
	return e.stderr
}

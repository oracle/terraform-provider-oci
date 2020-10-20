package tfexec

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

func parseError(err error, stderr string) error {
	if _, ok := err.(*exec.ExitError); !ok {
		return err
	}

	switch {
	// case ErrTerraformNotFound.regexp.MatchString(stderr):
	// return ErrTerraformNotFound
	case regexp.MustCompile(usageRegexp).MatchString(stderr):
		return &ErrCLIUsage{stderr: stderr}
	case regexp.MustCompile(`Error: Could not satisfy plugin requirements`).MatchString(stderr):
		return &ErrNoInit{stderr: stderr}
	case regexp.MustCompile(`Error: Could not load plugin`).MatchString(stderr):
		// this string is present in 0.13
		return &ErrNoInit{stderr: stderr}
	case regexp.MustCompile(`Error: No configuration files`).MatchString(stderr):
		return &ErrNoConfig{stderr: stderr}
	default:
		return errors.New(stderr)
	}
}

type ErrNoSuitableBinary struct {
	err error
}

func (e *ErrNoSuitableBinary) Error() string {
	return fmt.Sprintf("no suitable terraform binary could be found: %s", e.err.Error())
}

// ErrVersionMismatch is returned when the detected Terraform version is not compatible with the
// command or flags being used in this invocation.
type ErrVersionMismatch struct {
	MinInclusive string
	MaxExclusive string
	Actual       string
}

func (e *ErrVersionMismatch) Error() string {
	return fmt.Sprintf("unexpected version %s (min: %s, max: %s)", e.Actual, e.MinInclusive, e.MaxExclusive)
}

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

// ErrCLIUsage is returned when the combination of flags or arguments is incorrect.
//
//  CLI indicates usage errors in three different ways: either
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

// ErrManualEnvVar is returned when an env var that should be set programatically via an option or method
// is set via the manual environment passing functions.
type ErrManualEnvVar struct {
	name string
}

func (err *ErrManualEnvVar) Error() string {
	return fmt.Sprintf("manual setting of env var %q detected", err.name)
}

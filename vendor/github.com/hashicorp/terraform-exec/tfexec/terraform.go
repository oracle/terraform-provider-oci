package tfexec

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Terraform struct {
	execPath    string
	workingDir  string
	execVersion string
	env         []string
	logger      *log.Logger
}

// NewTerraform returns a Terraform struct with default values for all fields.
// If a blank execPath is supplied, NewTerraform will attempt to locate an
// appropriate binary on the system PATH.
func NewTerraform(workingDir string, execPath string) (*Terraform, error) {
	var err error
	if workingDir == "" {
		return nil, fmt.Errorf("Terraform cannot be initialised with empty workdir")
	}

	if _, err := os.Stat(workingDir); err != nil {
		return nil, fmt.Errorf("error initialising Terraform with workdir %s: %s", workingDir, err)
	}

	if execPath == "" {
		err := fmt.Errorf("NewTerraform: please supply the path to a Terraform executable using execPath, e.g. using the tfinstall package.")
		return nil, &ErrNoSuitableBinary{err: err}

	}
	tf := Terraform{
		execPath:   execPath,
		workingDir: workingDir,
		env:        os.Environ(),
		logger:     log.New(ioutil.Discard, "", 0),
	}

	execVersion, err := tf.version()
	if err != nil {
		return nil, &ErrNoSuitableBinary{err: fmt.Errorf("error running 'terraform version': %s", err)}
	}

	tf.execVersion = execVersion

	return &tf, nil
}

func (tf *Terraform) SetEnv(env map[string]string) {
	var tfenv []string

	// always propagate CHECKPOINT_DISABLE env var unless it is
	// explicitly overridden with tf.SetEnv
	if _, ok := env["CHECKPOINT_DISABLE"]; !ok {
		env["CHECKPOINT_DISABLE"] = os.Getenv("CHECKPOINT_DISABLE")
	}

	for k, v := range env {
		tfenv = append(tfenv, k+"="+v)
	}

	tf.env = tfenv
}

func (tf *Terraform) SetLogger(logger *log.Logger) {
	tf.logger = logger
}

func (tf *Terraform) version() (string, error) {
	versionCmd := tf.buildTerraformCmd(context.Background(), "version")
	var errBuf strings.Builder
	var outBuf bytes.Buffer
	versionCmd.Stderr = &errBuf
	versionCmd.Stdout = &outBuf

	err := versionCmd.Run()
	if err != nil {
		fmt.Println(errBuf.String())
		return "", fmt.Errorf("%s, %s", err, errBuf.String())
	}

	return outBuf.String(), nil
}

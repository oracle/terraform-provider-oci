package tfexec

import (
	"fmt"
	"os"
	"os/exec"
)

// FindTerraform attempts to find a Terraform CLI executable.
//
// As a first preference it will look for the environment variable
// TFEXEC_TERRAFORM_PATH and return its value. If that variable is not set, it will
// look in PATH for a program named "terraform",
// and, if one is found, return its absolute path.
func FindTerraform() (string, error) {
	if p := os.Getenv("TFEXEC_TERRAFORM_PATH"); p != "" {
		return p, nil
	}

	execName := "terraform"

	p, err := exec.LookPath(execName)
	if err != nil {
		return "", fmt.Errorf("terraform executable could not be found: %s", err)
	}
	return p, nil
}

package oci

import (
	"os/exec"
)

// runTerraform runs the configured Terraform CLI executable with the given
// arguments, returning an error if it produces a non-successful exit status.
func runTerraform(args []string) error {
	terraformExecPath, err := exec.LookPath("terraform")
	if err != nil {
		return err
	}

	allArgs := []string{"terraform"}
	allArgs = append(allArgs, args...)

	cmd := &exec.Cmd{
		Path: terraformExecPath,
		Args: allArgs,
	}
	err = cmd.Run()
	return err
}

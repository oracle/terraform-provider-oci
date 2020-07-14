package tfexec

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"strings"

	tfjson "github.com/hashicorp/terraform-json"
)

func (tf *Terraform) StateShow(ctx context.Context) (*tfjson.State, error) {
	var ret tfjson.State

	var errBuf strings.Builder
	var outBuf bytes.Buffer

	showCmd := tf.stateShowCmd(ctx)

	showCmd.Stderr = &errBuf
	showCmd.Stdout = &outBuf

	err := showCmd.Run()
	if err != nil {
		return nil, parseError(errBuf.String())
	}

	err = json.Unmarshal(outBuf.Bytes(), &ret)
	if err != nil {
		return nil, err
	}

	err = ret.Validate()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (tf *Terraform) stateShowCmd(ctx context.Context, args ...string) *exec.Cmd {
	allArgs := []string{"show", "-json", "-no-color"}
	allArgs = append(allArgs, args...)

	return tf.buildTerraformCmd(ctx, allArgs...)
}

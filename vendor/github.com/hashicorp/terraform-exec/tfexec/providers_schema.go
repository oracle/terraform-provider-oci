package tfexec

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"strings"

	tfjson "github.com/hashicorp/terraform-json"
)

func (tf *Terraform) ProvidersSchema(ctx context.Context) (*tfjson.ProviderSchemas, error) {
	var ret tfjson.ProviderSchemas

	var errBuf strings.Builder
	var outBuf bytes.Buffer

	schemaCmd := tf.providersSchemaCmd(ctx)

	schemaCmd.Stderr = &errBuf
	schemaCmd.Stdout = &outBuf

	err := schemaCmd.Run()
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

func (tf *Terraform) providersSchemaCmd(ctx context.Context, args ...string) *exec.Cmd {
	allArgs := []string{"providers", "schema", "-json", "-no-color"}
	allArgs = append(allArgs, args...)

	return tf.buildTerraformCmd(ctx, allArgs...)
}

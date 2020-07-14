package tfexec

import (
	"context"
	"os/exec"
)

func (tf *Terraform) buildTerraformCmd(ctx context.Context, args ...string) *exec.Cmd {
	env := append(tf.env, "TF_LOG=") // so logging can't pollute our stderr output

	cmd := exec.CommandContext(ctx, tf.execPath, args...)
	cmd.Env = env
	cmd.Dir = tf.workingDir

	tf.logger.Printf("Terraform command: %s", cmdString(cmd))

	return cmd
}

package tfexec

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
)

func (tf *Terraform) buildTerraformCmd(ctx context.Context, args ...string) *exec.Cmd {
	env := append(tf.env, "TF_LOG=") // so logging can't pollute our stderr output

	cmd := exec.CommandContext(ctx, tf.execPath, args...)
	cmd.Env = env
	cmd.Dir = tf.workingDir

	tf.logger.Printf("Terraform command: %s", cmd.String())

	return cmd
}

func (tf *Terraform) InitCmd(ctx context.Context, opts ...InitOption) *exec.Cmd {
	c := defaultInitOptions

	for _, o := range opts {
		o.configureInit(&c)
	}

	args := []string{"init", "-no-color", "-force-copy", "-input=false"}

	// string opts: only pass if set
	if c.fromModule != "" {
		args = append(args, "-from-module="+c.fromModule)
	}
	if c.lockTimeout != "" {
		args = append(args, "-lock-timeout="+c.lockTimeout)
	}

	// boolean opts: always pass
	args = append(args, "-backend="+fmt.Sprint(c.backend))
	args = append(args, "-get="+fmt.Sprint(c.get))
	args = append(args, "-get-plugins="+fmt.Sprint(c.getPlugins))
	args = append(args, "-lock="+fmt.Sprint(c.lock))
	args = append(args, "-upgrade="+fmt.Sprint(c.upgrade))
	args = append(args, "-verify-plugins="+fmt.Sprint(c.verifyPlugins))

	// unary flags: pass if true
	if c.reconfigure {
		args = append(args, "-reconfigure")
	}

	// string slice opts: split into separate args
	if c.backendConfig != nil {
		for _, bc := range c.backendConfig {
			args = append(args, "-backend-config="+bc)
		}
	}
	if c.pluginDir != nil {
		for _, pd := range c.pluginDir {
			args = append(args, "-plugin-dir="+pd)
		}
	}

	return tf.buildTerraformCmd(ctx, args...)
}

func (tf *Terraform) ApplyCmd(ctx context.Context, opts ...ApplyOption) *exec.Cmd {
	c := defaultApplyOptions

	for _, o := range opts {
		o.configureApply(&c)
	}

	args := []string{"apply", "-no-color", "-auto-approve", "-input=false"}

	// string opts: only pass if set
	if c.backup != "" {
		args = append(args, "-backup="+c.backup)
	}
	if c.lockTimeout != "" {
		args = append(args, "-lock-timeout="+c.lockTimeout)
	}
	if c.state != "" {
		args = append(args, "-state="+c.state)
	}
	if c.stateOut != "" {
		args = append(args, "-state-out="+c.stateOut)
	}
	if c.varFile != "" {
		args = append(args, "-var-file="+c.varFile)
	}

	// boolean and numerical opts: always pass
	args = append(args, "-lock="+strconv.FormatBool(c.lock))
	args = append(args, "-parallelism="+fmt.Sprint(c.parallelism))
	args = append(args, "-refresh="+strconv.FormatBool(c.refresh))

	// string slice opts: split into separate args
	if c.targets != nil {
		for _, ta := range c.targets {
			args = append(args, "-target="+ta)
		}
	}
	if c.vars != nil {
		for _, v := range c.vars {
			args = append(args, "-var '"+v+"'")
		}
	}

	// string argument: pass if set
	if c.dirOrPlan != "" {
		args = append(args, c.dirOrPlan)
	}

	return tf.buildTerraformCmd(ctx, args...)
}

func (tf *Terraform) DestroyCmd(ctx context.Context, opts ...DestroyOption) *exec.Cmd {
	c := defaultDestroyOptions

	for _, o := range opts {
		o.configureDestroy(&c)
	}

	args := []string{"destroy", "-no-color", "-auto-approve"}

	// string opts: only pass if set
	if c.backup != "" {
		args = append(args, "-backup="+c.backup)
	}
	if c.lockTimeout != "" {
		args = append(args, "-lock-timeout="+c.lockTimeout)
	}
	if c.state != "" {
		args = append(args, "-state="+c.state)
	}
	if c.stateOut != "" {
		args = append(args, "-state-out="+c.stateOut)
	}
	if c.varFile != "" {
		args = append(args, "-var-file="+c.varFile)
	}

	// boolean and numerical opts: always pass
	args = append(args, "-lock="+strconv.FormatBool(c.lock))
	args = append(args, "-parallelism="+fmt.Sprint(c.parallelism))
	args = append(args, "-refresh="+strconv.FormatBool(c.refresh))

	// string slice opts: split into separate args
	if c.targets != nil {
		for _, ta := range c.targets {
			args = append(args, "-target="+ta)
		}
	}
	if c.vars != nil {
		for _, v := range c.vars {
			args = append(args, "-var '"+v+"'")
		}
	}

	return tf.buildTerraformCmd(ctx, args...)
}

func (tf *Terraform) PlanCmd(ctx context.Context, opts ...PlanOption) *exec.Cmd {
	c := defaultPlanOptions

	for _, o := range opts {
		o.configurePlan(&c)
	}

	args := []string{"plan", "-no-color", "-input=false"}

	// string opts: only pass if set
	if c.lockTimeout != "" {
		args = append(args, "-lock-timeout="+c.lockTimeout)
	}
	if c.out != "" {
		args = append(args, "-out="+c.out)
	}
	if c.state != "" {
		args = append(args, "-state="+c.state)
	}
	if c.varFile != "" {
		args = append(args, "-var-file="+c.varFile)
	}

	// boolean and numerical opts: always pass
	args = append(args, "-lock="+strconv.FormatBool(c.lock))
	args = append(args, "-parallelism="+fmt.Sprint(c.parallelism))
	args = append(args, "-refresh="+strconv.FormatBool(c.refresh))

	// unary flags: pass if true
	if c.destroy {
		args = append(args, "-destroy")
	}

	// string slice opts: split into separate args
	if c.targets != nil {
		for _, ta := range c.targets {
			args = append(args, "-target="+ta)
		}
	}
	if c.vars != nil {
		for _, v := range c.vars {
			args = append(args, "-var '"+v+"'")
		}
	}

	return tf.buildTerraformCmd(ctx, args...)
}

func (tf *Terraform) ImportCmd(ctx context.Context, opts ...ImportOption) *exec.Cmd {
	c := defaultImportOptions

	for _, o := range opts {
		o.configureImport(&c)
	}

	args := []string{"import", "-no-color", "-input=false"}

	// string opts: only pass if set
	if c.backup != "" {
		args = append(args, "-backup="+c.backup)
	}
	if c.config != "" {
		args = append(args, "-config"+c.config)
	}
	if c.lockTimeout != "" {
		args = append(args, "-lock-timeout="+c.lockTimeout)
	}
	if c.state != "" {
		args = append(args, "-state="+c.state)
	}
	if c.stateOut != "" {
		args = append(args, "-state-out="+c.stateOut)
	}
	if c.varFile != "" {
		args = append(args, "-var-file="+c.varFile)
	}

	// boolean and numerical opts: always pass
	args = append(args, "-lock="+strconv.FormatBool(c.lock))

	// unary flags: pass if true
	if c.allowMissingConfig {
		args = append(args, "-allow-missing-config")
	}

	// string slice opts: split into separate args
	if c.vars != nil {
		for _, v := range c.vars {
			args = append(args, "-var '"+v+"'")
		}
	}

	return tf.buildTerraformCmd(ctx, args...)
}

func (tf *Terraform) OutputCmd(ctx context.Context, opts ...OutputOption) *exec.Cmd {
	c := defaultOutputOptions

	for _, o := range opts {
		o.configureOutput(&c)
	}

	args := []string{"output", "-no-color", "-json"}

	// string opts: only pass if set
	if c.state != "" {
		args = append(args, "-state="+c.state)
	}

	return tf.buildTerraformCmd(ctx, args...)
}

func (tf *Terraform) StateShowCmd(ctx context.Context, args ...string) *exec.Cmd {
	allArgs := []string{"show", "-json", "-no-color"}
	allArgs = append(allArgs, args...)

	return tf.buildTerraformCmd(ctx, allArgs...)
}

func (tf *Terraform) ProvidersSchemaCmd(ctx context.Context, args ...string) *exec.Cmd {
	allArgs := []string{"providers", "schema", "-json", "-no-color"}
	allArgs = append(allArgs, args...)

	return tf.buildTerraformCmd(ctx, allArgs...)
}

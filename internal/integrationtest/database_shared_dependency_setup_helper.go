// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

type terraformStateFile struct {
	Resources []terraformStateResource `json:"resources"`
}

type terraformStateResource struct {
	Module    string                   `json:"module"`
	Mode      string                   `json:"mode"`
	Type      string                   `json:"type"`
	Name      string                   `json:"name"`
	Instances []terraformStateInstance `json:"instances"`
}

type terraformStateInstance struct {
	Attributes map[string]any `json:"attributes"`
}

type testLogWriter struct {
	t *testing.T
}

func (w testLogWriter) Write(p []byte) (int, error) {
	text := strings.TrimRight(string(p), "\n")
	if text != "" {
		w.t.Log(text)
	}
	return len(p), nil
}

// ResolveOrCreateSharedDependenciesFromConfig reuses existing shared resource
// ids when all requested resources are supplied; otherwise it creates the
// dependencies from the provided Terraform configuration.
func ResolveOrCreateSharedDependenciesFromConfig(t *testing.T, existingResourceIDs map[string]string, setupConfig string, resourceAddresses []string) (map[string]string, func()) {
	t.Helper()

	if len(resourceAddresses) == 0 {
		t.Fatal("at least one resource address must be provided")
	}

	resolvedResourceIDs := make(map[string]string, len(resourceAddresses))
	canReuseAll := true
	for _, resourceAddress := range resourceAddresses {
		existingResourceID := strings.TrimSpace(existingResourceIDs[resourceAddress])
		if !isUsableSharedDependencyID(existingResourceID) {
			canReuseAll = false
			break
		}
		resolvedResourceIDs[resourceAddress] = existingResourceID
	}

	if canReuseAll {
		for _, resourceAddress := range resourceAddresses {
			log.Printf("[SHARED_DEP_SETUP] reusing shared dependency %s with id %s", resourceAddress, resolvedResourceIDs[resourceAddress])
		}
		return resolvedResourceIDs, nil
	}

	log.Printf("[SHARED_DEP_SETUP] creating shared dependencies %v", resourceAddresses)
	return SetupSharedDependenciesFromConfig(t, setupConfig, resourceAddresses)
}

func isUsableSharedDependencyID(resourceID string) bool {
	resourceID = strings.TrimSpace(resourceID)
	if resourceID == "" {
		return false
	}

	if strings.HasPrefix(resourceID, "${") && strings.HasSuffix(resourceID, "}") {
		return false
	}

	if strings.HasPrefix(resourceID, "TF_VAR_") {
		return false
	}

	return true
}

// SetupSharedDependenciesFromConfig applies the supplied Terraform configuration
// once in a temporary workspace and returns the ids for each requested resource
// address. The returned cleanup destroys the same workspace exactly once.
func SetupSharedDependenciesFromConfig(t *testing.T, setupConfig string, resourceAddresses []string) (map[string]string, func()) {
	t.Helper()

	if strings.TrimSpace(setupConfig) == "" {
		t.Fatal("setup config must not be empty")
	}

	if len(resourceAddresses) == 0 {
		t.Fatal("at least one resource address must be provided")
	}

	workingDirectory := t.TempDir()
	configPath := filepath.Join(workingDirectory, "main.tf")
	terraformCLIEnv, terraformProviderBootstrapMode, err := terraformWorkspaceProviderEnv(workingDirectory)
	if err != nil {
		t.Fatalf("failed to configure local provider installation for shared dependency setup: %v", err)
	}

	fullSetupConfig := terraformProviderBootstrapConfig(terraformProviderBootstrapMode) + setupConfig
	if err := os.WriteFile(configPath, []byte(fullSetupConfig), 0o600); err != nil {
		t.Fatalf("failed to write shared dependency config: %v", err)
	}

	destroySharedDependencyOnFailure := func(failureReason string) {
		if err := destroyTerraformWorkspace(t, workingDirectory, terraformCLIEnv); err != nil {
			t.Fatalf("%s; additionally failed to destroy shared dependency resources: %v", failureReason, err)
		}
		t.Fatal(failureReason)
	}

	if _, err := runTerraformWorkspaceCommand(t, workingDirectory, terraformCLIEnv, "init", "-input=false", "-no-color"); err != nil {
		t.Fatalf("terraform init failed for shared dependency setup: %v", err)
	}
	if _, err := runTerraformWorkspaceCommand(t, workingDirectory, terraformCLIEnv, "apply", "-auto-approve", "-input=false", "-no-color"); err != nil {
		destroySharedDependencyOnFailure(fmt.Sprintf("terraform apply failed for shared dependency setup: %v", err))
	}

	resourceIDs := make(map[string]string, len(resourceAddresses))
	for _, resourceAddress := range resourceAddresses {
		if strings.TrimSpace(resourceAddress) == "" {
			t.Fatal("Resource address must not be empty")
		}
		log.Printf("[SHARED_DEP_SETUP] applied setup config for resource address %s", resourceAddress)
		resourceID, err := terraformResourceIDForAddress(workingDirectory, resourceAddress)
		if err != nil {
			destroySharedDependencyOnFailure(fmt.Sprintf("failed to capture shared dependency id for %s: %v", resourceAddress, err))
		}
		log.Printf("[SHARED_DEP_SETUP] captured resource address %s with id %s", resourceAddress, resourceID)
		resourceIDs[resourceAddress] = resourceID
	}

	var destroyOnce sync.Once
	cleanup := func() {
		destroyOnce.Do(func() {
			if err := destroyTerraformWorkspace(t, workingDirectory, terraformCLIEnv); err != nil {
				t.Errorf("terraform destroy failed for shared dependencies %v: %v", resourceAddresses, err)
			}
		})
	}

	return resourceIDs, cleanup
}

func runTerraformWorkspaceCommand(t *testing.T, workingDirectory string, extraEnv []string, args ...string) (string, error) {
	t.Helper()

	terraformPath, err := exec.LookPath("terraform")
	if err != nil {
		return "", fmt.Errorf("terraform binary not found in PATH: %w", err)
	}

	cmd := exec.Command(terraformPath, args...)
	cmd.Dir = workingDirectory
	cmd.Env = append(os.Environ(), "TF_IN_AUTOMATION=1")
	cmd.Env = append(cmd.Env, extraEnv...)

	var output bytes.Buffer
	writer := io.MultiWriter(&output, testLogWriter{t: t})
	cmd.Stdout = writer
	cmd.Stderr = writer

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("terraform %s failed: %w\n%s", strings.Join(args, " "), err, output.String())
	}

	return output.String(), nil
}

func terraformResourceIDForAddress(workingDirectory string, resourceAddress string) (string, error) {
	statePath := filepath.Join(workingDirectory, "terraform.tfstate")
	output, err := os.ReadFile(statePath)
	if err != nil {
		return "", fmt.Errorf("failed to read terraform state file %s: %w", statePath, err)
	}

	var terraformState terraformStateFile
	if err := json.Unmarshal([]byte(output), &terraformState); err != nil {
		return "", fmt.Errorf("failed to parse terraform state file: %w", err)
	}
	resourceID, found := terraformResourceIDFromState(terraformState, resourceAddress)
	if !found {
		return "", fmt.Errorf("resource %s not found in terraform state", resourceAddress)
	}
	if resourceID == "" {
		return "", fmt.Errorf("resource %s does not have an id value", resourceAddress)
	}

	return resourceID, nil
}

func terraformResourceIDFromState(terraformState terraformStateFile, resourceAddress string) (string, bool) {
	for _, resource := range terraformState.Resources {
		if terraformStateResourceAddress(resource) != resourceAddress {
			continue
		}

		for _, instance := range resource.Instances {
			idValue, ok := instance.Attributes["id"]
			if !ok {
				continue
			}

			resourceID, ok := idValue.(string)
			if !ok {
				continue
			}

			return resourceID, true
		}

		return "", true
	}

	return "", false
}

func terraformStateResourceAddress(resource terraformStateResource) string {
	address := fmt.Sprintf("%s.%s", resource.Type, resource.Name)
	if resource.Module != "" {
		return fmt.Sprintf("%s.%s", resource.Module, address)
	}
	return address
}

func destroyTerraformWorkspace(t *testing.T, workingDirectory string, extraEnv []string) error {
	if _, err := runTerraformWorkspaceCommand(t, workingDirectory, extraEnv, "init", "-input=false", "-no-color"); err != nil {
		return fmt.Errorf("terraform init before destroy failed: %w", err)
	}
	if _, err := runTerraformWorkspaceCommand(t, workingDirectory, extraEnv, "destroy", "-auto-approve", "-input=false", "-no-color"); err != nil {
		return err
	}
	return nil
}

func terraformProviderBootstrapConfig(terraformProviderBootstrapMode string) string {
	if terraformProviderBootstrapMode == "local" {
		return `provider "oci" {}

`
	}

	return fmt.Sprintf(`terraform {
  required_providers {
    oci = {
      source  = "oracle/oci"
      version = "= %s"
    }
  }
}

provider "oci" {}

`, globalvar.Version)
}

func terraformWorkspaceProviderEnv(workingDirectory string) ([]string, string, error) {
	usePipelineProviderBootstrap, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("use_pipeline_provider_bootstrap", "false"))
	if !usePipelineProviderBootstrap {
		log.Printf("[SHARED_DEP_SETUP] provider bootstrap mode=local")
		providerBinaryPath := utils.GetEnvSettingWithBlankDefault("provider_bin_path")
		if providerBinaryPath == "" {
			return nil, "", fmt.Errorf("TF_VAR_provider_bin_path must be set for shared dependency setup")
		}

		resolvedProviderBinaryPath, err := resolveProviderBinaryPath(providerBinaryPath)
		if err != nil {
			return nil, "", err
		}

		localPluginDir, err := ensureLocalTerraformPluginDir(resolvedProviderBinaryPath)
		if err != nil {
			return nil, "", err
		}

		log.Printf("[SHARED_DEP_SETUP] using local provider binary %s", resolvedProviderBinaryPath)
		log.Printf("[SHARED_DEP_SETUP] using local terraform plugin dir %s", localPluginDir)
		return nil, "local", nil
	}

	providerBinaryPath := utils.GetEnvSettingWithBlankDefault("provider_bin_path")
	if providerBinaryPath == "" {
		return nil, "", fmt.Errorf("TF_VAR_provider_bin_path must be set for shared dependency setup")
	}

	resolvedProviderBinaryPath, err := resolveProviderBinaryPath(providerBinaryPath)
	if err != nil {
		return nil, "", err
	}

	mirrorRoot := filepath.Join(workingDirectory, ".terraform.d", "plugins")
	if err := os.MkdirAll(mirrorRoot, 0o755); err != nil {
		return nil, "", fmt.Errorf("failed to create Terraform mirror root: %w", err)
	}

	providerBinaryName := filepath.Base(resolvedProviderBinaryPath)
	versionedProviderBinaryName := fmt.Sprintf("%s_v%s", providerBinaryName, globalvar.Version)

	for _, platformDir := range terraformProviderPlatformDirs() {
		destinationDir := filepath.Join(mirrorRoot, "registry.terraform.io", "oracle", "oci", globalvar.Version, platformDir)
		if err := os.MkdirAll(destinationDir, 0o755); err != nil {
			return nil, "", fmt.Errorf("failed to create provider mirror directory %s: %w", destinationDir, err)
		}

		unversionedDestination := filepath.Join(destinationDir, providerBinaryName)
		if err := copyFile(resolvedProviderBinaryPath, unversionedDestination); err != nil {
			return nil, "", fmt.Errorf("failed to copy provider binary to %s: %w", unversionedDestination, err)
		}

		versionedDestination := filepath.Join(destinationDir, versionedProviderBinaryName)
		if err := copyFile(resolvedProviderBinaryPath, versionedDestination); err != nil {
			return nil, "", fmt.Errorf("failed to copy versioned provider binary to %s: %w", versionedDestination, err)
		}

		log.Printf("[SHARED_DEP_SETUP] copied provider binary to %s", unversionedDestination)
		log.Printf("[SHARED_DEP_SETUP] copied versioned provider binary to %s", versionedDestination)
	}

	terraformCLIConfigPath := filepath.Join(workingDirectory, "terraform.rc")
	terraformCLIConfig := fmt.Sprintf(`provider_installation {
  filesystem_mirror {
    path    = %q
    include = ["registry.terraform.io/oracle/oci"]
  }
  direct {
    exclude = ["registry.terraform.io/oracle/oci"]
  }
}
`, mirrorRoot)
	if err := os.WriteFile(terraformCLIConfigPath, []byte(terraformCLIConfig), 0o600); err != nil {
		return nil, "", fmt.Errorf("failed to write Terraform CLI config: %w", err)
	}

	log.Printf("[SHARED_DEP_SETUP] using provider binary %s", resolvedProviderBinaryPath)
	log.Printf("[SHARED_DEP_SETUP] using provider mirror root %s", mirrorRoot)
	log.Printf("[SHARED_DEP_SETUP] using terraform rc %s", terraformCLIConfigPath)
	log.Printf("[SHARED_DEP_SETUP] provider bootstrap mode=pipeline")

	return []string{fmt.Sprintf("TF_CLI_CONFIG_FILE=%s", terraformCLIConfigPath)}, "pipeline", nil
}

func resolveProviderBinaryPath(providerBinPath string) (string, error) {
	info, err := os.Stat(providerBinPath)
	if err != nil {
		return "", fmt.Errorf("failed to stat TF_VAR_provider_bin_path %s: %w", providerBinPath, err)
	}

	if !info.IsDir() {
		return providerBinPath, nil
	}

	candidates := []string{
		filepath.Join(providerBinPath, "terraform-provider-oci"),
		filepath.Join(providerBinPath, fmt.Sprintf("terraform-provider-oci_v%s", globalvar.Version)),
	}

	for _, candidate := range candidates {
		candidateInfo, candidateErr := os.Stat(candidate)
		if candidateErr == nil && !candidateInfo.IsDir() {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("TF_VAR_provider_bin_path %s is a directory and no terraform-provider-oci binary was found in it", providerBinPath)
}

func terraformProviderPlatformDirs() []string {
	return []string{
		"linux_arm",
	}
}

func ensureLocalTerraformPluginDir(resolvedProviderBinaryPath string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to resolve user home directory: %w", err)
	}

	localPluginDir := filepath.Join(homeDir, ".terraform.d", "plugins")
	if err := os.MkdirAll(localPluginDir, 0o755); err != nil {
		return "", fmt.Errorf("failed to create local Terraform plugin directory: %w", err)
	}

	providerBinaryName := filepath.Base(resolvedProviderBinaryPath)
	unversionedDestination := filepath.Join(localPluginDir, providerBinaryName)
	if err := copyFile(resolvedProviderBinaryPath, unversionedDestination); err != nil {
		return "", fmt.Errorf("failed to copy local provider binary to %s: %w", unversionedDestination, err)
	}

	versionedDestination := filepath.Join(localPluginDir, fmt.Sprintf("%s_v%s", providerBinaryName, globalvar.Version))
	if err := copyFile(resolvedProviderBinaryPath, versionedDestination); err != nil {
		return "", fmt.Errorf("failed to copy versioned local provider binary to %s: %w", versionedDestination, err)
	}

	return localPluginDir, nil
}

func copyFile(sourcePath string, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := sourceFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	info, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	destinationFile, err := os.OpenFile(destinationPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := destinationFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	if _, err := io.Copy(destinationFile, sourceFile); err != nil {
		return err
	}

	return nil
}

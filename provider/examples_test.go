// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const examplesTestStateFile = "test_examples.tfstate"
const defaultTerraformBinary = "terraform"
const vcnExamplePath = "../docs/examples/networking/vcn"
const localBinPath = "/usr/local/bin"
const tfPluginDir = "~/.terraform.d/plugins"

var examplesTestAllowedEnvironmentVariables = []string{
	"HOME",
	"PATH",
	"TF_VAR_user_ocid",
	"TF_VAR_tenancy_ocid",
	"TF_VAR_fingerprint",
	"TF_VAR_private_key_path",
	"TF_VAR_private_key_password",
	"TF_VAR_region",
	"TF_VAR_compartment_ocid",
	"TF_VAR_AD",
	"TF_VAR_ssh_public_key",
	"TF_VAR_ssh_private_key",
}

func TestExamplesPlan(t *testing.T) {
	RunExamples(t, true)
}

func TestExamplesApply(t *testing.T) {
	RunExamples(t, false)
}

func TestTerraformVersions(t *testing.T) {
	if RunConfigOnAllTerraformVersions(t, vcnExamplePath, false) {
		log.Printf("Successfully ran all Terraform version tests")
	}
}

func RunExamples(t *testing.T, planOnly bool) {
	rootPath := "../docs/examples"
	log.Printf("Testing examples under %v", rootPath)

	pathList, err := GetConfigPaths(t, rootPath)
	if err != nil {
		t.Errorf("Error scanning directories: %v", err)
		return
	}

	for _, dir := range pathList {
		if RunConfig(t, dir, planOnly, defaultTerraformBinary) {
			log.Printf("Success")
		}
	}
}

func GetTerraformBinaries(binPath string) ([]string, error) {
	results := []string{}

	entries, err := ioutil.ReadDir(binPath)
	if err != nil {
		return results, err
	}

	for _, entry := range entries {
		// Include any binaries that start with "terraform" prefix
		if name := entry.Name(); !entry.IsDir() && strings.HasPrefix(name, defaultTerraformBinary) {
			results = append(results, name)
		}
	}

	return results, nil
}

func GetConfigPaths(t *testing.T, rootPath string) (pathList []string, err error) {
	dirSet := make(map[string]struct{})

	var fileScanner = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
		// Assume that all directories containing *.tf files are configs that should be tested.
		// This assumption may need to be updated if we add examples that use modules.
		if !fileInfo.IsDir() && strings.HasSuffix(path, ".tf") {
			dir := filepath.Dir(path)

			// TODO: Skip the db_systems example for now, until this is updated to use the
			// new set of environment variables.
			if !strings.HasSuffix(dir, "/db_systems") {
				dirSet[dir] = struct{}{}
			}
		}
		return nil
	}

	err = filepath.Walk(rootPath, fileScanner)
	pathList = make([]string, 0, len(dirSet))

	for dir := range dirSet {
		pathList = append(pathList, dir)
	}

	return pathList, err
}

func RunConfigOnAllTerraformVersions(t *testing.T, path string, planOnly bool) bool {
	terraformBinaries, err := GetTerraformBinaries(localBinPath)
	if err != nil {
		t.Errorf("Error retrieving terraform binaries: %v", err)
		return false
	}

	if len(terraformBinaries) == 0 {
		t.Errorf("Did not find any terraform binaries")
		return false
	}

	result := true
	for _, tfBin := range terraformBinaries {
		log.Printf("=== Terraform Version ('%s'), Config Path ('%s') ===\n", tfBin, path)
		if !runCommandWithOutputOptions(t, path, fmt.Sprintf("%s version", tfBin), true) {
			log.Printf("Unable to run version command. Skipping test for %s.\n", tfBin)
			result = false
			continue
		}

		if !RunConfig(t, path, planOnly, tfBin) {
			log.Printf("Failed to run test on version '%s'\n", tfBin)
			result = false
		}
	}

	return result
}

func RunConfig(t *testing.T, path string, planOnly bool, terraformBinary string) bool {
	// Fail if a state file already exists, since that indicates that a previous run did not
	// properly clean up.
	if _, err := os.Stat(filepath.Join(path, examplesTestStateFile)); err == nil {
		t.Errorf("State file '%v' already exists at %v.", examplesTestStateFile, path)
		return false
	}

	terraformCommand := terraformBinary
	if terraformCommand == "" {
		terraformCommand = defaultTerraformBinary
	}

	if !RunCommand(t, path, fmt.Sprintf("%s init -plugin-dir %s", terraformCommand, tfPluginDir)) {
		return false
	}

	if planOnly {
		return RunCommand(t, path, fmt.Sprintf("%s plan -state=%v", terraformCommand, examplesTestStateFile))
	} else {
		result := RunCommand(t, path, fmt.Sprintf("%s apply -auto-approve -state=%v", terraformCommand, examplesTestStateFile))

		// Regardless of the result, attempt to destroy.
		if RunCommand(t, path, fmt.Sprintf("%s destroy -force -state=%v", terraformCommand, examplesTestStateFile)) {
			// Only remove the state file if destroy was successful. Otherwise, leave it in place so that further
			// cleanup can be done manually.
			result = RunCommand(t, path, fmt.Sprintf("rm %v*", examplesTestStateFile)) && result
		} else {
			result = false
		}

		return result
	}
}

func RunCommand(t *testing.T, path, command string) bool {
	return runCommandWithOutputOptions(t, path, command, false)
}

func runCommandWithOutputOptions(t *testing.T, path, command string, displayOutputOnSuccess bool) bool {
	log.Printf("Running command '%v' at %v", command, path)

	env := make([]string, len(examplesTestAllowedEnvironmentVariables))
	for index, variable := range examplesTestAllowedEnvironmentVariables {
		env[index] = fmt.Sprintf("%v=%v", variable, os.Getenv(variable))
	}

	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = path
	cmd.Env = env
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Error: Command Failed. Output: %s", output)
		t.Errorf("Error running command %v at %v: %v", command, path, err)
		return false
	}

	if displayOutputOnSuccess {
		log.Printf("Output: %s", output)
	}

	return true
}

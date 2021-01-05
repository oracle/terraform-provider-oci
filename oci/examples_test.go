// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

const examplesTestStateFile = "test_examples.tfstate"
const defaultTerraformBinary = "terraform"
const vcnExamplePath = "../examples/networking/vcn"
const localBinPath = "/usr/local/bin"

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
	"TF_VAR_compartment_id",
	"TF_VAR_compartment_ocid_acceptor",
	"TF_VAR_compartment_ocid_requestor",
	"TF_VAR_availability_domain",
	"TF_VAR_ssh_public_key",
	"TF_VAR_ssh_private_key",
	"TF_VAR_compartment_name_acceptor",
	"TF_VAR_compartment_name_requestor",
	"TF_VAR_compartment_ocid_acceptor",
	"TF_VAR_compartment_ocid_requestor",
	"TF_VAR_fingerprint_acceptor",
	"TF_VAR_fingerprint_requestor",
	"TF_VAR_identity_provider_metadata_file",
	"TF_VAR_private_key_path_acceptor",
	"TF_VAR_private_key_path_requestor",
	"TF_VAR_ssh_public_key",
	"TF_VAR_ssh_private_key",
	"TF_VAR_user_acceptor",
	"TF_VAR_user_requestor",
	"TF_VAR_tags_import_if_exists",
	"TF_VAR_defined_tag_namespace_name",
	"TF_VAR_simulate_db",
	"TF_VAR_avoid_waiting_for_delete_target",
	"TF_VAR_vault_id",
	"TF_VAR_database_software_image_id",
}

func TestExamplesPlan(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skipping TestExamplesPlan")
	}

	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "TestExamplesPlan") {
		t.Skip("Skipping TestExamplesPlan")
	}
	RunExamples(t, true)
}

func TestExamplesApply(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skipping TestExamplesApply")
	}
	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "TestExamplesApply") {
		t.Skip("Skipping TestExamplesApply")
	}
	RunExamples(t, false)
}

func TestTerraformVersions(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skipping TestTerraformVersions")
	}
	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "TestTerraformVersions") {
		t.Skip("Skipping TestTerraformVersions")
	}
	if RunConfigOnAllTerraformVersions(t, vcnExamplePath, false) {
		log.Printf("Successfully ran all Terraform version tests")
	}
}

/*
 * Test for end-to-end Terraform import
 * The acceptance import tests have a test gap, because they do not cover Terraform's resource state to state file
 * normalization logic. Import tests only verify that in-memory state is correct.
 */
func TestTerraformImport(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skipping TestTerraformImport")
	}
	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "TestTerraformImport") {
		t.Skip("Skipping TestTerraformImport")
	}
	if RunConfigAndImport(t) {
		log.Printf("Successfully ran all Terraform import tests")
	}
}

func RunExamples(t *testing.T, planOnly bool) {
	rootPath := getEnvSettingWithDefault("examples_root", "../examples")
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

func GetTerraformBinaries(binPath string, targetPrefixes []string) ([]string, error) {
	results := []string{}

	entries, err := ioutil.ReadDir(binPath)
	if err != nil {
		return results, err
	}

	for _, entry := range entries {
		// Include any binaries that start with "terraform" prefix
		if name := entry.Name(); !entry.IsDir() && strings.HasPrefix(name, defaultTerraformBinary) {
			if len(targetPrefixes) == 0 {
				results = append(results, name)
			} else {
				for _, prefix := range targetPrefixes {
					if strings.HasPrefix(name, prefix) {
						results = append(results, name)
						break
					}
				}
			}
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

			// TODO: need improvement in our test environment
			// Skip some terraform example tests for now, since currently they do not fit
			// well in our testing environment
			if !shouldSkip(dir) {
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

// Helper function to ignore some folders
func shouldSkip(dir string) bool {
	blackList := strings.Split(os.Getenv("black_list"), ",")

	var flag bool
	for _, blackDir := range blackList {
		if len(blackDir) > 0 {
			flag = flag || strings.HasSuffix(dir, blackDir)
			if flag {
				return flag
			}
		}
	}
	return flag
}

func RunConfigOnAllTerraformVersions(t *testing.T, path string, planOnly bool) bool {
	terraformBinaries, err := GetTerraformBinaries(localBinPath, []string{})
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
		if _, result = runCommandWithOutputOptions(t, path, fmt.Sprintf("%s version", tfBin), true); !result {
			log.Printf("Unable to run version command. Skipping test for %s.\n", tfBin)
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

	if !RunCommand(t, path, fmt.Sprintf("%s init", terraformCommand)) {
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

func RunConfigAndImport(t *testing.T) bool {
	path := vcnExamplePath
	terraformBinaries, err := GetTerraformBinaries(localBinPath, []string{"terraform-0.11", "terraform-0.12"})
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
		var output []byte
		var vcnId string
		log.Printf("=== Terraform Version ('%s'), Config Path ('%s') ===\n", tfBin, path)
		if _, result := runCommandWithOutputOptions(t, path, fmt.Sprintf("%s version", tfBin), true); !result {
			log.Printf("Unable to run version command. Skipping test for %s.\n", tfBin)
			continue
		}

		if !RunCommand(t, path, fmt.Sprintf("%s init", tfBin)) {
			log.Printf("Unable to init terraform. Skipping test for %s.\n", tfBin)
			return false
		}

		// Create a VCN resource
		if result = RunCommand(t, path, fmt.Sprintf("%s apply -auto-approve -state=%v", tfBin, examplesTestStateFile)); !result {
			goto cleanup
		}

		// Get the ID of the VCN
		if output, result = runCommandWithOutputOptions(t, path, fmt.Sprintf("%s output -state=%v vcn_id", tfBin, examplesTestStateFile), true); !result {
			goto cleanup
		}
		vcnId = string(output)

		// Remove VCN from state file
		if result = RunCommand(t, path, fmt.Sprintf("%s state rm -state=%v oci_core_vcn.vcn1", tfBin, examplesTestStateFile)); !result {
			goto cleanup
		}

		// Import the VCN
		if result = RunCommand(t, path, fmt.Sprintf("%s import -state=%v oci_core_vcn.vcn1 %v", tfBin, examplesTestStateFile, vcnId)); !result {
			goto cleanup
		}

		// Plan should show no differences.
		if result = RunCommand(t, path, fmt.Sprintf("%s plan -detailed-exitcode -state=%v", tfBin, examplesTestStateFile)); !result {
			goto cleanup
		}

	cleanup:
		// Regardless of the result, attempt to destroy.
		if RunCommand(t, path, fmt.Sprintf("%s destroy -force -state=%v", tfBin, examplesTestStateFile)) {
			// Only remove the state file if destroy was successful. Otherwise, leave it in place so that further
			// cleanup can be done manually.
			result = RunCommand(t, path, fmt.Sprintf("rm %v*", examplesTestStateFile)) && result
		} else {
			result = false
		}

		if !result {
			break
		}
	}

	return result
}

func RunCommand(t *testing.T, path, command string) bool {
	_, result := runCommandWithOutputOptions(t, path, command, false)
	return result
}

func runCommandWithOutputOptions(t *testing.T, path, command string, displayOutputOnSuccess bool) ([]byte, bool) {
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
		return output, false
	}

	if displayOutputOnSuccess {
		log.Printf("Output: %s", output)
	}

	return output, true
}

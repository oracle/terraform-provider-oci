// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

type (
	structConfigFile struct {
		Environment *map[string]string `yaml:"env-var"`
		Steps       []configStep       `yaml:"steps"`
	}

	configStep struct {
		LogMessage           *string       `yaml:"log-message"`
		TerraformSuccess     *configFiles  `yaml:"tf-success"`
		TerraformFail        *configFiles  `yaml:"tf-fail"`
		StateConfigCheck     *configChecks `yaml:"state"`
		StateConfigCheckFile *string       `yaml:"state-file"`
	}

	configFiles  []string
	configChecks map[string]map[string]string
)

const (
	terraformStateFile = "terraform.tfstate"
	stateKeyExistsFlag = "-EXIST-"
)

// examplesTestAllowedEnvironmentVariables is defined in examples_test.go
// var examplesTestAllowedEnvironmentVariables = []string{
// 	"PATH",
// 	"TF_VAR_user_ocid",
// 	"TF_VAR_tenancy_ocid",
// 	"TF_VAR_fingerprint",
// 	"TF_VAR_private_key_path",
// 	"TF_VAR_private_key_password",
// 	"TF_VAR_region",
// 	"TF_VAR_compartment_ocid",
// 	"TF_VAR_AD",
// 	"TF_VAR_ssh_public_key",
// 	"TF_VAR_ssh_private_key",
// }

func TestAllConfigPlan(t *testing.T) {
	runTestConfigs(t, true)
}

func TestAllConfigApply(t *testing.T) {
	runTestConfigs(t, false)
}

func zTestCheck(t *testing.T) {
	//config, _ := readConfigFile(t, "../docs/sequencetests/1/2-vcn_simple.yaml")
	// doStateChecks(t, "/Users/anberg/Oracle content - Accounts/Oracle Content/Terraform/Test files/20180110/terraform.tfstate-after-simple_1", config.Steps[1].StateConfigChecks)
	//doStateChecks(t, "/Users/anberg/Oracle content - Accounts/Oracle Content/Terraform/Test files/20180110", config.Steps[1].StateConfigCheck)

	// output, runres := runCommand(t, "../docs/sequencetests/1/1-vcn_simple-test.yaml-run", "terraform destroy -force")
	// log.Printf("Ran command, result: %v\n%v", runres, output)

	content, err := ioutil.ReadFile("../docs/sequencetests/1/tf-show-output")
	checks := parseStateString(string(content))
	log.Printf("Loaded configs, got %v (err: %v)", checks, err)
}

func runTestConfigs(t *testing.T, planOnly bool) {
	rootPath := "../docs/sequencetests"
	log.Printf("Testing examples under %v", rootPath)

	configList, err := getConfigFilePaths(t, rootPath)
	if err != nil {
		t.Errorf("Error scanning directories: %v", err)
		return
	}

	for _, configPath := range configList {
		if runConfigFile(t, configPath, planOnly) {
			log.Printf("Success")
		}
	}
}

func getConfigFilePaths(t *testing.T, rootPath string) (configList []string, err error) {
	configSet := make(map[string]struct{})

	var fileScanner = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
		// We're going to look for <foo>-test.yaml as a test config.
		if !fileInfo.IsDir() && strings.HasSuffix(path, "-test.yaml") {
			configSet[path] = struct{}{}
		}
		return nil
	}

	err = filepath.Walk(rootPath, fileScanner)
	configList = make([]string, 0, len(configSet))

	for configPath := range configSet {
		configList = append(configList, configPath)
	}

	return configList, err
}

func readConfigFile(t *testing.T, configPath string) (config structConfigFile, err error) {
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		t.Errorf("Unable to read config file '%v'", configPath)
		return
	}

	if err = yaml.Unmarshal([]byte(configData), &config); err != nil {
		t.Errorf("Unable to parse config file at '%v': %v", configPath, err)
	}
	return
}

type executionState struct {
	t                             *testing.T
	planOnly                      bool
	configPath, configDir, runDir string
	initDone                      bool
	result                        bool
	terraformOp                   string
	configFileCache               map[string][]byte
	expectedResult                bool
	successfulApplyStep           *configStep
	lastConfigFiles               *configFiles
}

func runConfigFile(t *testing.T, configPath string, planOnly bool) bool {
	config, err := readConfigFile(t, configPath)
	if nil != err {
		return false
	}

	runDir := configPath + "-run"
	if err = os.Mkdir(runDir, os.ModePerm); err != nil {
		t.Errorf("Error: Unable to make the run directory '%v'", runDir)
		return false
	}
	defer os.Remove(runDir)

	// Fail if a state file already exists, since that indicates that a previous run did not properly clean up.
	if _, err = os.Stat(filepath.Join(runDir, terraformStateFile)); err == nil {
		t.Errorf("Error: State file '%v' already exists at %v.", terraformStateFile, runDir)
		return false
	}

	state := executionState{
		t:               t,
		planOnly:        planOnly,
		configPath:      configPath,
		configDir:       filepath.Dir(configPath),
		runDir:          runDir,
		initDone:        false,
		result:          true,
		configFileCache: make(map[string][]byte),
	}
	if planOnly {
		state.terraformOp = "plan"
		defer os.RemoveAll(runDir)
	} else {
		state.terraformOp = "apply --auto-approve"
		defer func() {
			// If we've ever successfully applied anything, try to do a 'terraform destroy' operation.
			if nil != state.successfulApplyStep {
				// Terraform seems to not like to destroy if there is no .tf config file in the directory.
				setupTerraformFilesForStep(&state, state.successfulApplyStep)
				if _, success := runCommand(t, runDir, "terraform destroy -force"); success {
					err := os.RemoveAll(runDir)
					state.result = state.result && err == nil
				} else {
					state.result = false
				}
				cleanupTerraformFiles(&state)
			}
		}()
	}

	for _, step := range config.Steps {
		processStep(&state, step)
		if !state.result {
			break
		}
	}

	return state.result
}

func processStep(state *executionState, step configStep) {
	if step.LogMessage != nil {
		state.t.Log(*step.LogMessage)
		log.Print(*step.LogMessage)
	}
	if step.TerraformSuccess != nil || step.TerraformFail != nil {
		state.expectedResult = step.TerraformSuccess != nil
		setupTerraformFilesForStep(state, &step)
		defer cleanupTerraformFiles(state)

		if !state.initDone {
			if _, state.initDone = runCommand(state.t, state.runDir, "terraform init"); !state.initDone {
				state.t.Errorf("Error: Unable to perform manditory 'terraform init' in directory '%v'", state.runDir)
				state.result = false
				return
			}
		}

		_, runResult := runCommand(state.t, state.runDir, fmt.Sprintf("terraform %v", state.terraformOp))
		state.result = state.result && (runResult == state.expectedResult)
		if step.TerraformSuccess != nil && state.result {
			state.successfulApplyStep = &step
		}

		if !state.result {
			state.t.Errorf("Error: Unexpected result error(%v) from terraform with config %v", runResult, *state.lastConfigFiles)
			return
		}
	}
	if step.StateConfigCheck != nil && !state.planOnly {
		doStateChecks(state.t, state.runDir, *step.StateConfigCheck)
	}
	if step.StateConfigCheckFile != nil && !state.planOnly {
		doFileStateChecks(state.t, state.runDir, *step.StateConfigCheckFile)
	}
}

func setupTerraformFilesForStep(state *executionState, step *configStep) (err error) {
	if step.TerraformSuccess != nil {
		state.lastConfigFiles = step.TerraformSuccess
	} else {
		state.lastConfigFiles = step.TerraformFail
	}

	for _, fileName := range *state.lastConfigFiles {
		_, fnd := state.configFileCache[fileName]
		if !fnd {
			filePath := filepath.Join(state.configDir, fileName)
			if state.configFileCache[fileName], err = ioutil.ReadFile(filePath); err != nil {
				state.t.Errorf("Error: Unable to read terraform config file %v, err: %v", filePath, err)
				state.result = false
				return
			}
		}
		if err = ioutil.WriteFile(filepath.Join(state.runDir, fileName), state.configFileCache[fileName], os.ModePerm); err != nil {
			state.result = false
			return
		}
	}
	return
}

func cleanupTerraformFiles(state *executionState) {
	for _, fileName := range *state.lastConfigFiles {
		os.Remove(filepath.Join(state.runDir, fileName))
	}
}

func parseStateString(terraformState string) configChecks {
	// Parse out the 2-level data structure
	outputLines := strings.Split(terraformState, "\n")
	resetEscapeRE := regexp.MustCompile("\x1b\\[0m")
	resourceRE := regexp.MustCompile("^(.*):$")
	keyvalRE := regexp.MustCompile("^\\s*(.*) = (.*)$")
	stateInfo := make(configChecks)
	var currentResource map[string]string
	for _, line := range outputLines {
		line = resetEscapeRE.ReplaceAllLiteralString(line, "")
		if matches := resourceRE.FindStringSubmatch(line); matches != nil {
			currentResource = make(map[string]string)
			stateInfo[matches[1]] = currentResource
			continue
		}
		if matches := keyvalRE.FindStringSubmatch(line); matches != nil {
			currentResource[matches[1]] = matches[2]
		}
	}
	if len(stateInfo) == 0 {
		return nil
	}
	return stateInfo
}

func doFileStateChecks(t *testing.T, runDir, checkFile string) bool {
	content, err := ioutil.ReadFile(checkFile)
	if err != nil {
		t.Errorf("Error: Unable to read from expected state file '%v': %v", checkFile, err)
		return false
	}
	checks := parseStateString(string(content))
	if nil == checks {
		t.Errorf("Error: Unable to parse expected state file '%v'", checkFile)
		return false
	}
	return doStateChecks(t, runDir, checks)
}

func doStateChecks(t *testing.T, runDir string, checks configChecks) bool {
	env := make([]string, len(examplesTestAllowedEnvironmentVariables))
	for index, variable := range examplesTestAllowedEnvironmentVariables {
		env[index] = fmt.Sprintf("%v=%v", variable, os.Getenv(variable))
	}

	cmd := exec.Command("sh", "-c", "terraform show")
	cmd.Dir = runDir
	cmd.Env = env
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: Command Failed. Output: %s", output)
		t.Errorf("Error running 'terraform show' at %v: %v", runDir, err)
		return false
	}

	currentState := parseStateString(string(output))
	result := true
	if len(checks) != len(currentState) {
		t.Errorf("Error: Resource counts do not match; expected %d, actual %d", len(checks), len(currentState))
		return false
	}
	for resourceName, attributeChecks := range checks {
		if currentAttributes, fnd := currentState[resourceName]; fnd {
			if len(attributeChecks) != len(currentAttributes) {
				t.Errorf("Error: Attribute counts on resource %v do not match; expected %d, actual %d", resourceName, len(attributeChecks), len(currentAttributes))
				result = false
			}
			for checkKey, checkVal := range attributeChecks {
				if stateVal, fnd := currentAttributes[checkKey]; fnd {
					if checkVal != stateKeyExistsFlag && checkVal != stateVal {
						t.Errorf("Error: resource %v key %v value %v does not match %v", resourceName, checkKey, checkVal, stateVal)
						result = false
					}
				} else {
					t.Errorf("Error: resource %v missing/mismatch %v in actual state", resourceName, checkKey)
					result = false
				}
			}
		} else {
			t.Errorf("Error: resource %v missing in actual state", resourceName)
			result = false
		}
	}
	return result
}

func runCommand(t *testing.T, path, command string) ([]byte, bool) {
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
		return nil, false
	}

	return output, true
}

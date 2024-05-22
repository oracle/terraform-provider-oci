package resourcediscovery

import (
	"context"
	"fmt"
	"os"
	"testing"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/stretchr/testify/assert"

	hcinstall "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/src"
)

func TestUnitResourceSupportImport(t *testing.T) {
	type expected struct {
		isSupport bool
		gotError  error
	}
	tests := []struct {
		name         string
		resourceName string
		expected     expected
	}{
		{
			name:     "Test response with no resource name",
			expected: expected{isSupport: false, gotError: nil},
		},
		{
			name:         "Test response with valid resource name",
			resourceName: "oci_load_balancer",
			expected:     expected{isSupport: true, gotError: nil},
		},
		{
			name:         "Test response with valid resource name",
			resourceName: "oci_load_balancer.load_balancer",
			expected:     expected{isSupport: true, gotError: nil},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		res, _ := isResourceSupportImportVar(test.resourceName)
		if res != test.expected.isSupport {
			t.Errorf("Output %v not equal to expected %v", res, test.expected.isSupport)
		}

	}
}

func TestUnitExportCompartment(t *testing.T) {
	type expected struct {
		isSupport bool
		gotError  error
	}

	var exportCommandArgs tf_export.ExportCommandArgs
	var compartmentId = ""

	tests := []struct {
		name         string
		resourceName string
		mockFunc     func()
		expected     expected
	}{

		{
			name: "Test response with no resource name",
			mockFunc: func() {
				hcInstallerEnsure = func(installer *hcinstall.Installer, ctx context.Context, sources []src.Source) (string, error) {
					return "", nil
				}
				tfInitVar = func(tf *tfexec.Terraform, initArgs []tfexec.InitOption) error {
					return nil
				}
				tfPlanVar = func(tf *tfexec.Terraform, planArgs []tfexec.PlanOption) (bool, error) {
					return false, nil
				}
				newTerraformVar = func(workingDir string, execPath string) (*tfexec.Terraform, error) {
					return nil, nil
				}
				RunExportCommandVar = func(args *tf_export.ExportCommandArgs) (err error, status Status) {
					return nil, 404
				}

			},
			expected: expected{isSupport: false, gotError: nil},
		},
		{
			name: "Test response with new terraform struct",
			mockFunc: func() {
				hcInstallerEnsure = func(installer *hcinstall.Installer, ctx context.Context, sources []src.Source) (string, error) {
					return "", nil
				}
				tfInitVar = func(tf *tfexec.Terraform, initArgs []tfexec.InitOption) error {
					return nil
				}
				tfPlanVar = func(tf *tfexec.Terraform, planArgs []tfexec.PlanOption) (bool, error) {
					return false, nil
				}
				newTerraformVar = func(workingDir string, execPath string) (*tfexec.Terraform, error) {
					return nil, fmt.Errorf("Error during new terraform")
				}
				RunExportCommandVar = func(args *tf_export.ExportCommandArgs) (err error, status Status) {
					return nil, 404
				}

			},
			expected: expected{isSupport: false, gotError: fmt.Errorf("Error during new terraform")},
		},
		{
			name: "Test response with no resource name, empty provider bin and go path",
			mockFunc: func() {

				hcInstallerEnsure = func(installer *hcinstall.Installer, ctx context.Context, sources []src.Source) (string, error) {
					return "", nil
				}
				tfInitVar = func(tf *tfexec.Terraform, initArgs []tfexec.InitOption) error {
					return nil
				}
				tfPlanVar = func(tf *tfexec.Terraform, planArgs []tfexec.PlanOption) (bool, error) {
					return false, nil
				}
				newTerraformVar = func(workingDir string, execPath string) (*tfexec.Terraform, error) {
					return nil, nil
				}
				RunExportCommandVar = func(args *tf_export.ExportCommandArgs) (err error, status Status) {
					return nil, 404
				}
				getEnvSettingWithBlankDefaultVar = func(path string) string {
					return ""
				}
				utils.GetEnvSettingWithDefault("export_parallelism", "0")

			},
			expected: expected{isSupport: false, gotError: nil},
		},
		{
			name: "Test response with no resource name, empty go path",
			mockFunc: func() {

				hcInstallerEnsure = func(installer *hcinstall.Installer, ctx context.Context, sources []src.Source) (string, error) {
					return "", nil
				}
				tfInitVar = func(tf *tfexec.Terraform, initArgs []tfexec.InitOption) error {
					return nil
				}
				tfPlanVar = func(tf *tfexec.Terraform, planArgs []tfexec.PlanOption) (bool, error) {
					return false, nil
				}
				newTerraformVar = func(workingDir string, execPath string) (*tfexec.Terraform, error) {
					return nil, nil
				}
				RunExportCommandVar = func(args *tf_export.ExportCommandArgs) (err error, status Status) {
					return nil, 404
				}
				os.Setenv("GOPATH", "")
			},
			expected: expected{isSupport: false, gotError: fmt.Errorf("not able to set 'provider_bin_path', either specificy 'provider_bin_path' env variable or set GOPATH to use default provider bin path ($GOPATH/bin)")},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		test.mockFunc()
		res := TestExportCompartment(&compartmentId, &exportCommandArgs)
		if test.name != "Test response with no resource name, empty go path" && test.name != "Test response with new terraform struct" && res != nil {
			t.Errorf("Output %v not equal to expected %v", res, test.expected)
		} else if test.name == "Test response with new terraform struct" {
			if res.Error() != test.expected.gotError.Error() {
				t.Errorf("Output %v not equal to expected %v", res, test.expected.gotError)
			}
		} else if test.name == "Test response with no resource name, empty go path" && res != nil {
			if res.Error() != test.expected.gotError.Error() {
				t.Errorf("Output %v not equal to expected %v", res, test.expected.gotError.Error())
			}
		}
	}
}

func TestUnitExportCompartmentWithResourceName(t *testing.T) {
	type expected struct {
		isSupport bool
		gotError  error
	}

	var id = ""
	var compartmentId = ""

	tests := []struct {
		name         string
		resourceName string
		mockFunc     func()
		expected     expected
	}{
		{
			name: "Test response with no resource",
			mockFunc: func() {
				testExportCompartmentVar = func(compartmentId *string, exportCommandArgs *tf_export.ExportCommandArgs) error {
					return nil
				}
				isResourceSupportImportVar = func(string) (support bool, err error) {
					return true, nil
				}
				hcInstallerEnsure = func(installer *hcinstall.Installer, ctx context.Context, sources []src.Source) (string, error) {
					return "", nil
				}
				tfInitVar = func(tf *tfexec.Terraform, initArgs []tfexec.InitOption) error {
					return nil
				}
				RunExportCommandVar = func(args *tf_export.ExportCommandArgs) (err error, status Status) {
					return nil, 404
				}
				newTerraformVar = func(workingDir string, execPath string) (*tfexec.Terraform, error) {
					return nil, nil
				}
				utils.GetEnvSettingWithDefault("export_parallelism", "0")
				getEnvSettingWithBlankDefaultVar("")
			},
			expected: expected{isSupport: false, gotError: nil},
		},
		{
			name: "Test response with valid resource name",
			mockFunc: func() {
				testExportCompartmentVar = func(compartmentId *string, exportCommandArgs *tf_export.ExportCommandArgs) error {
					return nil
				}
				isResourceSupportImportVar = func(string) (support bool, err error) {
					return true, nil
				}
				hcInstallerEnsure = func(installer *hcinstall.Installer, ctx context.Context, sources []src.Source) (string, error) {
					return "", nil
				}
				tfInitVar = func(tf *tfexec.Terraform, initArgs []tfexec.InitOption) error {
					return nil
				}
				newTerraformVar = func(workingDir string, execPath string) (*tfexec.Terraform, error) {
					return nil, nil
				}
				getEnvSettingWithBlankDefaultVar("")
			},
			resourceName: "oci_load_balancer.load_balancer",
			expected:     expected{isSupport: false, gotError: nil},
		},
		{
			name: "Test validation of hint for resource name",
			mockFunc: func() {
				testExportCompartmentVar = func(compartmentId *string, exportCommandArgs *tf_export.ExportCommandArgs) error {
					return nil
				}
				isResourceSupportImportVar = func(string) (support bool, err error) {
					return true, nil
				}
				hcInstallerEnsure = func(installer *hcinstall.Installer, ctx context.Context, sources []src.Source) (string, error) {
					return "", nil
				}
				tfInitVar = func(tf *tfexec.Terraform, initArgs []tfexec.InitOption) error {
					return nil
				}
				newTerraformVar = func(workingDir string, execPath string) (*tfexec.Terraform, error) {
					return nil, nil
				}
				getEnvSettingWithBlankDefaultVar("")
			},
			resourceName: "oci_core_vcn",
			expected:     expected{isSupport: false, gotError: nil},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		test.mockFunc()
		res := TestExportCompartmentWithResourceName(&id, &compartmentId, test.resourceName)

		if test.expected.gotError == nil {
			assert.NoError(t, res)
		} else {
			if res.Error() != test.expected.gotError.Error() {
				t.Errorf("Output %v not equal to expected %v", res, test.expected.gotError)
			}
		}

	}
}

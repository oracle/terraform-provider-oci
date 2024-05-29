package resourcediscovery

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"

	"github.com/hashicorp/terraform-exec/tfexec"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	tf_provider "github.com/oracle/terraform-provider-oci/internal/provider"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	hcinstall "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/fs"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/src"
)

var (
	hcInstallerEnsure = func(installer *hcinstall.Installer, ctx context.Context, sources []src.Source) (string, error) {
		return installer.Ensure(ctx, sources)
	}
	testExportCompartmentVar   = TestExportCompartment
	isResourceSupportImportVar = isResourceSupportImport
	newTerraformVar            = tfexec.NewTerraform
	RunExportCommandVar        = RunExportCommand
)

var tfInitVar = func(tf *tfexec.Terraform, initArgs []tfexec.InitOption) error {
	return tf.Init(context.Background(), initArgs...)
}
var tfPlanVar = func(tf *tfexec.Terraform, planArgs []tfexec.PlanOption) (bool, error) {
	return tf.Plan(context.Background(), planArgs...)
}

func TestExportCompartmentWithResourceName(id *string, compartmentId *string, resourceName string) error {

	// add logs for notifying execution
	log.Println()
	log.Printf("-------------------------------- Executing Resource Discovery Sub-Step --------------------------------")
	log.Println()

	defer func() {
		// add logs for notifying execution
		log.Println()
		log.Printf("-------------------------------- Exiting Resource Discovery Sub-Step --------------------------------")
		log.Println()
	}()

	var exportCommandArgs tf_export.ExportCommandArgs
	if strings.Contains(resourceName, ".") {
		resourceName = strings.Split(resourceName, ".")[0]
	}

	var err error
	exportCommandArgs.GenerateState, err = isResourceSupportImportVar(resourceName)
	if err != nil {
		return err
	}

	for serviceName, resourceGraph := range tf_export.TenancyResourceGraphs {
		for _, association := range resourceGraph {
			for _, hint := range association {
				if hint.ResourceClass == resourceName {
					exportCommandArgs.Services = []string{serviceName}
					exportCommandArgs.IDs = []string{*id}
					return testExportCompartmentVar(compartmentId, &exportCommandArgs)
				}
			}
		}
	}

	for serviceName, resourceGraph := range tf_export.CompartmentResourceGraphs {
		for _, association := range resourceGraph {
			for _, hint := range association {
				if hint.ResourceClass == resourceName {
					exportCommandArgs.Services = []string{serviceName}
					exportCommandArgs.IDs = []string{*id}
					return testExportCompartmentVar(compartmentId, &exportCommandArgs)
				}
			}
		}
	}

	// compartment export not support yet
	log.Printf("[INFO] ===> Compartment export doesn't support this resource %v yet", resourceName)
	return nil
}

func TestExportCompartment(compartmentId *string, exportCommandArgs *tf_export.ExportCommandArgs) error {
	// checking for provider_bin_path here because parent func will also be
	// called for resources that do not support RD
	if providerBinPath := getEnvSettingWithBlankDefaultVar("provider_bin_path"); providerBinPath == "" {
		goPath := os.Getenv("GOPATH")
		if goPath == "" {
			return fmt.Errorf("not able to set 'provider_bin_path', either specificy 'provider_bin_path' env variable or set GOPATH to use default provider bin path ($GOPATH/bin)")
		}
		if err := os.Setenv("provider_bin_path", strings.Join([]string{os.Getenv("GOPATH"), string(os.PathSeparator), "bin"}, "")); err != nil {
			log.Printf("unable to set 'provider_bin_path' to GOPATH/bin")
			return err
		}
		log.Printf("'provider_bin_path' not provided for resource discovery testing, using GOPATH/bin as default provider location")
	}
	dir, _ := os.Getwd()
	outputDir := fmt.Sprintf(dir + "/exportCompartment")
	if err := os.RemoveAll(outputDir); err != nil {
		log.Printf("unable to remove existing '%s' due to error '%v'", outputDir, err)
		return err
	}
	if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
		log.Printf("unable to Create '%s' due to error '%v'", outputDir, err)
		return err
	}
	defer func() {
		if err := os.RemoveAll(outputDir); err != nil {
			log.Printf("unable to cleanup '%s' due to error '%v'", outputDir, err)
		}
	}()
	exportCommandArgs.Services = append(exportCommandArgs.Services, "availability_domain")
	exportCommandArgs.CompartmentId = compartmentId
	exportCommandArgs.OutputDir = &outputDir
	var tfVersion tf_export.TfHclVersion = &tf_export.TfHclVersion12{Value: tf_export.TfVersion12}
	exportCommandArgs.TFVersion = &tfVersion

	var parseErr error
	if exportCommandArgs.Parallelism, parseErr = strconv.Atoi(utils.GetEnvSettingWithDefault("export_parallelism", "10")); parseErr != nil {
		return fmt.Errorf("[ERROR] invalid value for resource discovery parallelism: %s", parseErr.Error())
	}
	log.Printf("[INFO] exportCommandArgs.Parallelism: %d", exportCommandArgs.Parallelism)

	if errExport, status := RunExportCommandVar(exportCommandArgs); errExport != nil || status == StatusPartialSuccess {
		if errExport != nil {
			return fmt.Errorf("[ERROR] RunExportCommand failed: %s", errExport.Error())
		}
		// For generated tests, RD will only return this error if one of the `ids` was not found
		// (which in case of tests is the id for the resource RD is looking for)
		if errExport != nil && status == StatusPartialSuccess {
			var idsNotFoundError tf_export.ResourceDiscoveryCustomError
			idsNotFoundError = tf_export.ResourceDiscoveryCustomError{
				TypeOfError: tf_export.PartiallyResourcesDiscoveredError,
				Message:     errExport,
				Suggestion:  tf_export.PartiallyResourcesDiscoveredSuggestion,
			}
			return idsNotFoundError.Error()
		}
	}

	// run init command

	terraformBinPath := getEnvSettingWithBlankDefaultVar(globalvar.TerraformBinPathName)
	if terraformBinPath == "" {
		var err error
		terraformBinPath, err = hcInstallerEnsure(hcinstall.NewInstaller(), context.Background(),
			[]src.Source{src.Findable(&fs.AnyVersion{Product: &product.Terraform})})
		if err != nil {
			return err
		}
	}
	tf, err := newTerraformVar(*exportCommandArgs.OutputDir, terraformBinPath)
	if err != nil {
		return err
	}
	//backgroundCtx := context.Background()

	var initArgs []tfexec.InitOption
	if pluginDir := getEnvSettingWithBlankDefaultVar("provider_bin_path"); pluginDir != "" {
		log.Printf("[INFO] plugin dir: '%s'", pluginDir)
		initArgs = append(initArgs, tfexec.PluginDir(pluginDir))
	}
	if err := tfInitVar(tf, initArgs); err != nil {
		return err
	}

	// Need to set the compartment id environment variable for plan step
	compartmentOcidVarName := "TF_VAR_compartment_ocid"
	storeCompartmentId := os.Getenv(compartmentOcidVarName)
	if err := os.Setenv(compartmentOcidVarName, *compartmentId); err != nil {
		return fmt.Errorf("could not set %s environment in export test", compartmentOcidVarName)
	}

	defer func() {
		if storeCompartmentId != "" {
			if err := os.Setenv(compartmentOcidVarName, storeCompartmentId); err != nil {
				log.Printf("[WARN] unable to restore %s to %s", compartmentOcidVarName, storeCompartmentId)
			}
		}
	}()

	// run plan command

	var planArgs []tfexec.PlanOption
	if exportCommandArgs.GenerateState {
		statefile := fmt.Sprintf(*exportCommandArgs.OutputDir + "/terraform.tfstate")
		planArgs = append(planArgs, tfexec.State(statefile))
	}

	if _, err := tfPlanVar(tf, planArgs); err != nil {
		return fmt.Errorf("[ERROR] terraform plan command failed %s", err.Error())
	}
	return nil
}

func isResourceSupportImport(resourceName string) (support bool, err error) {
	if strings.Contains(resourceName, ".") {
		resourceName = strings.Split(resourceName, ".")[0]
	}
	resource := tf_provider.ResourcesMap()[resourceName]
	if resource == nil {
		return false, fmt.Errorf("[ERROR]: resource %v is not found in resource Map", resourceName)
	}
	return resource.Importer != nil, nil
}

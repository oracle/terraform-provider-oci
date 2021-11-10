package resourcediscovery

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/hashicorp/terraform-exec/tfinstall"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"
	tf_provider "github.com/terraform-providers/terraform-provider-oci/internal/provider"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

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

	var exportCommandArgs ExportCommandArgs
	if strings.Contains(resourceName, ".") {
		resourceName = strings.Split(resourceName, ".")[0]
	}

	var err error
	exportCommandArgs.GenerateState, err = isResourceSupportImport(resourceName)
	if err != nil {
		return err
	}

	for serviceName, resourceGraph := range tenancyResourceGraphs {
		for _, association := range resourceGraph {
			for _, hint := range association {
				if hint.resourceClass == resourceName {
					exportCommandArgs.Services = []string{serviceName}
					exportCommandArgs.IDs = []string{*id}
					return testExportCompartment(compartmentId, &exportCommandArgs)
				}
			}
		}
	}

	for serviceName, resourceGraph := range compartmentResourceGraphs {
		for _, association := range resourceGraph {
			for _, hint := range association {
				if hint.resourceClass == resourceName {
					exportCommandArgs.Services = []string{serviceName}
					exportCommandArgs.IDs = []string{*id}
					return testExportCompartment(compartmentId, &exportCommandArgs)
				}
			}
		}
	}

	// compartment export not support yet
	log.Printf("[INFO] ===> Compartment export doesn't support this resource %v yet", resourceName)
	return nil
}
func testExportCompartment(compartmentId *string, exportCommandArgs *ExportCommandArgs) error {
	// checking for provider_bin_path here because parent func will also be
	// called for resources that do not support RD
	if providerBinPath := utils.GetEnvSettingWithBlankDefault("provider_bin_path"); providerBinPath == "" {
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
	var tfVersion TfHclVersion = &TfHclVersion12{Value: TfVersion12}
	exportCommandArgs.TFVersion = &tfVersion

	var parseErr error
	if exportCommandArgs.Parallelism, parseErr = strconv.Atoi(utils.GetEnvSettingWithDefault("export_parallelism", "10")); parseErr != nil {
		return fmt.Errorf("[ERROR] invalid value for resource discovery parallelism: %s", parseErr.Error())
	}
	log.Printf("[INFO] exportCommandArgs.Parallelism: %d", exportCommandArgs.Parallelism)

	if errExport, status := RunExportCommand(exportCommandArgs); errExport != nil || status == StatusPartialSuccess {
		if errExport != nil {
			return fmt.Errorf("[ERROR] RunExportCommand failed: %s", errExport.Error())
		}
		// For generated tests, RD will only return this error if one of the `ids` was not found
		// (which in case of tests is the id for the resource RD is looking for)
		if status == StatusPartialSuccess {
			return fmt.Errorf("[ERROR] expected resource was not found")
		}
	}

	// run init command

	terraformBinPath := utils.GetEnvSettingWithBlankDefault(globalvar.TerraformBinPathName)
	if terraformBinPath == "" {
		var err error
		terraformBinPath, err = tfinstall.Find(context.Background(), tfinstall.LookPath())
		if err != nil {
			return err
		}
	}
	tf, err := tfexec.NewTerraform(*exportCommandArgs.OutputDir, terraformBinPath)
	if err != nil {
		return err
	}
	backgroundCtx := context.Background()

	var initArgs []tfexec.InitOption
	if pluginDir := utils.GetEnvSettingWithBlankDefault("provider_bin_path"); pluginDir != "" {
		log.Printf("[INFO] plugin dir: '%s'", pluginDir)
		initArgs = append(initArgs, tfexec.PluginDir(pluginDir))
	}
	if err := tf.Init(backgroundCtx, initArgs...); err != nil {
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

	if _, err := tf.Plan(backgroundCtx, planArgs...); err != nil {
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
		return false, fmt.Errorf("[ERROR]: resouce %v is not found in resource Map", resourceName)
	}
	return resource.Importer != nil, nil
}

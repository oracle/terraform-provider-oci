package apm_config

import (
	"fmt"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportApmConfigConfigHints.GetIdFn = getApmConfigConfigId
	exportApmConfigDataFileHints.GetIdFn = getApmConfigDataFileId
	tf_export.RegisterCompartmentGraphs("apm_config", apmConfigResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getApmConfigConfigId(resource *tf_export.OCIResource) (string, error) {

	configId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find configId for ApmConfig Config")
	}
	apmDomainId, ok := resource.SourceAttributes["apm_domain_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find apmDomainId for ApmConfig Config")
	}
	return GetConfigCompositeId(configId, apmDomainId), nil
}

func getApmConfigDataFileId(resource *tf_export.OCIResource) (string, error) {

	apmDomainId, ok := resource.SourceAttributes["apm_domain_id"].(string)
	apmType, ok := resource.SourceAttributes["apm_type"].(string)
	dataFileName, ok := resource.SourceAttributes["data_file_name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataFileName for ApmConfig DataFile")
	}
	return GetDataFileCompositeId(apmDomainId, apmType, dataFileName), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportApmConfigConfigHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apm_config_config",
	DatasourceClass:        "oci_apm_config_configs",
	DatasourceItemsAttr:    "config_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "config",
	RequireResourceRefresh: true,
}

var exportApmConfigDataFileHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apm_config_data_file",
	DatasourceClass:        "oci_apm_config_data_files",
	DatasourceItemsAttr:    "data_file_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "data_file",
}

var apmConfigResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {},
}

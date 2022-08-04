package license_manager

import (
	"fmt"
	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportLicenseManagerConfigurationHints.GetIdFn = getLicenseManagerConfigurationId
	tf_export.RegisterCompartmentGraphs("license_manager", licenseManagerResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getLicenseManagerConfigurationId(resource *tf_export.OCIResource) (string, error) {
	compartmentId, ok := resource.SourceAttributes["compartment_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find compartment_id for Licensemanager Configuration")
	}
	return GetConfigurationId(compartmentId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportLicenseManagerConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_license_manager_configuration",
	DatasourceClass:      "oci_license_manager_configuration",
	ResourceAbbreviation: "configuration",
}

var exportLicenseManagerProductLicenseHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_license_manager_product_license",
	DatasourceClass:        "oci_license_manager_product_licenses",
	DatasourceItemsAttr:    "product_license_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "product_license",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_license_manager.LifeCycleStateActive),
	},
}

var exportLicenseManagerLicenseRecordHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_license_manager_license_record",
	DatasourceClass:        "oci_license_manager_license_records",
	DatasourceItemsAttr:    "license_record_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "license_record",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_license_manager.LifeCycleStateActive),
	},
}

var licenseManagerResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportLicenseManagerConfigurationHints},
		{TerraformResourceHints: exportLicenseManagerProductLicenseHints},
	},
	"oci_license_manager_product_license": {
		{
			TerraformResourceHints: exportLicenseManagerLicenseRecordHints,
			DatasourceQueryParams: map[string]string{
				"product_license_id": "id",
			},
		},
	},
}

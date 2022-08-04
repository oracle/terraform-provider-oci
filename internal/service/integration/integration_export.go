package integration

import (
	oci_integration "github.com/oracle/oci-go-sdk/v65/integration"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("integration", integrationResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportIntegrationIntegrationInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_integration_integration_instance",
	DatasourceClass:      "oci_integration_integration_instances",
	DatasourceItemsAttr:  "integration_instances",
	ResourceAbbreviation: "integration_instance",
	DiscoverableLifecycleStates: []string{
		string(oci_integration.IntegrationInstanceLifecycleStateActive),
	},
}

var integrationResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportIntegrationIntegrationInstanceHints},
	},
}

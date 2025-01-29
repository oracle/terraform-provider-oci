package tenantmanagercontrolplane

import (
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("tenantmanagercontrolplane", tenantmanagercontrolplaneResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportTenantmanagercontrolplaneSubscriptionMappingHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_tenantmanagercontrolplane_subscription_mapping",
	DatasourceClass:        "oci_tenantmanagercontrolplane_subscription_mappings",
	DatasourceItemsAttr:    "subscription_mapping_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "subscription_mapping",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_tenantmanagercontrolplane.SubscriptionMappingLifecycleStateActive),
	},
}

var tenantmanagercontrolplaneResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportTenantmanagercontrolplaneSubscriptionMappingHints},
	},
}

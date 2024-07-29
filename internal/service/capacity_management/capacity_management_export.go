package capacity_management

import (
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("capacity_management", capacityManagementResourceGraph)
}

// Hints for discovering and exporting this resource to configuration and state files
var exportCapacityManagementOccCustomerGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_occ_customer_group",
	DatasourceClass:        "oci_capacity_management_occ_customer_groups",
	DatasourceItemsAttr:    "occ_customer_group_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "occ_customer_group",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.OccCustomerGroupLifecycleStateActive),
	},
}

var exportCapacityManagementOccAvailabilityCatalogHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_occ_availability_catalog",
	DatasourceClass:        "oci_capacity_management_occ_availability_catalogs",
	DatasourceItemsAttr:    "occ_availability_catalog_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "occ_availability_catalog",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.OccAvailabilityCatalogLifecycleStateActive),
	},
}

var exportCapacityManagementOccCapacityRequestHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_occ_capacity_request",
	DatasourceClass:        "oci_capacity_management_occ_capacity_requests",
	DatasourceItemsAttr:    "occ_capacity_request_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "occ_capacity_request",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.OccCapacityRequestLifecycleStateActive),
	},
}

var exportCapacityManagementInternalOccCapacityRequestHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_internal_occ_capacity_request",
	DatasourceClass:        "oci_capacity_management_internal_occ_capacity_requests",
	DatasourceItemsAttr:    "occ_capacity_request_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "internal_occ_capacity_request",
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.OccCapacityRequestLifecycleStateActive),
	},
}

var capacityManagementResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCapacityManagementOccCustomerGroupHints},
		{TerraformResourceHints: exportCapacityManagementInternalOccCapacityRequestHints},
		{TerraformResourceHints: exportCapacityManagementOccAvailabilityCatalogHints},
		{TerraformResourceHints: exportCapacityManagementOccCapacityRequestHints},
	},
}

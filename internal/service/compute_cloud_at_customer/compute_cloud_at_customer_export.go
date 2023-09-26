package compute_cloud_at_customer

import (
	oci_compute_cloud_at_customer "github.com/oracle/oci-go-sdk/v65/computecloudatcustomer"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("compute_cloud_at_customer", computeCloudAtCustomerResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportComputeCloudAtCustomerCccUpgradeScheduleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_compute_cloud_at_customer_ccc_upgrade_schedule",
	DatasourceClass:        "oci_compute_cloud_at_customer_ccc_upgrade_schedules",
	DatasourceItemsAttr:    "ccc_upgrade_schedule_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "ccc_upgrade_schedule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_compute_cloud_at_customer.CccUpgradeScheduleLifecycleStateActive),
		string(oci_compute_cloud_at_customer.CccUpgradeScheduleLifecycleStateNeedsAttention),
	},
}

var exportComputeCloudAtCustomerCccInfrastructureHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_compute_cloud_at_customer_ccc_infrastructure",
	DatasourceClass:        "oci_compute_cloud_at_customer_ccc_infrastructures",
	DatasourceItemsAttr:    "ccc_infrastructure_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "ccc_infrastructure",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_compute_cloud_at_customer.CccInfrastructureLifecycleStateActive),
		string(oci_compute_cloud_at_customer.CccInfrastructureLifecycleStateNeedsAttention),
	},
}

var computeCloudAtCustomerResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportComputeCloudAtCustomerCccUpgradeScheduleHints},
		{TerraformResourceHints: exportComputeCloudAtCustomerCccInfrastructureHints},
	},
}

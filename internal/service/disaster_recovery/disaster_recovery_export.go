package disaster_recovery

import (
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("disaster_recovery", disasterRecoveryResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDisasterRecoveryDrProtectionGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_disaster_recovery_dr_protection_group",
	DatasourceClass:        "oci_disaster_recovery_dr_protection_groups",
	DatasourceItemsAttr:    "dr_protection_group_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "dr_protection_group",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_disaster_recovery.DrProtectionGroupLifecycleStateActive),
		string(oci_disaster_recovery.DrProtectionGroupLifecycleStateNeedsAttention),
	},
}

var exportDisasterRecoveryDrPlanExecutionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_disaster_recovery_dr_plan_execution",
	DatasourceClass:        "oci_disaster_recovery_dr_plan_executions",
	DatasourceItemsAttr:    "dr_plan_execution_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "dr_plan_execution",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_disaster_recovery.DrPlanExecutionLifecycleStateSucceeded),
	},
}

var exportDisasterRecoveryDrPlanHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_disaster_recovery_dr_plan",
	DatasourceClass:        "oci_disaster_recovery_dr_plans",
	DatasourceItemsAttr:    "dr_plan_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "dr_plan",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_disaster_recovery.DrPlanLifecycleStateActive),
		string(oci_disaster_recovery.DrPlanLifecycleStateNeedsAttention),
	},
}

var disasterRecoveryResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDisasterRecoveryDrProtectionGroupHints},
	},
	"oci_disaster_recovery_dr_protection_group": {
		{
			TerraformResourceHints: exportDisasterRecoveryDrPlanHints,
			DatasourceQueryParams: map[string]string{
				"dr_protection_group_id": "id",
			},
		},
		{
			TerraformResourceHints: exportDisasterRecoveryDrPlanExecutionHints,
			DatasourceQueryParams: map[string]string{
				"dr_protection_group_id": "id",
			},
		},
	},
}

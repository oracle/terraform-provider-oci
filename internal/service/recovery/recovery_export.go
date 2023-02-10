package recovery

import (
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("recovery", recoveryResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportRecoveryRecoveryServiceSubnetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_recovery_recovery_service_subnet",
	DatasourceClass:        "oci_recovery_recovery_service_subnets",
	DatasourceItemsAttr:    "recovery_service_subnet_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "recovery_service_subnet",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_recovery.LifecycleStateActive),
	},
}

var exportRecoveryProtectedDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_recovery_protected_database",
	DatasourceClass:        "oci_recovery_protected_databases",
	DatasourceItemsAttr:    "protected_database_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "protected_database",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_recovery.LifecycleStateActive),
	},
}

var exportRecoveryProtectionPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_recovery_protection_policy",
	DatasourceClass:        "oci_recovery_protection_policies",
	DatasourceItemsAttr:    "protection_policy_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "protection_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_recovery.LifecycleStateActive),
	},
}

var recoveryResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportRecoveryRecoveryServiceSubnetHints},
		{TerraformResourceHints: exportRecoveryProtectedDatabaseHints},
		{TerraformResourceHints: exportRecoveryProtectionPolicyHints},
	},
}

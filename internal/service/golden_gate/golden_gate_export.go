package golden_gate

import (
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("golden_gate", goldenGateResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportGoldenGateDatabaseRegistrationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_golden_gate_database_registration",
	DatasourceClass:        "oci_golden_gate_database_registrations",
	DatasourceItemsAttr:    "database_registration_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_registration",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_golden_gate.LifecycleStateActive),
		string(oci_golden_gate.LifecycleStateNeedsAttention),
		string(oci_golden_gate.LifecycleStateSucceeded),
	},
}

var exportGoldenGateDeploymentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_golden_gate_deployment",
	DatasourceClass:        "oci_golden_gate_deployments",
	DatasourceItemsAttr:    "deployment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "deployment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_golden_gate.LifecycleStateActive),
		string(oci_golden_gate.LifecycleStateNeedsAttention),
		string(oci_golden_gate.LifecycleStateSucceeded),
	},
}

var exportGoldenGateDeploymentBackupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_golden_gate_deployment_backup",
	DatasourceClass:        "oci_golden_gate_deployment_backups",
	DatasourceItemsAttr:    "deployment_backup_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "deployment_backup",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_golden_gate.LifecycleStateActive),
		string(oci_golden_gate.LifecycleStateNeedsAttention),
		string(oci_golden_gate.LifecycleStateSucceeded),
	},
}

var goldenGateResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportGoldenGateDatabaseRegistrationHints},
		{TerraformResourceHints: exportGoldenGateDeploymentHints},
		{TerraformResourceHints: exportGoldenGateDeploymentBackupHints},
	},
}

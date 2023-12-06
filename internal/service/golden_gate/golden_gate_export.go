package golden_gate

import (
	"fmt"

	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportGoldenGateDeploymentCertificateHints.GetIdFn = getGoldenGateDeploymentCertificateId
	tf_export.RegisterCompartmentGraphs("golden_gate", goldenGateResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getGoldenGateDeploymentCertificateId(resource *tf_export.OCIResource) (string, error) {

	certificateKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find certificateKey for GoldenGate DeploymentCertificate")
	}
	deploymentId := resource.Parent.Id
	return GetDeploymentCertificateCompositeId(certificateKey, deploymentId), nil
}

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

var exportGoldenGateConnectionAssignmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_golden_gate_connection_assignment",
	DatasourceClass:        "oci_golden_gate_connection_assignments",
	DatasourceItemsAttr:    "connection_assignment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "connection_assignment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_golden_gate.ConnectionAssignmentLifecycleStateActive),
	},
}

var exportGoldenGateConnectionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_golden_gate_connection",
	DatasourceClass:        "oci_golden_gate_connections",
	DatasourceItemsAttr:    "connection_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "connection",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_golden_gate.ConnectionLifecycleStateActive),
	},
}

var exportGoldenGateDeploymentCertificateHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_golden_gate_deployment_certificate",
	DatasourceClass:        "oci_golden_gate_deployment_certificates",
	DatasourceItemsAttr:    "certificate_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "deployment_certificate",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_golden_gate.CertificateLifecycleStateActive),
	},
}

var goldenGateResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportGoldenGateDatabaseRegistrationHints},
		{TerraformResourceHints: exportGoldenGateDeploymentHints},
		{TerraformResourceHints: exportGoldenGateDeploymentBackupHints},
		{TerraformResourceHints: exportGoldenGateConnectionAssignmentHints},
		{TerraformResourceHints: exportGoldenGateConnectionHints},
	},
	"oci_golden_gate_deployment": {
		{
			TerraformResourceHints: exportGoldenGateDeploymentCertificateHints,
			DatasourceQueryParams: map[string]string{
				"deployment_id": "id",
			},
		},
	},
}

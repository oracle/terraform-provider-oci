package operator_access_control

import (
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v65/operatoraccesscontrol"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("operator_access_control", operatorAccessControlResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportOperatorAccessControlOperatorControlHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_operator_access_control_operator_control",
	DatasourceClass:        "oci_operator_access_control_operator_controls",
	DatasourceItemsAttr:    "operator_control_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "operator_control",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_operator_access_control.OperatorControlLifecycleStatesCreated),
		string(oci_operator_access_control.OperatorControlLifecycleStatesAssigned),
	},
}

var exportOperatorAccessControlOperatorControlAssignmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_operator_access_control_operator_control_assignment",
	DatasourceClass:        "oci_operator_access_control_operator_control_assignments",
	DatasourceItemsAttr:    "operator_control_assignment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "operator_control_assignment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_operator_access_control.OperatorControlAssignmentLifecycleStatesCreated),
		string(oci_operator_access_control.OperatorControlAssignmentLifecycleStatesApplied),
		string(oci_operator_access_control.OperatorControlAssignmentLifecycleStatesApplyfailed),
	},
}

var operatorAccessControlResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOperatorAccessControlOperatorControlHints},
		{TerraformResourceHints: exportOperatorAccessControlOperatorControlAssignmentHints},
	},
}

package datacc

import (
	oci_datacc "github.com/oracle/oci-go-sdk/v65/datacc"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("datacc", dataccResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDataccVmInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datacc_vm_instance",
	DatasourceClass:        "oci_datacc_vm_instances",
	DatasourceItemsAttr:    "vm_instance_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "vm_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datacc.VmInstanceLifecycleStateActive),
	},
}

var exportDataccVmClusterNetworkHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datacc_vm_cluster_network",
	DatasourceClass:        "oci_datacc_vm_cluster_networks",
	DatasourceItemsAttr:    "vm_cluster_network_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "vm_cluster_network",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datacc.VmClusterNetworkLifecycleStateRequiresValidation),
		string(oci_datacc.VmClusterNetworkLifecycleStateValidated),
	},
}

var exportDataccInfrastructureHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datacc_infrastructure",
	DatasourceClass:        "oci_datacc_infrastructures",
	DatasourceItemsAttr:    "infrastructure_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "infrastructure",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datacc.InfrastructureLifecycleStateRequiresValidation),
		string(oci_datacc.InfrastructureLifecycleStateRequiresActivation),
		string(oci_datacc.InfrastructureLifecycleStateActive),
	},
}

var dataccResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataccVmInstanceHints},
		{TerraformResourceHints: exportDataccVmClusterNetworkHints},
		{TerraformResourceHints: exportDataccInfrastructureHints},
	},
}

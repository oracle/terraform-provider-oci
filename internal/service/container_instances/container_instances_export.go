package container_instances

import (
	oci_container_instances "github.com/oracle/oci-go-sdk/v65/containerinstances"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("container_instances", containerInstancesResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportContainerInstancesContainerInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_container_instances_container_instance",
	DatasourceClass:        "oci_container_instances_container_instances",
	DatasourceItemsAttr:    "container_instance_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "container_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_container_instances.ContainerInstanceLifecycleStateActive),
	},
}

var containerInstancesResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportContainerInstancesContainerInstanceHints},
	},
}

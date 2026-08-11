package ddfs

import (
	oci_ddfs "github.com/oracle/oci-go-sdk/v65/ddfs"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("ddfs", ddfsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDdfsInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ddfs_instance",
	DatasourceClass:        "oci_ddfs_instances",
	DatasourceItemsAttr:    "instance_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ddfs.InstanceLifecycleStateActive),
	},
}

var ddfsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDdfsInstanceHints},
	},
}

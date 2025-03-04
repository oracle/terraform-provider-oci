package lustre_file_storage

import (
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("lustre_file_storage", lustreFileStorageResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportLustreFileStorageLustreFileSystemHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_lustre_file_storage_lustre_file_system",
	DatasourceClass:        "oci_lustre_file_storage_lustre_file_systems",
	DatasourceItemsAttr:    "lustre_file_system_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "lustre_file_system",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_lustre_file_storage.LustreFileSystemLifecycleStateActive),
	},
}

var lustreFileStorageResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportLustreFileStorageLustreFileSystemHints},
	},
}

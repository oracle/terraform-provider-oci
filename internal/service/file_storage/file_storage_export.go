package file_storage

import (
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportFileStorageMountTargetHints.RequireResourceRefresh = true
	tf_export.RegisterCompartmentGraphs("file_storage", fileStorageResourceGraph)
	tf_export.BuildAvailabilityResourceGraph("oci_identity_availability_domain", customAssociationFileStorageIdentityAvailabilityDomain)
	tf_export.BuildAvailabilityResourceGraph("oci_file_storage_file_system", customAssociationFileStorageFileSystem)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportFileStorageFileSystemHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_file_storage_file_system",
	DatasourceClass:        "oci_file_storage_file_systems",
	DatasourceItemsAttr:    "file_systems",
	ResourceAbbreviation:   "file_system",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_file_storage.FileSystemLifecycleStateActive),
	},
}

var exportFileStorageMountTargetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_file_storage_mount_target",
	DatasourceClass:        "oci_file_storage_mount_targets",
	DatasourceItemsAttr:    "mount_targets",
	ResourceAbbreviation:   "mount_target",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_file_storage.MountTargetLifecycleStateActive),
	},
}

var exportFileStorageExportHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_file_storage_export",
	DatasourceClass:        "oci_file_storage_exports",
	DatasourceItemsAttr:    "exports",
	ResourceAbbreviation:   "export",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_file_storage.ExportLifecycleStateActive),
	},
}

var exportFileStorageSnapshotHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_file_storage_snapshot",
	DatasourceClass:      "oci_file_storage_snapshots",
	DatasourceItemsAttr:  "snapshots",
	ResourceAbbreviation: "snapshot",
	DiscoverableLifecycleStates: []string{
		string(oci_file_storage.SnapshotLifecycleStateActive),
	},
}

var exportFileStorageReplicationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_file_storage_replication",
	DatasourceClass:        "oci_file_storage_replications",
	DatasourceItemsAttr:    "replications",
	ResourceAbbreviation:   "replication",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_file_storage.ReplicationLifecycleStateActive),
	},
}

var exportFileStorageFilesystemSnapshotPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_file_storage_filesystem_snapshot_policy",
	DatasourceClass:        "oci_file_storage_filesystem_snapshot_policies",
	DatasourceItemsAttr:    "filesystem_snapshot_policies",
	ResourceAbbreviation:   "filesystem_snapshot_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_file_storage.FilesystemSnapshotPolicyLifecycleStateActive),
		string(oci_file_storage.FilesystemSnapshotPolicyLifecycleStateInactive),
	},
}

var exportFileStorageOutboundConnectorHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_file_storage_outbound_connector",
	DatasourceClass:        "oci_file_storage_outbound_connectors",
	DatasourceItemsAttr:    "outbound_connectors",
	ResourceAbbreviation:   "outbound_connector",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_file_storage.OutboundConnectorLifecycleStateActive),
	},
}

var fileStorageResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportFileStorageExportHints},
	},
}

var customAssociationFileStorageIdentityAvailabilityDomain = []tf_export.TerraformResourceAssociation{
	{
		TerraformResourceHints: exportFileStorageFileSystemHints,
		DatasourceQueryParams: map[string]string{
			"availability_domain": "name",
		},
	},
	{
		TerraformResourceHints: exportFileStorageMountTargetHints,
		DatasourceQueryParams: map[string]string{
			"availability_domain": "name",
		},
	},
	{
		TerraformResourceHints: exportFileStorageReplicationHints,
		DatasourceQueryParams: map[string]string{
			"availability_domain": "name",
		},
	},
	{
		TerraformResourceHints: exportFileStorageFilesystemSnapshotPolicyHints,
		DatasourceQueryParams: map[string]string{
			"availability_domain": "name",
		},
	},
	{
		TerraformResourceHints: exportFileStorageOutboundConnectorHints,
		DatasourceQueryParams: map[string]string{
			"availability_domain": "name",
		},
	},
}

var customAssociationFileStorageFileSystem = []tf_export.TerraformResourceAssociation{
	{
		TerraformResourceHints: exportFileStorageSnapshotHints,
		DatasourceQueryParams: map[string]string{
			"file_system_id": "id",
		},
	},
}

// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_file_storage_export", FileStorageExportResource())
	tfresource.RegisterResource("oci_file_storage_export_set", FileStorageExportSetResource())
	tfresource.RegisterResource("oci_file_storage_file_system", FileStorageFileSystemResource())
	tfresource.RegisterResource("oci_file_storage_filesystem_snapshot_policy", FileStorageFilesystemSnapshotPolicyResource())
	tfresource.RegisterResource("oci_file_storage_mount_target", FileStorageMountTargetResource())
	tfresource.RegisterResource("oci_file_storage_replication", FileStorageReplicationResource())
	tfresource.RegisterResource("oci_file_storage_outbound_connector", FileStorageOutboundConnectorResource())
	tfresource.RegisterResource("oci_file_storage_snapshot", FileStorageSnapshotResource())
}

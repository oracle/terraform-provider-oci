// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_file_storage_export_sets", FileStorageExportSetsDataSource())
	tfresource.RegisterDatasource("oci_file_storage_exports", FileStorageExportsDataSource())
	tfresource.RegisterDatasource("oci_file_storage_file_systems", FileStorageFileSystemsDataSource())
	tfresource.RegisterDatasource("oci_file_storage_filesystem_snapshot_policies", FileStorageFilesystemSnapshotPoliciesDataSource())
	tfresource.RegisterDatasource("oci_file_storage_filesystem_snapshot_policy", FileStorageFilesystemSnapshotPolicyDataSource())
	tfresource.RegisterDatasource("oci_file_storage_mount_targets", FileStorageMountTargetsDataSource())
	tfresource.RegisterDatasource("oci_file_storage_replication", FileStorageReplicationDataSource())
	tfresource.RegisterDatasource("oci_file_storage_replication_target", FileStorageReplicationTargetDataSource())
	tfresource.RegisterDatasource("oci_file_storage_replication_targets", FileStorageReplicationTargetsDataSource())
	tfresource.RegisterDatasource("oci_file_storage_replications", FileStorageReplicationsDataSource())
	tfresource.RegisterDatasource("oci_file_storage_outbound_connector", FileStorageOutboundConnectorDataSource())
	tfresource.RegisterDatasource("oci_file_storage_outbound_connectors", FileStorageOutboundConnectorsDataSource())
	tfresource.RegisterDatasource("oci_file_storage_snapshot", FileStorageSnapshotDataSource())
	tfresource.RegisterDatasource("oci_file_storage_snapshots", FileStorageSnapshotsDataSource())
}

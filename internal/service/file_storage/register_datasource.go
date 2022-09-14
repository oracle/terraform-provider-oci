// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_file_storage_export_sets", FileStorageExportSetsDataSource())
	tfresource.RegisterDatasource("oci_file_storage_exports", FileStorageExportsDataSource())
	tfresource.RegisterDatasource("oci_file_storage_file_systems", FileStorageFileSystemsDataSource())
	tfresource.RegisterDatasource("oci_file_storage_mount_targets", FileStorageMountTargetsDataSource())
	tfresource.RegisterDatasource("oci_file_storage_snapshot", FileStorageSnapshotDataSource())
	tfresource.RegisterDatasource("oci_file_storage_snapshots", FileStorageSnapshotsDataSource())
}

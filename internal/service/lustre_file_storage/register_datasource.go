// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package lustre_file_storage

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_lustre_file_storage_lustre_file_system", LustreFileStorageLustreFileSystemDataSource())
	tfresource.RegisterDatasource("oci_lustre_file_storage_lustre_file_systems", LustreFileStorageLustreFileSystemsDataSource())
	tfresource.RegisterDatasource("oci_lustre_file_storage_object_storage_link", LustreFileStorageObjectStorageLinkDataSource())
	tfresource.RegisterDatasource("oci_lustre_file_storage_object_storage_link_sync_job", LustreFileStorageObjectStorageLinkSyncJobDataSource())
	tfresource.RegisterDatasource("oci_lustre_file_storage_object_storage_link_sync_jobs", LustreFileStorageObjectStorageLinkSyncJobsDataSource())
	tfresource.RegisterDatasource("oci_lustre_file_storage_object_storage_links", LustreFileStorageObjectStorageLinksDataSource())
}

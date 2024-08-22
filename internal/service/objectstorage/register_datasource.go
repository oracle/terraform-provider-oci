// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_objectstorage_bucket", ObjectStorageBucketDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_bucket_summaries", ObjectStorageBucketsDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_namespace", ObjectStorageNamespaceDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_object", ObjectStorageObjectDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_object_lifecycle_policy", ObjectStorageObjectLifecyclePolicyDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_object_versions", ObjectStorageObjectVersionsDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_objects", ObjectStorageObjectsDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_preauthrequest", ObjectStoragePreauthenticatedRequestDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_preauthrequests", ObjectStoragePreauthenticatedRequestsDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_replication_policies", ObjectStorageReplicationPoliciesDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_replication_policy", ObjectStorageReplicationPolicyDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_replication_sources", ObjectStorageReplicationSourcesDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_namespace_metadata", ObjectStorageNamespaceMetadataDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_object_head", ObjectStorageObjectHeadDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_private_endpoint", ObjectStoragePrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_objectstorage_private_endpoint_summaries", ObjectStoragePrivateEndpointsDataSource())
}

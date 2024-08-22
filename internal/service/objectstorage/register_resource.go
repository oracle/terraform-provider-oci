// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_objectstorage_bucket", ObjectStorageBucketResource())
	tfresource.RegisterResource("oci_objectstorage_object", ObjectStorageObjectResource())
	tfresource.RegisterResource("oci_objectstorage_object_lifecycle_policy", ObjectStorageObjectLifecyclePolicyResource())
	tfresource.RegisterResource("oci_objectstorage_preauthrequest", ObjectStoragePreauthenticatedRequestResource())
	tfresource.RegisterResource("oci_objectstorage_replication_policy", ObjectStorageReplicationPolicyResource())
	tfresource.RegisterResource("oci_objectstorage_namespace_metadata", ObjectStorageNamespaceMetadataResource())
	tfresource.RegisterResource("oci_objectstorage_private_endpoint", ObjectStoragePrivateEndpointResource())
}

---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_bucket"
sidebar_current: "docs-oci-resource-objectstorage-bucket"
description: |-
  Provides the Bucket resource in Oracle Cloud Infrastructure Object Storage service
---

# oci_objectstorage_bucket
This resource provides the Bucket resource in Oracle Cloud Infrastructure Object Storage service.

Creates a bucket in the given namespace with a bucket name and optional user-defined metadata. Avoid entering
confidential information in bucket names.


## Example Usage

```hcl
resource "oci_objectstorage_bucket" "test_bucket" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${var.bucket_name}"
	namespace = "${var.bucket_namespace}"

	#Optional
	access_type = "${var.bucket_access_type}"
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	kms_key_id = "${oci_kms_key.test_key.id}"
	metadata = "${var.bucket_metadata}"
	object_events_enabled = "${var.bucket_object_events_enabled}"
	storage_tier = "${var.bucket_storage_tier}"
}
```

## Argument Reference

The following arguments are supported:

* `access_type` - (Optional) (Updatable) The type of public access enabled on this bucket. A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the bucket, public access is allowed for the `GetObject` and `HeadObject` operations. 
* `compartment_id` - (Required) (Updatable) The ID of the compartment in which to create the bucket.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `kms_key_id` - (Optional) (Updatable) The OCID of a master encryption key used to call the Key Management service to generate a data  encryption key or to encrypt or decrypt a data encryption key. 
* `metadata` - (Optional) (Updatable) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
* `name` - (Required) The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, and dashes. Bucket names must be unique within an Object Storage namespace. Avoid entering confidential information. example: Example: my-new-bucket1 
* `namespace` - (Required) The Object Storage namespace used for the request.
* `object_events_enabled` - (Optional) (Updatable) Whether or not events are emitted for object state changes in this bucket. By default, `objectEventsEnabled` is set to `false`. Set `objectEventsEnabled` to `true` to emit events for object state changes. For more information about events, see [Overview of Events](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm). 
* `storage_tier` - (Optional) The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_type` - The type of public access enabled on this bucket. A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the bucket, public access is allowed for the `GetObject` and `HeadObject` operations. 
* `approximate_count` - The approximate number of objects in the bucket. Count statistics are reported periodically. You will see a lag between what is displayed and the actual object count. 
* `approximate_size` - The approximate total size in bytes of all objects in the bucket. Size statistics are reported periodically. You will see a lag between what is displayed and the actual size of the bucket. 
* `compartment_id` - The compartment ID in which the bucket is authorized.
* `created_by` - The OCID of the user who created the bucket.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `etag` - The entity tag (ETag) for the bucket.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `kms_key_id` - The OCID of a master encryption key used to call the Key Management service to generate a data encryption key  or to encrypt or decrypt a data encryption key. 
* `metadata` - Arbitrary string keys and values for user-defined metadata.
* `name` - The name of the bucket. Avoid entering confidential information. Example: my-new-bucket1 
* `namespace` - The Object Storage namespace in which the bucket lives.
* `object_events_enabled` - Whether or not events are emitted for object state changes in this bucket. By default, `objectEventsEnabled` is set to `false`. Set `objectEventsEnabled` to `true` to emit events for object state changes. For more information about events, see [Overview of Events](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm). 
* `object_lifecycle_policy_etag` - The entity tag (ETag) for the live object lifecycle policy on the bucket.
* `storage_tier` - The storage tier type assigned to the bucket. A bucket is set to 'Standard' tier by default, which means objects uploaded or copied to the bucket will be in the standard storage tier. When the 'Archive' tier type is set explicitly for a bucket, objects uploaded or copied to the bucket will be stored in archive storage. The 'storageTier' property is immutable after bucket is created. 
* `time_created` - The date and time the bucket was created, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616), section 14.29.

## Import

Buckets can be imported using the `id`, e.g.

```
$ terraform import oci_objectstorage_bucket.test_bucket "n/{namespaceName}/b/{bucketName}" 
```


---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_bucket"
sidebar_current: "docs-oci-datasource-object_storage-bucket"
description: |-
  Provides details about a specific Bucket in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_bucket
This data source provides details about a specific Bucket resource in Oracle Cloud Infrastructure Object Storage service.

Gets the current representation of the given bucket in the given Object Storage namespace.


## Example Usage

```hcl
data "oci_objectstorage_bucket" "test_bucket" {
	#Required
	name = "${var.bucket_name}"
	namespace = "${var.bucket_namespace}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `namespace` - (Required) The Object Storage namespace used for the request.


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
* `kms_key_id` - The OCID of a KMS key id used to call KMS to generate data key or decrypt the encrypted data key. 
* `metadata` - Arbitrary string keys and values for user-defined metadata.
* `name` - The name of the bucket. Avoid entering confidential information. Example: my-new-bucket1 
* `namespace` - The Object Storage namespace in which the bucket lives.
* `object_lifecycle_policy_etag` - The entity tag (ETag) for the live object lifecycle policy on the bucket.
* `storage_tier` - The storage tier type assigned to the bucket. A bucket is set to 'Standard' tier by default, which means objects uploaded or copied to the bucket will be in the standard storage tier. When the 'Archive' tier type is set explicitly for a bucket, objects uploaded or copied to the bucket will be stored in archive storage. The 'storageTier' property is immutable after bucket is created. 
* `time_created` - The date and time the bucket was created, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616), section 14.29.


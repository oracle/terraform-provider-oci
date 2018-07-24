---
layout: "oci"
page_title: "OCI: oci_objectstorage_bucket"
sidebar_current: "docs-oci-resource-object_storage-bucket"
description: |-
Creates and manages an OCI Bucket
---

# oci_objectstorage_bucket
The `oci_objectstorage_bucket` resource creates and manages an OCI Bucket

Creates a bucket in the given namespace with a bucket name and optional user-defined metadata.


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
	metadata = "${var.bucket_metadata}"
	storage_tier = "${var.bucket_storage_tier}"
}
```

## Argument Reference

The following arguments are supported:

* `access_type` - (Optional) (Updatable) The type of public access enabled on this bucket. A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the bucket, public access is allowed for the `GetObject` and `HeadObject` operations. 
* `compartment_id` - (Required) (Updatable) The ID of the compartment in which to create the bucket.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `metadata` - (Optional) (Updatable) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
* `name` - (Required) The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, and dashes. Bucket names must be unique within the namespace. Avoid entering confidential information. example: Example: my-new-bucket1 
* `storage_tier` - (Optional) The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_type` - The type of public access enabled on this bucket. A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the bucket, public access is allowed for the `GetObject` and `HeadObject` operations. 
* `compartment_id` - The compartment ID in which the bucket is authorized.
* `created_by` - The OCID of the user who created the bucket.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `etag` - The entity tag for the bucket.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `metadata` - Arbitrary string keys and values for user-defined metadata.
* `name` - The name of the bucket. Avoid entering confidential information. Example: my-new-bucket1 
* `namespace` - The namespace in which the bucket lives.
* `storage_tier` - The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the archive storage tier. The 'storageTier' property is immutable after bucket is created. 
* `time_created` - The date and time the bucket was created, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616), section 14.29.

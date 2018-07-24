---
layout: "oci"
page_title: "OCI: oci_object_storage_buckets"
sidebar_current: "docs-oci-datasource-object_storage-buckets"
description: |-
Provides a list of Buckets
---
# Data Source: oci_objectstorage_bucketsummaries
The BucketSummaries data source allows access to the list of OCI buckets

Gets a list of all `BucketSummary`s in a compartment. A `BucketSummary` contains only summary fields for the bucket
and does not contain fields like the user-defined metadata.

To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
talk to an administrator. If you're an administrator who needs to write policies to give users access, see
[Getting Started with Policies](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).


## Example Usage

```hcl
data "oci_objectstorage_bucketsummaries" "test_buckets" {
	#Required
	compartment_id = "${var.compartment_id}"
	namespace = "${var.bucket_namespace}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list buckets.
* `namespace` - (Required) The top-level namespace used for the request.


## Attributes Reference

The following attributes are exported:

* `bucket_summaries` - The list of buckets.

### Bucket Reference

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


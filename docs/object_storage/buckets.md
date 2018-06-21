# oci_objectstorage_bucket

## Bucket Resource

### Bucket Reference

The following attributes are exported:

* `access_type` - The type of public access available on this bucket. Allows authenticated caller to access the bucket or contents of this bucket. By default a bucket is set to NoPublicAccess. It is treated as NoPublicAccess when this value is not specified. When the type is NoPublicAccess the bucket does not allow any public access. When the type is ObjectRead the bucket allows public access to the GetObject, HeadObject, ListObjects. 
* `compartment_id` - The compartment ID in which the bucket is authorized.
* `created_by` - The OCID of the user who created the bucket.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `etag` - The entity tag for the bucket.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `metadata` - Arbitrary string keys and values for user-defined metadata.
* `name` - The name of the bucket.
* `namespace` - The namespace in which the bucket lives.
* `storage_tier` - The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the archive storage tier. The 'storageTier' property is immutable after bucket is created. 
* `time_created` - The date and time at which the bucket was created.



### Create Operation
Creates a bucket in the given namespace with a bucket name and optional user-defined metadata.

To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized, 
talk to an administrator. If you're an administrator who needs to write policies to give users access, see 
[Getting Started with Policies](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).


The following arguments are supported:

* `access_type` - (Optional) The type of public access available on this bucket. Allows authenticated caller to access the bucket or contents of this bucket. By default a bucket is set to NoPublicAccess. It is treated as NoPublicAccess when this value is not specified. When the type is NoPublicAccess the bucket does not allow any public access. When the type is ObjectRead the bucket allows public access to the GetObject, HeadObject, ListObjects. 
* `compartment_id` - (Required) The ID of the compartment in which to create the bucket.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `metadata` - (Optional) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
* `name` - (Required) The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, and dashes. Bucket names must be unique within the namespace. 
* `namespace` - (Required) The top-level namespace used for the request.
* `storage_tier` - (Optional) The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created. 


### Update Operation

The following arguments support updates:
* `access_type` - The type of public access available on this bucket. Allows authenticated caller to access the bucket or contents of this bucket. By default a bucket is set to NoPublicAccess. It is treated as NoPublicAccess when this value is not specified. When the type is NoPublicAccess the bucket does not allow any public access. When the type is ObjectRead the bucket allows public access to the GetObject, HeadObject, ListObjects. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `metadata` - Arbitrary string, up to 4KB, of keys and values for user-defined metadata.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

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

# oci_objectstorage_bucketsummaries

## BucketSummaries DataSource

Gets a list of buckets.

### List Operation
Gets a list of all `BucketSummary`s in a compartment. A `BucketSummary` contains only summary fields for the bucket
and does not contain fields like the user-defined metadata.

To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized, 
talk to an administrator. If you're an administrator who needs to write policies to give users access, see 
[Getting Started with Policies](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list buckets.
* `namespace` - (Required) The top-level namespace used for the request.


The following attributes are exported:

* `bucket_summaries` - The list of buckets.

### Example Usage

```hcl
data "oci_objectstorage_bucketsummaries" "test_buckets" {
	#Required
	compartment_id = "${var.compartment_id}"
	namespace = "${var.bucket_namespace}"
}
```
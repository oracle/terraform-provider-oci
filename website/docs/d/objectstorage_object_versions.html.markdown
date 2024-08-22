---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_object_versions"
sidebar_current: "docs-oci-datasource-objectstorage-object_versions"
description: |-
  Provides the list of Object Versions in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_object_versions
This data source provides the list of Object Versions in Oracle Cloud Infrastructure Object Storage service.

Lists the object versions in a bucket.

ListObjectVersions returns an ObjectVersionCollection containing at most 1000 object versions. To paginate through
more object versions, use the returned `opc-next-page` value with the `page` request parameter.

To use this and other API operations, you must be authorized in an IAM policy. If you are not authorized,
talk to an administrator. If you are an administrator who needs to write policies to give users access, see
[Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).


## Example Usage

```hcl
data "oci_objectstorage_object_versions" "test_object_versions" {
	#Required
	bucket = var.object_version_bucket
	namespace = var.object_version_namespace

	#Optional
	delimiter = var.object_version_delimiter
	end = var.object_version_end
	fields = var.object_version_fields
	prefix = var.object_version_prefix
	start = var.object_version_start
	start_after = var.object_version_start_after
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `delimiter` - (Optional) When this parameter is set, only objects whose names do not contain the delimiter character (after an optionally specified prefix) are returned in the objects key of the response body. Scanned objects whose names contain the delimiter have the part of their name up to the first occurrence of the delimiter (including the optional prefix) returned as a set of prefixes. Note that only '/' is a supported delimiter character at this time. 
* `end` - (Optional) Returns object names which are lexicographically strictly less than this parameter.
* `fields` - (Optional) Object summary by default includes only the 'name' field. Use this parameter to also include 'size' (object size in bytes), 'etag', 'md5', 'timeCreated' (object creation date and time), 'timeModified' (object modification date and time), 'storageTier' and 'archivalState' fields. Specify the value of this parameter as a comma-separated, case-insensitive list of those field names.  For example 'name,etag,timeCreated,md5,timeModified,storageTier,archivalState'. 
* `namespace` - (Required) The Object Storage namespace used for the request.
* `prefix` - (Optional) The string to use for matching against the start of object names in a list query.
* `start` - (Optional) Returns object names which are lexicographically greater than or equal to this parameter.
* `start_after` - (Optional) Returns object names which are lexicographically strictly greater than this parameter.


## Attributes Reference

The following attributes are exported:

* `object_version_collection` - The list of object_version_collection.

### ObjectVersion Reference

The following attributes are exported:

* `items` - An array of object version summaries. 
	* `archival_state` - Archival state of an object. This field is set only for objects in Archive tier.
	* `etag` - The current entity tag (ETag) for the object.
	* `is_delete_marker` - This flag will indicate if the version is deleted or not.
	* `md5` - Base64-encoded MD5 hash of the object data.
	* `name` - The name of the object. Avoid entering confidential information. Example: test/object1.log 
	* `size` - Size of the object in bytes.
	* `storage_tier` - The storage tier that the object is stored in.
	* `time_created` - The date and time the object was created, as described in [RFC 2616](https://tools.ietf.org/html/rfc2616#section-14.29).
	* `time_modified` - The date and time the object was modified, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616#section-14.29).
	* `version_id` - VersionId of the object.
* `prefixes` - Prefixes that are common to the results returned by the request if the request specified a delimiter. 


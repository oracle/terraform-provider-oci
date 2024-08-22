---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_objects"
sidebar_current: "docs-oci-datasource-objectstorage-objects"
description: |-
  Provides the list of Objects in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_objects
This data source provides the list of Objects in Oracle Cloud Infrastructure Object Storage service.

Lists the objects in a bucket. By default, ListObjects returns object names only. See the `fields`
parameter for other fields that you can optionally include in ListObjects response.

ListObjects returns at most 1000 objects. To paginate through more objects, use the returned 'nextStartWith'
value with the 'start' parameter. To filter which objects ListObjects returns, use the 'start' and 'end'
parameters.

To use this and other API operations, you must be authorized in an IAM policy. If you are not authorized,
talk to an administrator. If you are an administrator who needs to write policies to give users access, see
[Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).


## Example Usage

```hcl
data "oci_objectstorage_objects" "test_objects" {
	#Required
	bucket = var.object_bucket
	namespace = var.object_namespace

	#Optional
	delimiter = var.object_delimiter
	end = var.object_end
	fields = var.object_fields
	prefix = var.object_prefix
	start = var.object_start
	start_after = var.object_start_after
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

* `objects` - The list of list_objects.

### Object Reference

The following attributes are exported:

* `name` - The name of the object. 
* `size` - Size of the object in bytes.
* `md5` - Base64-encoded MD5 hash of the object data.
* `time_created` - The date and time the object was created, as described in [RFC 2616](https://tools.ietf.org/html/rfc2616#section-14.29).
* `time_modified` - The date and time the object was modified, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616#section-14.29).
* `etag` - The current entity tag (ETag) for the object.
* `storage_tier` - The storage tier that the object is stored in.
* `archival-state` - Archival state of an object. This field is set only for objects in Archive tier.

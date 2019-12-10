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

Lists the objects in a bucket.

To use this and other API operations, you must be authorized in an IAM policy. If you are not authorized,
talk to an administrator. If you are an administrator who needs to write policies to give users access, see
[Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).


## Example Usage

```hcl
data "oci_objectstorage_objects" "test_objects" {
	#Required
	bucket = "${var.object_bucket}"
	namespace = "${var.object_namespace}"

	#Optional
	delimiter = "${var.object_delimiter}"
	end = "${var.object_end}"
	prefix = "${var.object_prefix}"
	start = "${var.object_start}"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `delimiter` - (Optional) When this parameter is set, only objects whose names do not contain the delimiter character (after an optionally specified prefix) are returned in the objects key of the response body. Scanned objects whose names contain the delimiter have the part of their name up to the first occurrence of the delimiter (including the optional prefix) returned as a set of prefixes. Note that only '/' is a supported delimiter character at this time. 
* `end` - (Optional) Object names returned by a list query must be strictly less than this parameter.
* `namespace` - (Required) The Object Storage namespace used for the request.
* `prefix` - (Optional) The string to use for matching against the start of object names in a list query.
* `start` - (Optional) Object names returned by a list query must be greater or equal to this parameter.


## Attributes Reference

The following attributes are exported:

* `list_objects` - The list of list_objects.

### Object Reference

The following attributes are exported:

* `bucket` - The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `content` - The object to upload to the object store.
* `content_encoding` - The content encoding of the object.
* `content_language` - The content language of the object.
* `content_length` - The content length of the body.
* `content_md5` - The base-64 encoded MD5 hash of the body.
* `content_type` - The content type of the object.  Defaults to 'application/octet-stream' if not overridden during the PutObject call.
* `etag` - The current entity tag (ETag) for the object.
* `metadata` - Optional user-defined metadata key and value.
Note: Metadata keys are case-insensitive and all returned keys will be lower case.
* `namespace` - The top-level namespace used for the request.
* `object` - The name of the object. Avoid entering confidential information. Example: `test/object1.log` 


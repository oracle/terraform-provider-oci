---
subcategory: "Object Storage"
layout: "oci"
page_title: "OCI: oci_objectstorage_object_head"
sidebar_current: "docs-oci-datasource-objectstorage-object_head"
description: |-
  Provides a list of Object metadata
---

# Data Source: oci_objectstorage_object_head
This data source provides details about metadata of a specific Object resource in Oracle Cloud Infrastructure Object Storage service.

Gets the metadata of an object.

## Example Usage

```hcl
data "oci_objectstorage_object_head" "test_object_head" {
	#Required
	bucket = "${var.object_bucket}"
	namespace = "${var.object_namespace}"
	object = "${var.object_object}"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `namespace` - (Required) The top-level namespace used for the request.
* `object` - (Required) The name of the object. Avoid entering confidential information. Example: `test/object1.log` 


## Attributes Reference

The following attributes are exported:
 
* `metadata` - The metadata of the object
* `content_type` - The content-type of the object
* `content_length` - The content-length of the object
* `etag` - The etag of the object


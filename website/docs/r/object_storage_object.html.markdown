---
layout: "oci"
page_title: "OCI: oci_objectstorage_object"
sidebar_current: "docs-oci-resource-object_storage-object"
description: |-
Creates and manages an OCI Object
---

# oci_objectstorage_object
The `oci_objectstorage_object` resource creates and manages an OCI Object

Creates a new object or overwrites an existing one.


## Example Usage

```hcl
resource "oci_objectstorage_object" "test_object" {
	#Required
	bucket = "${var.object_bucket}"
	content = "${var.object_content}"
	namespace = "${var.object_namespace}"
	object = "${var.object_object}"

	#Optional
	content_encoding = "${var.object_content_encoding}"
	content_language = "${var.object_content_language}"
	content_type = "${var.object_content_type}"
	metadata = "${var.object_metadata}"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `content_encoding` - (Optional) The content encoding of the object.
* `content_language` - (Optional) The content language of the object.
* `content_type` - (Optional) The content type of the object.  Defaults to 'application/octet-stream' if not overridden during the PutObject call.
* `content` - (Required) The object to upload to the object store.
* `metadata` - (Optional) Optional user-defined metadata key and value.
Note: All specified keys must be in lower case.
* `namespace` - (Required) The top-level namespace used for the request.
* `object` - (Required) The name of the object. Avoid entering confidential information. Example: `test/object1.log` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bucket` - The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `content` - The object to upload to the object store.
* `content_encoding` - The content encoding of the object.
* `content_language` - The content language of the object.
* `content_length` - The content length of the body.
* `content_md5` - The base-64 encoded MD5 hash of the body.
* `content_type` - The content type of the object.  Defaults to 'application/octet-stream' if not overridden during the PutObject call.
* `metadata` - Optional user-defined metadata key and value.
Note: Metadata keys are case-insensitive and all returned keys will be lower case.
* `namespace` - The top-level namespace used for the request.
* `object` - The name of the object. Avoid entering confidential information. Example: `test/object1.log` 

## Import

Objects can be imported using the `id`, e.g.

```
$ terraform import oci_object_storage_object.test_object "id"
```

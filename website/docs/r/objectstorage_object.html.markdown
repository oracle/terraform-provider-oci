---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_object"
sidebar_current: "docs-oci-resource-objectstorage-object"
description: |-
  Provides the Object resource in Oracle Cloud Infrastructure Object Storage service
---

# oci_objectstorage_object
This resource provides the Object resource in Oracle Cloud Infrastructure Object Storage service.

Creates a new object or overwrites an existing object with the same name. The maximum object size allowed by
PutObject is 50 GiB.

See [Special Instructions for Object Storage PUT](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/signingrequests.htm#ObjectStoragePut)
for request signature requirements.


## Example Usage

```hcl
resource "oci_objectstorage_object" "test_object" {
	#Required
	bucket = "${var.object_bucket}"
	content = "${var.object_content}"
	namespace = "${var.object_namespace}"
	object = "${var.object_object}"

	#Optional
	cache_control = "${var.object_cache_control}"
	content_disposition = "${var.object_content_disposition}"
	content_encoding = "${var.object_content_encoding}"
	content_language = "${var.object_content_language}"
	content_type = "${var.object_content_type}"
	delete_all_object_versions = "${var.object_delete_all_object_versions}"
	metadata = "${var.object_metadata}"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `cache_control` - (Optional) The cache-control header value to be returned in GetObjectResponse.
* `content` - (Required) The object to upload to the object store. Cannot be defined if `source` or `source_uri_details` is defined.
* `content_disposition` - (Optional) The Content-Disposition header value to be returned in GetObjectResponse.
* `content_encoding` - (Optional) The content encoding of the object.
* `content_language` - (Optional) The content language of the object.
* `content_type` - (Optional) The content type of the object.  Defaults to 'application/octet-stream' if not overridden during the PutObject call.
* `delete_all_object_versions` - (Optional) (Updatable) A boolean to delete all object versions for an object in a bucket that has or ever had versioning enabled.
* `metadata` - (Optional) Optional user-defined metadata key and value.
Note: All specified keys must be in lower case.
* `namespace` - (Required) The Object Storage namespace used for the request.
* `object` - (Required) (Updatable) The name of the object. Avoid entering confidential information. Example: `test/object1.log` 
* `source` - (Optional) An absolute path to a file on the local system. Cannot be defined if `content` or `source_uri_details` is defined.
* `source_uri_details` - (Optional) Details of the source URI of the object in the cloud. Cannot be defined if `content` or `source` is defined. 
Note: To enable object copy, you must authorize the service to manage objects on your behalf.
    * `region` - (Required) The region of the source object.
    * `namespace` - (Required) The top-level namespace of the source object.
    * `bucket` - (Required) The name of the bucket for the source object.
    * `object` - (Required) The name of the source object.
    * `source_object_if_match_etag` - (Optional) The entity tag to match the source object.
    * `destination_object_if_match_etag` - (Optional) The entity tag to match the target object.
    * `destination_object_if_none_match_etag` - (Optional) The entity tag to not match the target object.
    * `source_version_id` - (Optional) The version id of the object to be restored.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bucket` - The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1`
* `cache_control` - The cache-control header value to be returned in GetObjectResponse. 
* `content` - The object to upload to the object store.
* `content_disposition` - The Content-Disposition header value to be returned in GetObjectResponse.
* `content_encoding` - The content encoding of the object.
* `content_language` - The content language of the object.
* `content_length` - The content length of the body.
* `content_md5` - The base-64 encoded MD5 hash of the body.
* `content_type` - The content type of the object.  Defaults to 'application/octet-stream' if not overridden during the PutObject call.
* `metadata` - Optional user-defined metadata key and value.
Note: Metadata keys are case-insensitive and all returned keys will be lower case.
* `namespace` - The top-level namespace used for the request.
* `object` - The name of the object. Avoid entering confidential information. Example: `test/object1.log` 
* `source` - An absolute path to a file on the local system to upload to the object store.
* `source_uri_details` - Details of the source URI of the object in the cloud. 
    * `region` - The region of the source object.
    * `namespace` - The top-level namespace of the source object.
    * `bucket` - The name of the bucket for the source object.
    * `object` - The name of the source object.

## Import

Objects can be imported using the `id`, e.g.

```
$ terraform import oci_objectstorage_object.test_object "n/{namespaceName}/b/{bucketName}/o/{objectName}" 
```


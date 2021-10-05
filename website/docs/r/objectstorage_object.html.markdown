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

See [Object Names](https://docs.cloud.oracle.com/iaas/Content/Object/Tasks/managingobjects.htm#namerequirements)
for object naming requirements. 

See [Special Instructions for Object Storage PUT](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/signingrequests.htm#ObjectStoragePut)
for request signature requirements.


## Example Usage

```hcl
resource "oci_objectstorage_object" "test_object" {
	#Required
	bucket = var.object_bucket
	content = var.object_content
	namespace = var.object_namespace
	object = var.object_object

	#Optional
	cache_control = var.object_cache_control
	content_disposition = var.object_content_disposition
	content_encoding = var.object_content_encoding
	content_language = var.object_content_language
	content_type = var.object_content_type
	delete_all_object_versions = var.object_delete_all_object_versions
	metadata = var.object_metadata
	storage_tier = var.object_storage_tier
    opc_sse_kms_key_id = var.object_opc_sse_kms_key_id
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `cache_control` - (Optional) The optional Cache-Control header that defines the caching behavior value to be returned in GetObject and HeadObject responses. Specifying values for this header has no effect on Object Storage behavior. Programs that read the object determine what to do based on the value provided. For example, you could use this header to identify objects that require caching restrictions. 
* `content` - (Optional) The object to upload to the object store. Cannot be defined if `source` or `source_uri_details` is defined.
* `content_disposition` - (Optional) The optional Content-Disposition header that defines presentational information for the object to be returned in GetObject and HeadObject responses. Specifying values for this header has no effect on Object Storage behavior. Programs that read the object determine what to do based on the value provided. For example, you could use this header to let users download objects with custom filenames in a browser. 
* `content_encoding` - (Optional) The optional Content-Encoding header that defines the content encodings that were applied to the object to upload. Specifying values for this header has no effect on Object Storage behavior. Programs that read the object determine what to do based on the value provided. For example, you could use this header to determine what decoding mechanisms need to be applied to obtain the media-type specified by the Content-Type header of the object. 
* `content_language` - (Optional) The optional Content-Language header that defines the content language of the object to upload. Specifying values for this header has no effect on Object Storage behavior. Programs that read the object determine what to do based on the value provided. For example, you could use this header to identify and differentiate objects based on a particular language. 
* `content_length` - (Required) (Updatable) The content length of the body.
* `content_md5` - (Optional) (Updatable) The optional base-64 header that defines the encoded MD5 hash of the body. If the optional Content-MD5 header is present, Object Storage performs an integrity check on the body of the HTTP request by computing the MD5 hash for the body and comparing it to the MD5 hash supplied in the header. If the two hashes do not match, the object is rejected and an HTTP-400 Unmatched Content MD5 error is returned with the message:

	"The computed MD5 of the request body (ACTUAL_MD5) does not match the Content-MD5 header (HEADER_MD5)" 
* `content_type` - (Optional) The optional Content-Type header that defines the standard MIME type format of the object. Content type defaults to 'application/octet-stream' if not specified in the PutObject call. Specifying values for this header has no effect on Object Storage behavior. Programs that read the object determine what to do based on the value provided. For example, you could use this header to identify and perform special operations on text only objects. 
* `delete_all_object_versions` - (Optional) (Updatable) A boolean to delete all object versions for an object in a bucket that has or ever had versioning enabled.
* `metadata` - (Optional) Optional user-defined metadata key and value.
Note: All specified keys must be in lower case.
* `namespace` - (Required) The Object Storage namespace used for the request.
* `object` - (Required) (Updatable) The name of the object. Avoid entering confidential information. Example: `test/object1.log` 
* `opc_sse_kms_key_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
* `storage_tier` - (Optional) (Updatable) The storage tier that the object should be stored in. If not specified, the object will be stored in the same storage tier as the bucket. 
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Object
	* `update` - (Defaults to 20 minutes), when updating the Object
	* `delete` - (Defaults to 20 minutes), when destroying the Object


## Import

Objects can be imported using the `id`, e.g.

```
$ terraform import oci_objectstorage_object.test_object "n/{namespaceName}/b/{bucketName}/o/{objectName}" 
```


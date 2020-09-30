---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_object"
sidebar_current: "docs-oci-datasource-objectstorage-object"
description: |-
  Provides details about a specific Object in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_object
This data source provides details about a specific Object resource in Oracle Cloud Infrastructure Object Storage service.

Gets the metadata and body of an object.


## Example Usage

```hcl
data "oci_objectstorage_object" "test_object" {
	#Required
	bucket = var.object_bucket
	namespace = var.object_namespace
	object = var.object_object

	#Optional
	http_response_cache_control = var.object_http_response_cache_control
	http_response_content_disposition = var.object_http_response_content_disposition
	http_response_content_encoding = var.object_http_response_content_encoding
	http_response_content_language = var.object_http_response_content_language
	http_response_content_type = var.object_http_response_content_type
	http_response_expires = var.object_http_response_expires
	version_id = oci_objectstorage_version.test_version.id
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `http_response_cache_control` - (Optional) This value will be used in Cache-Control header of the response.
* `http_response_content_disposition` - (Optional) This value will be used in Content-Disposition header of the response.
* `http_response_content_encoding` - (Optional) This value will be used in Content-Encoding header of the response
* `http_response_content_language` - (Optional) This value will be used in Content-Language header of the response.
* `http_response_content_type` - (Optional) This value will be used in Content-Type header of the response.
* `http_response_expires` - (Optional) This value will be used in Expires header of the response
* `namespace` - (Required) The Object Storage namespace used for the request.
* `object` - (Required) The name of the object. Avoid entering confidential information. Example: `test/object1.log` 
* `content_length_limit` - (Optional) The limit of the content length of the object body to download from the object store. The default is 1Mb.
* `base64_encode_content` - (Optional) Encodes the downloaded content in base64. It is recommended to set this to `true` for binary content to avoid corrupting the zip file in Terraform state. The default value is `false` to preserve backwards compatibility with Terraform v0.11 configurations.
If passing the base64 encoded content to a `local_file` resource, please use the `content_base64` attribute of the `local_file` resource.
* `version_id` - (Optional) VersionId used to identify a particular version of the object


## Attributes Reference

The following attributes are exported:

* `bucket` - The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `content` - The object to upload to the object store.
* `content_encoding` - The content encoding of the object.
* `content_language` - The content language of the object.
* `content_length` - The content length of the body.
* `content_md5` - The base-64 encoded MD5 hash of the body.
* `content_type` - The content type of the object.  Defaults to 'application/octet-stream' if not overridden during the PutObject call.
* `metadata` - Optional user-defined metadata key and value. Note: Metadata keys are case-insensitive and all returned keys will be lower case.
* `namespace` - The top-level namespace used for the request.
* `object` - The name of the object. Avoid entering confidential information. Example: `test/object1.log` 



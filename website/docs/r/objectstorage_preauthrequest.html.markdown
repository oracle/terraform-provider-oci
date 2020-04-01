---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_preauthrequest"
sidebar_current: "docs-oci-resource-objectstorage-preauthrequest"
description: |-
  Provides the Preauthenticated Request resource in Oracle Cloud Infrastructure Object Storage service
---

# oci_objectstorage_preauthrequest
This resource provides the Preauthenticated Request resource in Oracle Cloud Infrastructure Object Storage service.

Creates a pre-authenticated request specific to the bucket.


## Example Usage

```hcl
resource "oci_objectstorage_preauthrequest" "test_preauthenticated_request" {
	#Required
	access_type = "${var.preauthenticated_request_access_type}"
	bucket = "${var.preauthenticated_request_bucket}"
	name = "${var.preauthenticated_request_name}"
	namespace = "${var.preauthenticated_request_namespace}"
	time_expires = "${var.preauthenticated_request_time_expires}"

	#Optional
	object = "${var.preauthenticated_request_object}"
}
```

## Argument Reference

The following arguments are supported:

* `access_type` - (Required) The operation that can be performed on this resource. Allowed Values: `ObjectRead`, `ObjectWrite`, `ObjectReadWrite`, or `AnyObjectWrite`
* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `name` - (Required) A user-specified name for the pre-authenticated request. Names can be helpful in managing pre-authenticated requests.
* `namespace` - (Required) The Object Storage namespace used for the request.
* `object` - (Optional) The name of the object that is being granted access to by the pre-authenticated request. Avoid entering confidential information. The object name can be null and if so, the pre-authenticated request grants access to the entire bucket. 
* `time_expires` - (Required) The expiration date for the pre-authenticated request as per [RFC 3339](https://tools.ietf.org/html/rfc3339). After this date the pre-authenticated request will no longer be valid. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_type` - The operation that can be performed on this resource.
* `access_uri` - The URI to embed in the URL `https://objectstorage.${var.region}.oraclecloud.com{var.access_uri}` when using the pre-authenticated request.
* `bucket` - The name of the bucket.  Example: `my-new-bucket1` 
* `id` - The unique identifier to use when directly addressing the pre-authenticated request.
* `name` - The user-provided name of the pre-authenticated request.
* `object` - The name of the object that is being granted access to by the pre-authenticated request. Avoid entering confidential information. The object name can be null and if so, the pre-authenticated request grants access to the entire bucket. Example: test/object1.log 
* `time_created` - The date when the pre-authenticated request was created as per specification [RFC 3339](https://tools.ietf.org/html/rfc3339). 
* `time_expires` - The expiration date for the pre-authenticated request as per [RFC 3339](https://tools.ietf.org/html/rfc3339). After this date the pre-authenticated request will no longer be valid. 
* `namespace` - The top-level namespace used for the request.
* `par_id` - The unique identifier for the pre-authenticated request. This can be used to manage operations against the pre-authenticated request, such as GET or DELETE.

## Import

PreauthenticatedRequests can be imported using the `id`, e.g.

```
$ terraform import oci_objectstorage_preauthrequest.test_preauthenticated_request "n/{namespaceName}/b/{bucketName}/p/{parId}" 
```


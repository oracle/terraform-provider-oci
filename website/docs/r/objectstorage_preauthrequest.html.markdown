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
	access_type = var.preauthenticated_request_access_type
	bucket = var.preauthenticated_request_bucket
	name = var.preauthenticated_request_name
	namespace = var.preauthenticated_request_namespace
	time_expires = var.preauthenticated_request_time_expires

	#Optional
	bucket_listing_action = var.preauthenticated_request_bucket_listing_action
	object = var.preauthenticated_request_object
}
```

## Argument Reference

The following arguments are supported:

* `access_type` - (Required) The operation that can be performed on this resource. Allowed Values: `ObjectRead`, `ObjectWrite`, `ObjectReadWrite`, `AnyObjectReadWrite` or `AnyObjectRead`
* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `bucket_listing_action` - (Optional) Specifies whether a list operation is allowed on a PAR with accessType "AnyObjectRead" or "AnyObjectReadWrite". Deny: Prevents the user from performing a list operation. ListObjects: Authorizes the user to perform a list operation. 
* `name` - (Required) A user-specified name for the pre-authenticated request. Names can be helpful in managing pre-authenticated requests. Avoid entering confidential information. 
* `namespace` - (Required) The Object Storage namespace used for the request.
* `object` - Deprecated. Instead use `object_name`. Requests that include both `object` and `object_name` will be rejected. (Optional) The name of the object that is being granted access to by the pre-authenticated request. Avoid entering confidential information. The object name can be null and if so, the pre-authenticated request grants access to the entire bucket if the access type allows that. The object name can be a prefix as well, in that case pre-authenticated request grants access to all the objects within the bucket starting with that prefix provided that we have the correct access type.
* `object_name` - (Optional) The name of the object that is being granted access to by the pre-authenticated request. Avoid entering confidential information. The object name can be null and if so, the pre-authenticated request grants access to the entire bucket if the access type allows that. The object name can be a prefix as well, in that case pre-authenticated request grants access to all the objects within the bucket starting with that prefix provided that we have the correct access type.
* `time_expires` - (Required) The expiration date for the pre-authenticated request as per [RFC 3339](https://tools.ietf.org/html/rfc3339). After this date the pre-authenticated request will no longer be valid. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_type` - The operation that can be performed on this resource.
* `bucket_listing_action` - Specifies whether a list operation is allowed on a PAR with accessType "AnyObjectRead" or "AnyObjectReadWrite". Deny: Prevents the user from performing a list operation. ListObjects: Authorizes the user to perform a list operation. 
* `access_uri` - The URI to embed in the URL `https://objectstorage.${var.region}.oraclecloud.com{var.access_uri}` when using the pre-authenticated request.
* `bucket` - The name of the bucket.  Example: `my-new-bucket1` 
* `id` - The unique identifier to use when directly addressing the pre-authenticated request.
* `name` - The user-provided name of the pre-authenticated request.
* `namespace` - The top-level namespace used for the request.
* `object` - The name of the object that is being granted access to by the pre-authenticated request. Avoid entering confidential information. The object name can be null and if so, the pre-authenticated request grants access to the entire bucket. Example: test/object1.log 
* `par_id` - The unique identifier for the pre-authenticated request. This can be used to manage operations against the pre-authenticated request, such as GET or DELETE.
* `time_created` - The date when the pre-authenticated request was created as per specification [RFC 3339](https://tools.ietf.org/html/rfc3339). 
* `time_expires` - The expiration date for the pre-authenticated request as per [RFC 3339](https://tools.ietf.org/html/rfc3339). After this date the pre-authenticated request will no longer be valid. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Preauthenticated Request
	* `update` - (Defaults to 20 minutes), when updating the Preauthenticated Request
	* `delete` - (Defaults to 20 minutes), when destroying the Preauthenticated Request


## Import

PreauthenticatedRequests can be imported using the `id`, e.g.

```
$ terraform import oci_objectstorage_preauthrequest.test_preauthenticated_request "n/{namespaceName}/b/{bucketName}/p/{parId}" 
```


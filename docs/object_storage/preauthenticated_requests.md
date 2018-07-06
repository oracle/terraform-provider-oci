# oci_objectstorage_preauthrequest

## PreauthenticatedRequest Resource

### PreauthenticatedRequest Reference

The following attributes are exported:

* `access_type` - The operation that can be performed on this resource.
* `access_uri` - The URI to embed in the URL when using the pre-authenticated request.
* `bucket` - The name of the bucket.  Example: `my-new-bucket1` 
* `id` - The unique identifier to use when directly addressing the pre-authenticated request.
* `name` - The user-provided name of the pre-authenticated request.
* `namespace` - The top-level namespace used for the request.
* `object` - The name of object that is being granted access to by the pre-authenticated request. This can be null and if it is, the pre-authenticated request grants access to the entire bucket.
* `time_created` - The date when the pre-authenticated request was created as per [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_expires` - The expiration date for the pre-authenticated request as per [RFC 3339](https://tools.ietf.org/rfc/rfc3339). After this date the pre-authenticated request will no longer be valid. 



### Create Operation
Creates a pre-authenticated request specific to the bucket.


The following arguments are supported:

* `access_type` - (Required) The operation that can be performed on this resource.
* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `name` - (Required) A user-specified name for the pre-authenticated request. Helpful for management purposes.
* `namespace` - (Required) The top-level namespace used for the request.
* `object` - (Optional) The name of object that is being granted access to by the pre-authenticated request. This can be null and if it is, the pre-authenticated request grants access to the entire bucket.
* `time_expires` - (Required) The expiration date for the pre-authenticated request as per [RFC 3339](https://tools.ietf.org/rfc/rfc3339). After this date the pre-authenticated request will no longer be valid. 


### Update Operation


The following arguments support updates:
* NO arguments in this resource support updates

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

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


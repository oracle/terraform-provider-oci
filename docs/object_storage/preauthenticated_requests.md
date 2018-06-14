# oci_objectstorage_preauthrequest

## PreauthenticatedRequest Resource

### PreauthenticatedRequest Reference

The following attributes are exported:

* `access_type` - the operation that can be performed on this resource e.g PUT or GET.
* `access_uri` - The URI to embed in the URL when using the pre-authenticated request.
* `bucket` - (Required) The name of the bucket.  Example: `my-new-bucket1` 
* `id` - the unique identifier to use when directly addressing the pre-authenticated request
* `name` - the user supplied name of the pre-authenticated request
* `namespace` - (Required) The top-level namespace used for the request.
* `object` - Name of object that is being granted access to by the pre-authenticated request. This can be null and that would mean that the pre-authenticated request is granting access to the entire bucket
* `time_created` - the date when the pre-authenticated request was created as per spec [RFC 3339](https://tools.ietf.org/rfc/rfc3339) 
* `time_expires` - the expiration date after which the pre authenticated request will no longer be valid as per spec [RFC 3339](https://tools.ietf.org/rfc/rfc3339) 



### Create Operation
Create a pre-authenticated request specific to the bucket


The following arguments are supported:

* `access_type` - (Required) the operation that can be performed on this resource e.g PUT or GET.
* `bucket` - (Required) The name of the bucket.  Example: `my-new-bucket1` 
* `name` - (Required) user specified name for pre-authenticated request. Helpful for management purposes.
* `namespace` - (Required) The top-level namespace used for the request.
* `object` - (Optional) Name of object that is being granted access to by the pre-authenticated request. This can be null and that would mean that the pre-authenticated request is granting access to the entire bucket
* `time_expires` - (Required) The expiration date after which the pre-authenticated request will no longer be valid per spec [RFC 3339](https://tools.ietf.org/rfc/rfc3339) 


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


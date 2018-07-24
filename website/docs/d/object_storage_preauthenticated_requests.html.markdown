---
layout: "oci"
page_title: "OCI: oci_objectstorage_preauthenticated_requests"
sidebar_current: "docs-oci-datasource-object_storage-preauthenticated_requests"
description: |-
Provides a list of PreauthenticatedRequests
---
# Data Source: oci_objectstorage_preauthenticated_requests
The PreauthenticatedRequests data source allows access to the list of OCI preauthenticated_requests

Lists pre-authenticated requests for the bucket.


## Example Usage

```hcl
data "oci_objectstorage_preauthenticated_requests" "test_preauthenticated_requests" {
	#Required
	bucket = "${var.preauthenticated_request_bucket}"
	namespace = "${var.preauthenticated_request_namespace}"

	#Optional
	object_name_prefix = "${var.preauthenticated_request_object_name_prefix}"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `namespace` - (Required) The top-level namespace used for the request.
* `object_name_prefix` - (Optional) User-specified object name prefixes can be used to query and return a list of pre-authenticated requests.


## Attributes Reference

The following attributes are exported:

* `preauthenticated_requests` - The list of preauthenticated_requests.

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


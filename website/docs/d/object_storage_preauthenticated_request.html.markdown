---
layout: "oci"
page_title: "OCI: oci_objectstorage_preauthrequest"
sidebar_current: "docs-oci-datasource-object_storage-preauthenticated_request"
description: |-
  Provides details about a specific PreauthenticatedRequest
---

# Data Source: oci_objectstorage_preauthrequest
The PreauthenticatedRequest data source provides details about a specific PreauthenticatedRequest

Gets the pre-authenticated request for the bucket.

## Example Usage

```hcl
data "oci_objectstorage_preauthrequest" "test_preauthenticated_request" {
	#Required
	bucket = "${var.preauthenticated_request_bucket}"
	namespace = "${var.preauthenticated_request_namespace}"
	par_id = "${var.preauthenticated_request_par_id}"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `namespace` - (Required) The top-level namespace used for the request.
* `par_id` - (Required) The unique identifier for the pre-authenticated request. This can be used to manage operations against the pre-authenticated request, such as GET or DELETE. 


## Attributes Reference

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


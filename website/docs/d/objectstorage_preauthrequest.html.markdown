---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_preauthrequest"
sidebar_current: "docs-oci-datasource-objectstorage-preauthenticated_request"
description: |-
  Provides details about a specific Preauthenticated Request in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_preauthrequest
This data source provides details about a specific Preauthenticated Request resource in Oracle Cloud Infrastructure Object Storage service.

Gets the pre-authenticated request for the bucket.

## Example Usage

```hcl
data "oci_objectstorage_preauthrequest" "test_preauthenticated_request" {
	#Required
	bucket = var.preauthenticated_request_bucket
	namespace = var.preauthenticated_request_namespace
	par_id = oci_objectstorage_preauthrequest.test_par.id
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `namespace` - (Required) The Object Storage namespace used for the request.
* `par_id` - (Required) The unique identifier for the pre-authenticated request. This can be used to manage operations against the pre-authenticated request, such as GET or DELETE. 


## Attributes Reference

The following attributes are exported:

* `access_type` - The operation that can be performed on this resource.
* `bucket_listing_action` - Specifies whether a list operation is allowed on a PAR with accessType "AnyObjectRead" or "AnyObjectReadWrite". Deny: Prevents the user from performing a list operation. ListObjects: Authorizes the user to perform a list operation. 
* `access_uri` - The URI to embed in the URL when using the pre-authenticated request.
* `bucket` - The name of the bucket.  Example: `my-new-bucket1` 
* `id` - The unique identifier to use when directly addressing the pre-authenticated request.
* `name` - The user-provided name of the pre-authenticated request.
* `namespace` - The Object Storage namespace used for the request.
* `object` - Deprecated. Instead use `object_name`.The name of the object that is being granted access to by the pre-authenticated request. Avoid entering confidential information. The object name can be null and if so, the pre-authenticated request grants access to the entire bucket. Example: test/object1.log
* `object_name` - The name of the object that is being granted access to by the pre-authenticated request. Avoid entering confidential information. The object name can be null and if so, the pre-authenticated request grants access to the entire bucket. Example: test/object1.log  
* `time_created` - The date when the pre-authenticated request was created as per specification [RFC 3339](https://tools.ietf.org/html/rfc3339). 
* `time_expires` - The expiration date for the pre-authenticated request as per [RFC 3339](https://tools.ietf.org/html/rfc3339). After this date the pre-authenticated request will no longer be valid. 


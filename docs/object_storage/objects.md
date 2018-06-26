# oci_objectstorage_object

## Object Resource

### Object Reference

The following attributes are exported:

* `bucket` - The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `content` - The object to upload to the object store.
* `content_encoding` - The content encoding of the object.
* `content_language` - The content language of the object.
* `content_length` - The content length of the body.
* `content_md5` - The base-64 encoded MD5 hash of the body.
* `content_type` - The content type of the object.  Defaults to 'application/octet-stream' if not overridden during the PutObject call.
* `metadata` - Optional user-defined metadata key and value.
Note: Metadata keys are case-insensitive and all returned keys will be lower case.
* `namespace` - The top-level namespace used for the request.
* `object` - The name of the object. Avoid entering confidential information. Example: `test/object1.log` 



### Create Operation
Creates a new object or overwrites an existing one.


The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `content` - (Optional) The object to upload to the object store.
* `content_encoding` - (Optional) The content encoding of the object.
* `content_language` - (Optional) The content language of the object.
* `content_type` - (Optional) The content type of the object.  Defaults to 'application/octet-stream' if not overridden during the PutObject call.
* `metadata` - (Optional) Optional user-defined metadata key and value.
Note: All specified keys must be in lower case.
* `namespace` - (Required) The top-level namespace used for the request.
* `object` - (Required) The name of the object. Avoid entering confidential information. Example: `test/object1.log` 


### Update Operation
Creates a new object or overwrites an existing one.


The following arguments support updates:
* `object` - The name of the object. Avoid entering confidential information. Example: `test/object1.log` 

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_objectstorage_object" "test_object" {
	#Required
	bucket = "${var.object_bucket}"
	content = "${var.object_content}"
	namespace = "${var.object_namespace}"
	object = "${var.object_object}"

	#Optional
	content_encoding = "${var.object_content_encoding}"
	content_language = "${var.object_content_language}"
	content_type = "${var.object_content_type}"
	metadata = "${var.object_metadata}"
}
```

# oci_objectstorage_objects

## Objects DataSource

Lists object summaries

### List Operation
Lists the objects in a bucket.

To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized, 
talk to an administrator. If you're an administrator who needs to write policies to give users access, see 
[Getting Started with Policies](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `delimiter` - (Optional) When this parameter is set, only objects whose names do not contain the delimiter character (after an optionally specified prefix) are returned in the objects key of the response body. Scanned objects whose names contain the delimiter have the part of their name up to the first occurrence of the delimiter (including the optional prefix) returned as a set of prefixes. Note that only '/' is a supported delimiter character at this time. 
* `end` - (Optional) Object names returned by a list query must be strictly less than this parameter.
* `namespace` - (Required) The top-level namespace used for the request.
* `prefix` - (Optional) The string to use for matching against the start of object names in a list query.
* `start` - (Optional) Object names returned by a list query must be greater or equal to this parameter.


The following attributes are exported:
 
* `objects` - An array of object summaries. 
	* `md5` - Base64-encoded MD5 hash of the object data.
	* `name` - The name of the object. Avoid entering confidential information. Example: test/object1.log 
	* `size` - Size of the object in bytes.
	* `time_created` - The date and time the object was created, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616), section 14.29.
* `prefixes` - Prefixes that are common to the results returned by the request if the request specified a delimiter. 


### Example Usage

```hcl
data "oci_objectstorage_objects" "test_objects" {
	#Required
	bucket = "${var.object_bucket}"
	namespace = "${var.object_namespace}"

	#Optional
	delimiter = "${var.object_delimiter}"
	end = "${var.object_end}"
	prefix = "${var.object_prefix}"
	start = "${var.object_start}"
}
```

# oci_objectstorage_object_head

## Object_Head DataSource

Provides a datasource for fetching object metadata.

### Get Operation
Gets the user-defined metadata and entity tag for an object.

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `namespace` - (Required) The top-level namespace used for the request.
* `object` - (Required) The name of the object. Avoid entering confidential information. Example: `test/object1.log` 


The following attributes are exported:
 
* `metadata` - The metadata of the object
* `content-type` - The content-type of the object
* `content-length` - The content-length of the object


### Example Usage

```hcl
data "oci_objectstorage_object_head" "test_object_head" {
	#Required
	bucket = "${var.object_bucket}"
	namespace = "${var.object_namespace}"
	object = "${var.object_object}"
}
```
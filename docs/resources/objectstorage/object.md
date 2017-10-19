# oci\_objectstorage\_bucket

Provides an Objectstorage resource for CRUD operations on objects.
`**Dan: not sure which links to use here**`

## Example Usage

### Object

```
resource "oci_objectstorage_object" "t" {
    namespace = "namespaceID"
    bucket = "bucketID"
    object = "objectID"
    content = "the content"
    content_type = "text/plain"
    content_language = "en-US"
    content_encoding = "identity"
    metadata = {
        "version" = "1"
    }
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The namespace of the object store that the object is in.
* `bucket` - (Required) The name of the bucket. Avoid entering confidential information.
* `object` - (Required) The name of the object. Avoid entering confidential information.
* `content` - (Optional) A string that will form the body of the object.
* `metadata` - (Optional) User-defined metadata key value pairs.
* `content_type` - (Optional) The content type of the object. Defaults to 'application/octet-stream' if not overridden during the PutObject call.
* `content_language` - (Optional) The content language of the object.
* `content_encoding` - (Optional) The content encoding of the object.

## Additional Attributes
* `content_length` - The content length of the body.
* `content_md5` - The base-64 encoded MD5 hash of the body.

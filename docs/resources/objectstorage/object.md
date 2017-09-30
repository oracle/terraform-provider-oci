# oci\_objectstorage\_bucket

Provides an Objectstorage resource for CRUD operations on objects.
`**Dan: not sure which links to use here**`

## Example Usage

### Object w/ Metadata

```
resource "oci_objectstorage_object" "t" {
    namespace = "namespaceID"
    bucket = "bucketID"
    object = "objectID"
    content = "bodyContent"
    metadata = {
        "foo" = "bar"
    }
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The namespace of the object store that the object is in.
* `bucket` - (Required) The name of the bucket. Avoid entering confidential information.
* `object` - (Required) The name of the object. Avoid entering confidential information.
* `content` - (Optional) A string that will form the body of the object.

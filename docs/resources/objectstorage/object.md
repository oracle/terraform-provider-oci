# baremetal\_objectstorage\_bucket

Provides an Objectstorage resource for CRUD operations on objects.

## Example Usage

### Object w/ Metadata

```
resource "baremetal_objectstorage_object" "t" {
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

* `namespace` - (Required) The namespace of the object storage that the object is in.
* `bucket` - (Required) The name of the bucket.
* `object` - (Required) The name of the object.
* `content` - (Optional) A string that will form the body of the object.

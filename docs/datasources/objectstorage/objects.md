# baremetal\_objectstorage\_bucket

Provides an Objectstorage datasource for listing objects.

## Example Usage

### Object w/ Metadata

```
data "baremetal_objectstorage_objects" "t" {
    namespace = "namespaceID"
    bucket = "bucketID"
    prefix = "startswith-"
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The namespace of the object storage that the object is in.
* `bucket` - (Required) The name of the bucket.
* `prefix` - (Required) The name of the object.
* `start` - (Optional) The lexigraphically "minimum" string to return.
* `end` - (Optional) The lexigraphically "maximum" string to return.
* `limit` - (Optional) The maximum number of value to return

## Attributes Reference

The following attributes are exported:

* `objects` - The list of objects. They will have the fields: name, md5, size and time_created

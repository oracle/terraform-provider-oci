# oci\_objectstorage\_object_head

Provides a datasource for fetching object metadata.

## Example Usage

### Object Metadata

```
data "oci_objectstorage_object_head" "t" {
     namespace = "namespaceID"
     bucket = "bucketID"
     object = "objectID"
}
```

## Argument Reference

* `namespace` - (Required) The namespace of the object storage that the object is in.
* `bucket` - (Required) The name of the bucket in the namespace that the object is in.
* `object` - (Required) The name of the object in the bucket.

## Attribute Reference

* `metadata` - (Computed) The metadata of the object
* `content-type` - (Computed) The content-type of the object
* `content-length` - (Computed) The content-length of the object

# baremetal\_objectstorage\_bucket

Provides an Objectstorage datasource for fetching an objects metadata

## Example Usage

### Object w/ Metadata

```
data "baremetal_objectstorage_objecthead" "t" {
     namespace = "namespaceID"
     bucket = "bucketID"
     object = "objectID"
}
```

## Argument Reference

* `namespace` - (Required) The namespace of the object storage that the object is in.
* `bucket` - (Required) The name of the bucket in the namespace that the object is in.
* `object` - (Required) The name of the object in the bucket

## Attribute Reference

* `metadata` - (Computed) The metadata of the object
* `content-type` - (Computed) The content-type of the object
* `content-length` - (Computed) The content-length of the object
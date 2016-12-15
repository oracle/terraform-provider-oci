# baremetal\_objectstorage\_bucket

Provides an Objectstorage datasource for fetching the authenticated user's namespace.

## Example Usage

### Object w/ Metadata

```
data "baremetal_objectstorage_namespace" "t" {}
```

## Argument Reference

This datasource takes no arguments

* `namespace` - (Computed) The namespace of the object storage that the object is in.
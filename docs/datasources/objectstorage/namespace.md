# oci\_objectstorage\_namespace


Provides a datasource for fetching the authenticated user's namespace.

## Example Usage

### Object

```
data "oci_objectstorage_namespace" "t" {}
```

## Argument Reference

This datasource takes no arguments.

## Attribute Reference

* `namespace` - (Computed) The namespace of the bucket that the object is in.

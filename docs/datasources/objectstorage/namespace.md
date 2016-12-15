# baremetal\_objectstorage\_namespace

Provides an Objectstorage datasource for fetching the authenticated user's namespace.

## Example Usage

### Object

```
data "baremetal_objectstorage_namespace" "t" {}
```

## Argument Reference

This datasource takes no arguments

## Attribute Reference

* `namespace` - (Computed) The namespace of the object storage that the object is in.
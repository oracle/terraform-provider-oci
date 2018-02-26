
# oci\_objectstorage\_namespaces

## Namespace DataSource

Gets a namespace

### Get Operation
Gets the name of the namespace for the user making the request. An account name must be unique, must start with a
letter, and can have up to 15 lowercase letters and numbers. You cannot use spaces or special characters.

The following arguments are supported:

* No arguments are necessary

The following attributes are exported:

* `namespace` - (Computed) The name of the user's namespace. 

### Example Usage

```
data "oci_objectstorage_namespaces" "test_namespaces" {
}
```
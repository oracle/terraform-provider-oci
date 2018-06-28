# oci_objectstorage_namespace

## Namespace Singular DataSource

### Namespace Reference

The following attributes are exported:

* `namespace` - (Computed) The name of the user's namespace. 



### Get Operation
Gets the name of the namespace for the user making the request. An account name must be unique, must start with a
letter, and can have up to 15 lowercase letters and numbers. You cannot use spaces or special characters.


The following arguments are supported:

* No arguments are necessary


### Example Usage

```hcl
data "oci_objectstorage_namespace" "test_namespace" {
}
```

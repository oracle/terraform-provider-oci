# oci_objectstorage_namespace

## Namespace Singular DataSource

### Namespace Reference

The following attributes are exported:

* `namespace` - (Computed) The name of the user's namespace. 



### Get Operation
Namespaces are unique. Namespaces are either the tenancy name or a random string automatically generated during
account creation. You cannot edit a namespace.


The following arguments are supported:

* No arguments are necessary


### Example Usage

```hcl
data "oci_objectstorage_namespace" "test_namespace" {
}
```

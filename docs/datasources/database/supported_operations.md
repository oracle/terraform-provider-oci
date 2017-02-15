# baremetal\_database\_supported\_operations

Gets a list of database supported operations.

## Example Usage

```
data "baremetal_database_supported_operations" "t" {}
```

## Argument Reference

This datasource takes no arguments.

## Attributes Reference

The following attributes are exported:

* `supported_operations` - A list of the database supported operations.

## Supported Operations Reference

* `id` - The OCID that uniquely identifies the operation.
